package app

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/dominik-zeglen/inkster/middleware"
	"github.com/graph-gophers/graphql-go/relay"
)

type GraphQLHandler struct {
	http.Handler
}

func (_ GraphQLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dat, err := ioutil.ReadFile("app/graphql.html")
		check(err)
		_, err = w.Write(dat)
		check(err)
	} else {
		middleware.WithJwt(
			&relay.Handler{Schema: schema},
			os.Getenv("INKSTER_SECRET_KEY"),
		).ServeHTTP(w, r)
	}
}

func newGraphQLHandler() GraphQLHandler {
	return GraphQLHandler{}
}
