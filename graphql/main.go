package graphql

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/dominik-zeglen/ecoknow/core"
	"github.com/globalsign/mgo/bson"
)

var Schema = `
	schema {
		query: Query
		mutation: Mutation
	}
	
	type Query {
		getContainer(id: String!): Container
		getContainers: [Container]
		getRootContainers: [Container]
	}
	
	type Mutation {
		createContainer(name: String!, parentId: String): Container
		removeContainer(id: String!): OperationResult
	}
	
	type Container {
		id: String!
		name: String!
		parent: Container
		children: [Container]
	}

	type OperationResult {
		success: Boolean!
		message: String!
	}
`

type Resolver struct {
	dataSource core.Adapter
}

func NewResolver(dataSource core.Adapter) Resolver {
	return Resolver{dataSource: dataSource}
}

func toGlobalID(dataType string, ID bson.ObjectId) string {
	data := dataType + ":" + string(ID)
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func fromGlobalID(dataType string, ID string) (bson.ObjectId, error) {
	data, err := base64.StdEncoding.DecodeString(ID)
	if err != nil {
		panic(err)
	}
	portionedData := strings.Split(string(data), ":")
	if portionedData[0] == dataType {
		return bson.ObjectId(portionedData[1]), nil
	}
	return "", fmt.Errorf("Object types do not match")
}

type operationResult struct {
	success bool
	message string
}

type operationResultResolver struct {
	dataSource core.Adapter
	data       *operationResult
}

func (res *operationResultResolver) Success() bool {
	return res.data.success
}

func (res *operationResultResolver) Message() string {
	return res.data.message
}
