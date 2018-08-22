package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/dominik-zeglen/inkster/api"
	"github.com/dominik-zeglen/inkster/mongodb"
	"github.com/globalsign/mgo"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

var schema *graphql.Schema
var dataSource = mongodb.Adapter{
	DBName: os.Getenv("INKSTER_DB_NAME"),
}

func checkEnv() bool {
	vars := []string{
		"INKSTER_DB_URI",
		"INKSTER_DB_NAME",
		"INKSTER_STATIC",
		"INKSTER_PORT",
		"INKSTER_SERVE_STATIC",
	}
	for _, env := range vars {
		if os.Getenv(env) == "" {
			return false
		}
	}
	return true
}

func init() {
	if !checkEnv() {
		log.Fatalln("ERROR: Missing environment variables.")
	}
	session, err := mgo.Dial(os.Getenv("INKSTER_DB_URI"))
	if err != nil {
		log.Println("WARNING: Database is offline.")
	}
	dataSource.Session = session
	resolver := api.NewResolver(&dataSource)
	schema = graphql.MustParseSchema(api.Schema, &resolver)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
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

	http.Handle("/graphql/", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("INKSTER_PORT")), nil))
	log.Printf("Running server on port %s\n", os.Getenv("INKSTER_PORT"))
}
