package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dominik-zeglen/inkster/api"
	apiSchema "github.com/dominik-zeglen/inkster/api/schema"
	"github.com/dominik-zeglen/inkster/mailer"
	"github.com/dominik-zeglen/inkster/postgres"
	"github.com/go-pg/pg"
	"github.com/graph-gophers/graphql-go"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type AppServer struct {
	Config     AppConfig
	DataSource postgres.Adapter
	MailClient mailer.Mailer
	Schema     *graphql.Schema
}

func (app *AppServer) initDataSource() *AppServer {
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

	app.DataSource = postgres.Adapter{
		Session: pgSession,
	}

	return app
}

func (app *AppServer) initMailer() *AppServer {
	if app.Config.SMTP.UseDummy {
		app.MailClient = &mailer.MockMailClient{}
	} else {
		app.MailClient = mailer.NewSmtpMailClient(
			app.Config.SMTP.Login,
			app.Config.SMTP.Address,
			app.Config.SMTP.Password,
			app.Config.SMTP.Host,
			app.Config.SMTP.Port,
		)
	}

	return app
}

func (app *AppServer) initSchema() *AppServer {
	resolver := api.NewResolver(
		&app.DataSource,
		app.MailClient,
		app.Config.Server.SecretKey,
	)
	app.Schema = graphql.MustParseSchema(apiSchema.String(), &resolver)

	return app
}

func (app *AppServer) Init(configFilePath string) *AppServer {
	appConfig, err := LoadConfig(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	app.Config = *appConfig
	return app.
		initDataSource().
		initMailer().
		initSchema()
}

func (app *AppServer) Run() {
	http.Handle("/panel/static/",
		http.StripPrefix(
			"/panel/static/",
			http.FileServer(http.Dir("panel/build/static")),
		))
	http.Handle("/panel/",
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				dat, err := ioutil.ReadFile("panel/build/index.html")
				check(err)
				_, err = w.Write(dat)
				check(err)
			},
		),
	)
	if app.Config.Server.ServeStatic {
		http.Handle("/static/",
			http.StripPrefix(
				"/static/",
				http.FileServer(http.Dir(app.Config.Server.StaticPath)),
			))
	}
	http.Handle("/graphql/", newGraphQLHandler(
		app.Schema,
		app.Config.Server.SecretKey,
	))
	http.Handle("/upload", http.HandlerFunc(api.UploadHandler))

	log.Printf("Running server on port %s\n", app.Config.Server.Port)
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%s", app.Config.Server.Port),
		nil,
	))
}
