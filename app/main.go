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
	"github.com/dominik-zeglen/inkster/mailer"
	"github.com/dominik-zeglen/inkster/middleware"
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
	Config     appConfig.Config
	DataSource core.DataContext
	MailClient mailer.Mailer
	Schema     *graphql.Schema
}

func (app *Server) initDataSource() *Server {
	pgOptions, err := pg.ParseURL(app.Config.Postgres.URI)
	if err != nil {
		panic(err)
	}

	pgSession := pg.Connect(pgOptions)
	if app.Config.Postgres.LogQueries {
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
	if app.Config.Mail.WebhookURL != "" {
		app.MailClient = mailer.NewWebhookMailClient(
			app.Config.Mail.WebhookURL,
			app.Config.Mail.WebhookSecret,
		)
	} else if app.Config.Mail.SESAccessKey != "" {
		app.MailClient = mailer.NewSESMailClient(
			app.Config.Mail.SESAccessKey,
			app.Config.Mail.SESSecretAccessKey,
			app.Config.AWS.Region,
			app.Config.Mail.SESSender,
		)
	} else {
		log.Println("Warning: using dummy mailer")
		app.MailClient = &mailer.MockMailClient{}
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

// Init all settings
func (app *Server) Init() *Server {
	appConfig := appConfig.Load()

	app.Config = *appConfig
	return app.
		initDataSource().
		initMailer().
		initSchema()
}

// Run Inkster app
func (app *Server) Run() {
	if app.Config.Server.ServeStatic {
		http.Handle("/static/",
			http.StripPrefix(
				"/static/",
				http.FileServer(http.Dir(app.Config.Server.StaticPath)),
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
	http.Handle("/upload", middleware.WithCors(
		app.Config.Server.AllowedHosts,
		http.HandlerFunc(api.UploadHandler),
	))

	log.Printf("Running server on port %d\n", app.Config.Server.Port)
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%d", app.Config.Server.Port),
		nil,
	))
}
