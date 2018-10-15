package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/dominik-zeglen/inkster/api"
	apiSchema "github.com/dominik-zeglen/inkster/api/schema"
	"github.com/dominik-zeglen/inkster/mailer"
	"github.com/dominik-zeglen/inkster/middleware"
	"github.com/dominik-zeglen/inkster/mongodb"
	"github.com/globalsign/mgo"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

var schema *graphql.Schema
var DataSource mongodb.Adapter

func checkEnv() {
	vars := []string{
		"INKSTER_DB_URI",
		"INKSTER_STATIC",
		"INKSTER_PORT",
		"INKSTER_SERVE_STATIC",
		"INKSTER_SMTP_HOST",
		"INKSTER_SMTP_LOGIN",
		"INKSTER_SMTP_ADDR",
		"INKSTER_SMTP_PASS",
		"INKSTER_SMTP_PORT",
	}
	for _, env := range vars {
		if os.Getenv(env) == "" {
			log.Fatalf("ERROR: Missing environment variable: %s", env)
		}
	}
}

func init() {
	checkEnv()
	DataSource = mongodb.Adapter{
		DBName: os.Getenv("INKSTER_DB_NAME"),
	}
	session, err := mgo.Dial(os.Getenv("INKSTER_DB_URI"))
	if err != nil {
		log.Println("WARNING: Database is offline.")
	}
	var mailClient mailer.Mailer
	DataSource.Session = session
	if os.Getenv("INKSTER_DEBUG") == "1" {
		mailClient = &mailer.MockMailClient{}
	} else {
		mailClient = mailer.NewSmtpMailClient(
			os.Getenv("INKSTER_SMTP_LOGIN"),
			os.Getenv("INKSTER_SMTP_ADDR"),
			os.Getenv("INKSTER_SMTP_PASS"),
			os.Getenv("INKSTER_SMTP_HOST"),
			os.Getenv("INKSTER_SMTP_PORT"),
		)
	}
	resolver := api.NewResolver(&DataSource, mailClient, os.Getenv("INKSTER_SECRET_KEY"))
	schema = graphql.MustParseSchema(apiSchema.String(), &resolver)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Run() {
	http.Handle("/panel/static/",
		http.StripPrefix(
			"/panel/static/",
			http.FileServer(http.Dir("panel/build/static")),
		))
	http.Handle("/panel/",
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				log.Println(r.URL)
				dat, err := ioutil.ReadFile("panel/build/index.html")
				check(err)
				_, err = w.Write(dat)
				check(err)
			},
		),
	)
	if os.Getenv("INKSTER_SERVE_STATIC") == "1" {
		http.Handle("/static/",
			http.StripPrefix(
				"/static/",
				http.FileServer(http.Dir(os.Getenv("INKSTER_STATIC"))),
			))
	}
	http.Handle("/graphiql/",
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				dat, err := ioutil.ReadFile("app/graphiql.html")
				check(err)
				_, err = w.Write(dat)
				check(err)
			},
		),
	)
	http.Handle("/upload", http.HandlerFunc(api.UploadHandler))

	http.Handle("/graphql/",
		middleware.WithJwt(
			&relay.Handler{Schema: schema},
			os.Getenv("INKSTER_SECRET_KEY"),
		),
	)

	log.Printf("Running server on port %s\n", os.Getenv("INKSTER_PORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("INKSTER_PORT")), nil))
}
