package api

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

		page(id: ID!): Page
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

		createPage(input: PageCreateInput!): PageCreateResult
		createPageField(id: ID!, input: PageFieldCreateInput!): PageFieldOperationResult
		renamePageField(id: ID!, input: PageFieldRenameInput!): PageFieldOperationResult
		updatePageField(id: ID!, input: PageFieldUpdateInput!): PageFieldOperationResult
		removePageField(id: ID!, input: PageFieldRemoveInput!): PageFieldOperationResult
		removePage(id: ID!): PageRemoveResult
	}
	
	interface Node {
		id: ID!
	}
	interface CreateResult {
		userErrors: [UserError]
	}
	interface UpdateResult {
		userErrors: [UserError]
	}
	interface RemoveResult {
		userErrors: [UserError]
		removedObjectId: ID
	}
	interface OperationResult {
		userErrors: [UserError]
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
		pages: [Page]
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

	type Page implements Node {
		id: ID!
		name: String!
		fields: [PageField]
		parent: Container!
	}
	type PageField {
		name: String!
		type: String!
		value: String
	}
	type PageCreateResult implements CreateResult {
		userErrors: [UserError]
		page: Page
	}
	type PageRemoveResult implements RemoveResult {
		removedObjectId: ID!
	}	
	type PageFieldOperationResult implements OperationResult {
		userErrors: [UserError]
		page: Page
	}

	input PageCreateInput {
		name: String!
		parentId: ID!
		fields: [PageFieldCreateInput]
	}
  input PageFieldCreateInput {
    name: String!
    type: String!
    value: String
  }
	input PageFieldRenameInput {
		name: String!
		renameTo: String!
	}
	input PageFieldUpdateInput {
		name: String!
		value: String!
	}
	input PageFieldRemoveInput {
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
