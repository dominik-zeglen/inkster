package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/dominik-zeglen/inkster/mock"
	test "github.com/dominik-zeglen/inkster/testing"
	gql "github.com/graph-gophers/graphql-go"
)

var dataSource = mock.Adapter{
	GetTime: func() string { return "2017-07-07T10:00:00.000Z" },
}
var resolver = NewResolver(dataSource)
var schema = gql.MustParseSchema(Schema, &resolver)

var ErrNoError = fmt.Errorf("Did not return error")

func execQuery(query string, variables string) (string, error) {
	var vs map[string]interface{}

	err := json.Unmarshal([]byte(variables), &vs)
	if err != nil {
		return "", err
	}

	result := schema.Exec(context.TODO(), query, "", vs)
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
	dataSource.ResetMockDatabase(test.Directories, test.Templates, test.Pages)
}
