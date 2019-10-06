package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dominik-zeglen/inkster/api"
	apiSchema "github.com/dominik-zeglen/inkster/api/schema"
	appConfig "github.com/dominik-zeglen/inkster/config"
	"github.com/dominik-zeglen/inkster/core"
	"github.com/dominik-zeglen/inkster/mail"
	"github.com/dominik-zeglen/inkster/middleware"
	"github.com/dominik-zeglen/inkster/storage"
	"github.com/go-pg/pg"
	"github.com/graph-gophers/graphql-go"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Server is a type Inkster app uses to hold connections
type Server struct {
	Config       appConfig.Config
	DataSource   core.DataContext
	FileUploader storage.FileUploader
	MailClient   mail.Mailer
	Schema       *graphql.Schema
}

func (app *Server) initDataSource() *Server {
	pgOptions, err := pg.ParseURL(app.Config.Postgres.URI)
	if err != nil {
		panic(err)
	}

	pgSession := pg.Connect(pgOptions)
	if app.Config.Debug.LogQueries {
		pgSession.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
			query, err := event.FormattedQuery()
			if err != nil {
				panic(err)
			}

			log.Printf("%s %s", time.Since(event.StartTime), query)
		})
	}

	app.DataSource = core.DataContext{
		Session: pgSession,
	}

	return app
}

func (app *Server) initMailer() *Server {
	if app.Config.Mail.Backend == appConfig.MailTerm {
		app.MailClient = mail.NewTerminalMailer(app.Config)
	} else if app.Config.Mail.Backend == appConfig.MailAwsSes {
		app.MailClient = mail.NewAwsSesMailer(app.Config)
	}

	return app
}

func (app *Server) initSchema() *Server {
	resolver := api.NewResolver(
		&app.DataSource,
		app.MailClient,
		app.Config.Server.SecretKey,
	)
	app.Schema = graphql.MustParseSchema(apiSchema.String(), &resolver)

	return app
}

func (app *Server) initStorage() *Server {
	if app.Config.Storage.Backend == appConfig.StorageLocal {
		app.FileUploader = storage.NewLocalFileUploader()
	} else if app.Config.Storage.Backend == appConfig.StorageAwsS3 {
		app.FileUploader = storage.NewAwsS3FileUploader(app.Config)
	}

	return app
}

// Init all settings
func (app *Server) Init(configPath string) *Server {
	appConfig := appConfig.Load(configPath)

	app.Config = *appConfig
	return app.
		initDataSource().
		initMailer().
		initSchema().
		initStorage()
}

// Run Inkster app
func (app *Server) Run() {
	if app.Config.Storage.Backend == "local" {
		http.Handle("/static/",
			http.StripPrefix(
				"/static/",
				http.FileServer(http.Dir("/static/")),
			))
	}
	http.Handle("/graphql/", middleware.WithCors(
		app.Config.Server.AllowedHosts,
		newGraphQLHandler(
			app.Schema,
			app.Config.Server.SecretKey,
			app.DataSource,
			app.Config,
		),
	))
	http.Handle("/upload/", middleware.WithCors(
		app.Config.Server.AllowedHosts,
		newUploadHandler(app.FileUploader),
	))

	log.Printf("Running server on port %d\n", app.Config.Server.Port)
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%d", app.Config.Server.Port),
		nil,
	))
}
