package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/dominik-zeglen/ecoknow/api"
	"github.com/dominik-zeglen/ecoknow/mongodb"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

var schema *graphql.Schema
var dataSource = mongodb.Adapter{
	ConnectionURI: os.Getenv("FOXXY_DB_URI"),
	DBName:        os.Getenv("FOXXY_DB_NAME"),
}

func init() {
	resolver := api.NewResolver(&dataSource)
	schema = graphql.MustParseSchema(api.Schema, &resolver)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	http.Handle("/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		))
	http.Handle("/panel/",
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				dat, err := ioutil.ReadFile("app/graphiql.html")
				check(err)
				w.Write(dat)
			},
		),
	)
	http.Handle("/graphiql/",
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				dat, err := ioutil.ReadFile("app/graphiql.html")
				check(err)
				w.Write(dat)
			},
		),
	)

	http.Handle("/graphql/", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
