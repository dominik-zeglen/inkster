package graphql

import "github.com/dominik-zeglen/ecoknow/core"

// Type resolvers
type containerResolver struct {
	dataSource core.Adapter
	data       *core.Container
}

func (res *containerResolver) Id() int32 {
	return res.data.ID
}
func (res *containerResolver) Name() string {
	return res.data.Name
}
func (res *containerResolver) Parent() *containerResolver {
	if &res.data.ParentID == nil {
		return nil
	}
	parent, err := res.dataSource.GetContainer(res.data.ParentID)
	if err != nil {
		panic(err)
	}
	return &containerResolver{
		dataSource: res.dataSource,
		data:       &parent,
	}
}
func (res *containerResolver) Children() *[]*containerResolver {
	var resolverList []*containerResolver
	containers, err := res.dataSource.GetContainerChildrenList(res.data.ID)
	if err != nil {
		panic(err)
	}
	for index := range containers {
		resolverList = append(
			resolverList,
			&containerResolver{
				dataSource: res.dataSource,
				data:       &containers[index],
			},
		)
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
			ParentID: *args.ParentId,
		}
	} else {
		container = core.Container{
			Name: args.Name,
		}
	}
	container, err := res.dataSource.AddContainer(container)
	if err != nil {
		panic(err)
	}
	return &containerResolver{
		dataSource: res.dataSource,
		data:       &container,
	}
}

func (res *Resolver) GetContainer(args struct{ Id int32 }) *containerResolver {
	container, err := res.dataSource.GetContainer(args.Id)
	if err != nil {
		panic(err)
	}
	return &containerResolver{
		dataSource: res.dataSource,
		data:       &container,
	}
}

func (res *Resolver) GetContainers() *[]*containerResolver {
	var resolverList []*containerResolver
	containers, err := res.dataSource.GetContainerList()
	if err != nil {
		panic(err)
	}
	for index := range containers {
		resolverList = append(
			resolverList,
			&containerResolver{
				dataSource: res.dataSource,
				data:       &containers[index],
			},
		)
	}
	return &resolverList
}

func (res *Resolver) GetRootContainers() *[]*containerResolver {
	var resolverList []*containerResolver
	containers, err := res.dataSource.GetRootContainerList()
	if err != nil {
		panic(err)
	}
	for index := range containers {
		resolverList = append(
			resolverList,
			&containerResolver{
				dataSource: res.dataSource,
				data:       &containers[index],
			},
		)
	}
	return &resolverList
}

// RemoveContainer resolves RemoveContainer query
func (res *Resolver) RemoveContainer(args struct{ Id int32 }) *operationResultResolver {
	err := res.dataSource.RemoveContainer(args.Id)
	if err != nil {
		message := err.Error()
		return &operationResultResolver{
			dataSource: res.dataSource,
			data: &operationResult{
				success: false,
				message: message,
			},
		}
	}
	return &operationResultResolver{
		dataSource: res.dataSource,
		data: &operationResult{
			success: true,
			message: "",
		},
	}
}
