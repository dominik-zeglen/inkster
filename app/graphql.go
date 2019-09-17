package app

import (
	"io/ioutil"
	"net/http"

	appConfig "github.com/dominik-zeglen/inkster/config"
	"github.com/dominik-zeglen/inkster/core"
	"github.com/dominik-zeglen/inkster/middleware"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type GraphQLHandler struct {
	http.Handler
	dataSource core.DataContext
	schema     *graphql.Schema
	secretKey  string
	config     appConfig.Config
}

func (handler GraphQLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dat, err := ioutil.ReadFile("app/graphql.html")
		check(err)
		_, err = w.Write(dat)
		check(err)
	} else {
		middleware.WithConfig(middleware.WithWebsite(middleware.WithJwt(
			&relay.Handler{
				Schema: handler.schema,
			},
			handler.secretKey,
			handler.dataSource,
		), handler.dataSource), handler.config).ServeHTTP(w, r)
	}
}

func newGraphQLHandler(
	schema *graphql.Schema,
	secretKey string,
	dataSource core.DataContext,
	config appConfig.Config,
) GraphQLHandler {
	return GraphQLHandler{
		dataSource: dataSource,
		schema:     schema,
		secretKey:  secretKey,
		config:     config,
	}
}
