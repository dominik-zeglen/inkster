package graphql

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/dominik-zeglen/ecoknow/mock"
	gql "github.com/graph-gophers/graphql-go"
)

var dataSource = mock.Adapter{}
var resolver = NewResolver(dataSource)
var schema = gql.MustParseSchema(Schema, &resolver)

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
	json.Indent(&out, jsonResult, "", "    ")
	return out.String(), nil
}
