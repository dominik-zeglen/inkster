package api

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"os"
	"regexp"

	apiSchema "github.com/dominik-zeglen/inkster/api/schema"
	"github.com/dominik-zeglen/inkster/core"
	"github.com/dominik-zeglen/inkster/mailer"
	"github.com/dominik-zeglen/inkster/middleware"
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
	dbHost := os.Getenv("POSTGRES_HOST")
	pgOptions, err := pg.ParseURL(dbHost)
	if err != nil {
		panic(err)
	}
	if os.Getenv("CI") == "" {
		pgOptions.Database = "test_" + pgOptions.Database
		dbOptions, err := pq.ParseURL(os.Getenv("POSTGRES_HOST"))
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

func execQuery(query string, variables string, ctx *context.Context) (string, error) {
	defaultClaims := middleware.UserClaims{
		Email: "user1@example.com",
		ID:    1,
	}
	defaultContext := context.WithValue(context.TODO(), "user", &defaultClaims)

	userContext := ctx
	if ctx == nil {
		userContext = &defaultContext
	}

	var vs map[string]interface{}

	err := json.Unmarshal([]byte(variables), &vs)
	if err != nil {
		return "", err
	}

	result := schema.Exec(*userContext, query, "", vs)
	jsonResult, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	var out bytes.Buffer
	err = json.Indent(&out, jsonResult, "", "    ")
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func resetDatabase() {
	if err := fixtures.Load(); err != nil {
		panic(err)
	}
	mailClient.Reset()
}
