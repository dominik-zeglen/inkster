package graphql

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/dominik-zeglen/ecoknow/core"
	"github.com/globalsign/mgo/bson"
)

var Schema = `
	type Query {
		getContainer(id: ID!): Container
		getContainers: [Container]
		getRootContainers: [Container]
	}
	
	type Mutation {
		createContainer(input: ContainerCreateInput!): Container
		updateContainer(id: ID!, input: ContainerUpdateInput!): Container
		removeContainer(id: ID!): Boolean!	
	}
	
	type Container {
		id: String!
		name: String!
		parent: Container
		children: [Container]
	}
	input ContainerCreateInput {
		name: String!
		parentId: ID
	}
	input ContainerUpdateInput {
		name: String
		parentId: ID
	}
	
	schema {
		query: Query
		mutation: Mutation
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
