package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/lib/pq"
	_ "github.com/lib/pq"

	apiSchema "github.com/dominik-zeglen/inkster/api/schema"
	"github.com/dominik-zeglen/inkster/config"
	"github.com/dominik-zeglen/inkster/core"
	"github.com/dominik-zeglen/inkster/mailer"
	"github.com/dominik-zeglen/inkster/middleware"
	"github.com/dominik-zeglen/inkster/utils"
	"github.com/go-pg/pg"
	"github.com/go-testfixtures/testfixtures"
	gql "github.com/graph-gophers/graphql-go"
)

var dataSource core.MockContext
var mailClient = mailer.MockMailClient{}
var resolver = NewResolver(&dataSource, &mailClient, "secretKey")
var schema = gql.MustParseSchema(apiSchema.String(), &resolver)
var fixtures *testfixtures.Context

var ErrNoError = fmt.Errorf("Did not return error")

func init() {
	conf := config.Load("../")
	dbHost := conf.Postgres.URI
	pgOptions, err := pg.ParseURL(dbHost)
	if err != nil {
		panic(err)
	}
	if !conf.Miscellaneous.CI {
		pgOptions.Database = "test_" + pgOptions.Database
		dbOptions, err := pq.ParseURL(dbHost)
		if err != nil {
			panic(err)
		}
		regex := regexp.MustCompile("dbname=")
		replace := "${1}dbname=test_$2"
		dbHost = regex.ReplaceAllString(dbOptions, replace)
	}

	pgSession := pg.Connect(pgOptions)
	pgAdapter := core.MockContext{
		Session: pgSession,
	}
	dataSource = pgAdapter

	db, err := sql.Open("postgres", dbHost)
	if err != nil {
		panic(err)
	}

	fixtures, err = testfixtures.NewFolder(
		db,
		&testfixtures.PostgreSQL{},
		"../fixtures",
	)
	if err != nil {
		panic(err)
	}

	resetDatabase()
}

func execQuery(
	query string,
	variables string,
	userContext *context.Context,
) (string, error) {
	user := core.User{
		Email: "user1@example.com",
	}
	user.ID = 1

	defaultContext := context.WithValue(
		context.TODO(),
		middleware.UserContextKey,
		&user,
	)

	if userContext == nil {
		userContext = &defaultContext
	}

	var vs map[string]interface{}

	err := json.Unmarshal([]byte(variables), &vs)
	if err != nil {
		return "", err
	}

	result := schema.Exec(*userContext, query, "", vs)

	return utils.PrintJSON(result)
}

func resetDatabase() {
	if err := fixtures.Load(); err != nil {
		panic(err)
	}
	mailClient.Reset()
}
