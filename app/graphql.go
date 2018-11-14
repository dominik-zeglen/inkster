package app

import (
	"io/ioutil"
	"net/http"

	"github.com/dominik-zeglen/inkster/middleware"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type GraphQLHandler struct {
	http.Handler
	schema    *graphql.Schema
	secretKey string
}

func (handler GraphQLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dat, err := ioutil.ReadFile("app/graphql.html")
		check(err)
		_, err = w.Write(dat)
		check(err)
	} else {
		middleware.WithJwt(
			&relay.Handler{
				Schema: handler.schema,
			},
			handler.secretKey,
		).ServeHTTP(w, r)
	}
}

func newGraphQLHandler(
	schema *graphql.Schema,
	secretKey string,
) GraphQLHandler {
	return GraphQLHandler{
		schema:    schema,
		secretKey: secretKey,
	}
}
