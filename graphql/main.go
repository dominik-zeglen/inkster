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

		template(id: ID!): Template
		templates: [Template]
	}
	
	type Mutation {
		createContainer(input: ContainerCreateInput!): Container
		updateContainer(id: ID!, input: ContainerUpdateInput!): Container
		removeContainer(id: ID!): Boolean!

		createTemplate(input: TemplateCreateInput!): Template
		templateUpdate(id: ID!, input: TemplateUpdateInput!): TemplateUpdateResult
		createTemplateField(id: ID!, input: TemplateFieldCreateInput!): TemplateUpdateResult
		removeTemplateField(id: ID!, input: TemplateFieldRemoveInput!): TemplateUpdateResult
		removeTemplate(id: ID!): TemplateRemoveResult
	}
	
	interface Node {
		id: ID!
	}
	interface UpdateResult {
		userErrors: [UserError]
	}
	interface RemoveResult {
		userErrors: [UserError]
		removedObjectId: ID
	}

	type UserError {
		field: String!
		message: String!
	}

	type Container implements Node {
		id: ID!
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

	type Template implements Node {
		id: ID!
		name: String!
		fields: [TemplateField]
	}
	type TemplateField {
		name: String!
		type: String!
	}
	type TemplateUpdateResult implements UpdateResult {
		userErrors: [UserError]
		template: Template
	}
	type TemplateRemoveResult implements RemoveResult {
		userErrors: [UserError]
		removedObjectId: ID
	}
		
	input TemplateCreateInput {
		name: String!
		fields: [TemplateFieldCreateInput!]
	}
	input TemplateUpdateInput {
		name: String!
	}
	input TemplateFieldCreateInput {
		name: String!
		type: String!
	}
	input TemplateFieldRemoveInput {
		name: String!
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

type userError struct {
	field   string
	message string
}

type userErrorResolver struct {
	data userError
}

func (res *userErrorResolver) Field() string {
	return res.data.field
}

func (res *userErrorResolver) Message() string {
	return res.data.message
}
