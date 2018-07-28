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
	ConnectionURI: os.Getenv("INKSTER_DB_URI"),
	DBName:        os.Getenv("INKSTER_DB_NAME"),
}

func init() {
	if dataSource.ConnectionURI == "" || dataSource.DBName == "" {
		log.Fatalln("ERROR: Missing environment variables.")
	}
	sess, err := mgo.Dial(dataSource.ConnectionURI)
	log.Printf("Making connection to %s\n", dataSource.ConnectionURI)
	if err != nil {
		log.Println("WARNING: Database is offline.")
	} else {
		sess.Close()
	}
	resolver := api.NewResolver(&dataSource)
	schema = graphql.MustParseSchema(api.Schema, &resolver)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.Handle("/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("panel/build/static")),
		))
	http.Handle("/panel/",
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				dat, err := ioutil.ReadFile("panel/build/index.html")
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
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("INKSTER_PORT")), nil))
	log.Printf("Running server on port %s\n", os.Getenv("INKSTER_PORT"))
}
