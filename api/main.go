package api

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/dominik-zeglen/inkster/core"
	"github.com/dominik-zeglen/inkster/mailer"
	"github.com/globalsign/mgo/bson"
)

var Schema = `
	type Query {
		getDirectory(id: ID!): Directory
		getDirectories: [Directory]
		getRootDirectories: [Directory]

		template(id: ID!): Template
		templates: [Template]

		page(id: ID!): Page

		user(id: ID!): User
		users: [User]

		verifyToken(token: String!): Boolean!
	}
	
	type Mutation {
		createDirectory(input: DirectoryCreateInput!): Directory
		updateDirectory(id: ID!, input: DirectoryUpdateInput!): Directory
		removeDirectory(id: ID!): Boolean!

		createTemplate(input: TemplateCreateInput!): Template
		templateUpdate(id: ID!, input: TemplateUpdateInput!): TemplateUpdateResult
		createTemplateField(id: ID!, input: TemplateFieldCreateInput!): TemplateUpdateResult
		removeTemplateField(id: ID!, input: TemplateFieldRemoveInput!): TemplateUpdateResult
		removeTemplate(id: ID!): TemplateRemoveResult

		createPage(input: PageCreateInput!): PageOperationResult
		createPageField(id: ID!, input: PageFieldCreateInput!): PageFieldOperationResult
		renamePageField(id: ID!, input: PageFieldRenameInput!): PageFieldOperationResult
		updatePage(id: ID!, input: PageUpdateInput, addFields: [PageFieldCreateInput!], removeFields: [String!]): PageOperationResult

		updatePageField(id: ID!, input: PageFieldUpdateInput!): PageFieldOperationResult
		removePageField(id: ID!, input: PageFieldRemoveInput!): PageFieldOperationResult
		removePage(id: ID!): PageRemoveResult

		createUser(input: UserCreateInput!, sendInvitation: Boolean): UserOperationResult!
		removeUser(id: ID!): UserRemoveResult!
		updateUser(id: ID!, input: UserUpdateInput!): UserOperationResult!
	}
	
	type UserError {
		field: String!
		message: String!
	}

	type Directory {
		id: ID!
		createdAt: String!
		updatedAt: String!
		name: String!
		parent: Directory
		children: [Directory]
		pages: [Page]
	}
	input DirectoryCreateInput {
		name: String!
		parentId: ID
	}
	input DirectoryUpdateInput {
		name: String
		parentId: ID
	}

	type Template{
		id: ID!
		createdAt: String!
		updatedAt: String!
		name: String!
		fields: [TemplateField]
	}
	type TemplateField {
		name: String!
		type: String!
	}
	type TemplateUpdateResult  {
		userErrors: [UserError]
		template: Template
	}
	type TemplateRemoveResult  {
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

	type Page {
		id: ID!
		createdAt: String!
		updatedAt: String!
		name: String!
		slug: String!
		fields: [PageField]
		parent: Directory!
	}
	type PageField {
		name: String!
		type: String!
		value: String
	}
	type PageOperationResult {
		userErrors: [UserError]
		page: Page
	}
	type PageRemoveResult {
		removedObjectId: ID!
	}	
	type PageFieldOperationResult {
		userErrors: [UserError]
		page: Page
	}

	input PageCreateInput {
		name: String!
		parentId: ID!
		fields: [PageFieldCreateInput]
	}
	input PageUpdateInput {
		name: String
		slug: String
		parentId: ID
		fields: [PageFieldUpdate2Input!]
	}
  input PageFieldCreateInput {
    name: String!
    type: String!
    value: String!
  }
	input PageFieldRenameInput {
		name: String!
		renameTo: String!
	}
	input PageFieldUpdateInput {
		name: String!
		value: String!
	}
	input PageFieldUpdate2Input {
		name: String!
		update: PageFieldUpdate2InputPartial!
	}
	input PageFieldUpdate2InputPartial {
		name: String
		value: String
	}
	input PageFieldRemoveInput {
		name: String!
	}

	type User {
		id: ID!
		createdAt: String!
		updatedAt: String!
		email: String!
		isActive: Boolean!
	}
	type UserOperationResult {
		errors: [UserError!]!
		user: User
	}
	type UserRemoveResult {
		removedObjectId: ID
	}

	input UserCreateInput {
		email: String!
		password: String
	}
	input UserUpdateInput {
		isActive: Boolean
		email: String
	}
	
	schema {
		query: Query
		mutation: Mutation
	}
`

type Resolver struct {
	dataSource core.Adapter
	key        string
	mailer     mailer.Mailer
}

func NewResolver(dataSource core.Adapter, mailer mailer.Mailer, key string) Resolver {
	return Resolver{
		dataSource: dataSource,
		key:        key,
		mailer:     mailer,
	}
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
