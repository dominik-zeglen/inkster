package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/dominik-zeglen/inkster/mailer"
	"github.com/dominik-zeglen/inkster/middleware"
	"github.com/dominik-zeglen/inkster/mock"
	test "github.com/dominik-zeglen/inkster/testing"
	gql "github.com/graph-gophers/graphql-go"
)

var dataSource = mock.Adapter{
	GetTime: func() string { return "2017-07-07T10:00:00.000Z" },
}
var resolver = NewResolver(dataSource, mailer.MockMailClient{}, "secretKey")
var schema = gql.MustParseSchema(Schema, &resolver)

var ErrNoError = fmt.Errorf("Did not return error")

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
	dataSource.ResetMockDatabase(
		test.Directories,
		test.Templates,
		test.Pages,
		test.Users,
	)
}
