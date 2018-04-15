package graphql

import (
	"github.com/dominik-zeglen/foxxy/core"
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
		removeContainer(id: Int!): String
	}
	
	type Container {
		id: Int!
		name: String!
		parent: Container
		children: [Container]
	}
`

type Resolver struct{}

// Type resolvers ----
type containerResolver struct {
	data *core.Container
}

func (container *containerResolver) Id() int32 {
	return container.data.Id
}
func (container *containerResolver) Name() string {
	return container.data.Name
}
func (container *containerResolver) Parent() *containerResolver {
	if &container.data.ParentId == nil {
		return nil
	}
	parent, err := core.GetContainer(container.data.ParentId)
	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			return nil
		}
		panic(err)
	}
	return &containerResolver{&parent}
}
func (container *containerResolver) Children() *[]*containerResolver {
	var resolverList []*containerResolver
	containers, err := core.GetContainerChildrenList(container.data.Id)
	if err != nil {
		panic(err)
	}
	for index := range containers {
		resolverList = append(resolverList, &containerResolver{&containers[index]})
	}
	return &resolverList
}

// Query resolvers ----
type addContainerArgs struct {
	Name     string
	ParentId *int32
}

func (res *Resolver) CreateContainer(args addContainerArgs) *containerResolver {
	var container core.Container
	if args.ParentId != nil {
		container = core.Container{
			Name:     args.Name,
			ParentId: *args.ParentId,
		}
	} else {
		container = core.Container{
			Name: args.Name,
		}
	}
	container, err := core.AddContainer(container)
	if err != nil {
		panic(err)
	}
	return &containerResolver{&container}
}

func (res *Resolver) GetContainer(args struct{ Id int32 }) *containerResolver {
	container, err := core.GetContainer(args.Id)
	if err != nil {
		panic(err)
	}
	return &containerResolver{&container}
}

func (res *Resolver) GetContainers() *[]*containerResolver {
	var resolverList []*containerResolver
	containers, err := core.GetContainerList()
	if err != nil {
		panic(err)
	}
	for index := range containers {
		resolverList = append(resolverList, &containerResolver{&containers[index]})
	}
	return &resolverList
}

func (res *Resolver) GetRootContainers() *[]*containerResolver {
	var resolverList []*containerResolver
	containers, err := core.GetRootContainerList()
	if err != nil {
		panic(err)
	}
	for index := range containers {
		resolverList = append(resolverList, &containerResolver{&containers[index]})
	}
	return &resolverList
}

func (res *Resolver) RemoveContainer(args struct{ Id int32 }) *string {
	err := core.RemoveContainer(args.Id).Error()
	return &err
}
