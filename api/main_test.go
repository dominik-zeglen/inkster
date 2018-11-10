package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"

	apiSchema "github.com/dominik-zeglen/inkster/api/schema"
	"github.com/dominik-zeglen/inkster/mailer"
	"github.com/dominik-zeglen/inkster/middleware"
	"github.com/dominik-zeglen/inkster/postgres"
	test "github.com/dominik-zeglen/inkster/testing"
	"github.com/go-pg/pg"
	gql "github.com/graph-gophers/graphql-go"
)

var dataSource postgres.Adapter
var mailClient = mailer.MockMailClient{}
var resolver = NewResolver(&dataSource, &mailClient, "secretKey")
var schema = gql.MustParseSchema(apiSchema.String(), &resolver)

var ErrNoError = fmt.Errorf("Did not return error")

func init() {
	pgOptions, err := pg.ParseURL(os.Getenv("POSTGRES_HOST"))
	if err != nil {
		panic(err)
	}
	if os.Getenv("CI") != "" {
		pgOptions.Database = "test_" + pgOptions.Database
	}

	pgSession := pg.Connect(pgOptions)
	pgAdapter := postgres.Adapter{
		GetTime: func() string { return "2017-07-07T10:00:00.000Z" },
		Session: pgSession,
	}
	dataSource = pgAdapter

	resetDatabase()
}

func execQuery(query string, variables string, ctx *context.Context) (string, error) {
	defaultClaims := middleware.UserClaims{
		Email: test.Users[0].Email,
		ID:    test.Users[0].ID,
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
	err := dataSource.ResetMockDatabase(
		test.Directories,
		test.Templates,
		test.Pages,
		test.Users,
	)
	if err != nil {
		panic(err)
	}
	mailClient.Reset()
}
