package graphql

import (
	"github.com/dominik-zeglen/ecoknow/core"
)

var Schema = `
	schema {
		query: Query
		mutation: Mutation
	}
	
	type Query {
		getContainer(id: Int!): Container
		getContainers: [Container]
		getRootContainers: [Container]
	}
	
	type Mutation {
		createContainer(name: String!, parentId: Int): Container
		removeContainer(id: Int!): OperationResult
	}
	
	type Container {
		id: Int!
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
