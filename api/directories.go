package api

import (
	"github.com/dominik-zeglen/inkster/core"

	"github.com/globalsign/mgo/bson"
	gql "github.com/graph-gophers/graphql-go"
)

// Type resolvers
type directoryResolver struct {
	dataSource core.Adapter
	data       *core.Directory
}

func (res *directoryResolver) ID() gql.ID {
	globalID := toGlobalID("directory", res.data.ID)
	return gql.ID(globalID)
}
func (res *directoryResolver) CreatedAt() string {
	return res.data.CreatedAt
}
func (res *directoryResolver) UpdatedAt() string {
	return res.data.UpdatedAt
}
func (res *directoryResolver) Name() string {
	return res.data.Name
}
func (res *directoryResolver) Parent() *directoryResolver {
	if res.data.ParentID == "" {
		return nil
	}
	parent, err := res.dataSource.GetDirectory(res.data.ParentID)
	if err != nil {
		panic(err)
	}
	return &directoryResolver{
		dataSource: res.dataSource,
		data:       &parent,
	}
}
func (res *directoryResolver) Children() *[]*directoryResolver {
	var resolverList []*directoryResolver
	directories, err := res.dataSource.GetDirectoryChildrenList(res.data.ID)
	if err != nil {
		panic(err)
	}
	for index := range directories {
		resolverList = append(
			resolverList,
			&directoryResolver{
				dataSource: res.dataSource,
				data:       &directories[index],
			},
		)
	}
	return &resolverList
}
func (res *directoryResolver) Pages() (*[]*pageResolver, error) {
	var resolverList []*pageResolver
	pages, err := res.dataSource.GetPagesFromDirectory(res.data.ID)
	if err != nil {
		return nil, err
	}
	for index := range pages {
		resolverList = append(
			resolverList,
			&pageResolver{
				dataSource: res.dataSource,
				data:       &pages[index],
			},
		)
	}
	return &resolverList, nil
}

type directoryAddInput struct {
	Name     string
	ParentID *string
}
type createDirectoryArgs struct {
	Input directoryAddInput
}

func (res *Resolver) CreateDirectory(args createDirectoryArgs) *directoryResolver {
	var directory core.Directory
	input := args.Input
	if input.ParentID != nil {
		parentID, err := fromGlobalID("directory", *input.ParentID)
		if err != nil {
			return nil
		}
		directory = core.Directory{
			Name:     input.Name,
			ParentID: parentID,
		}
	} else {
		directory = core.Directory{
			Name: input.Name,
		}
	}
	directory, err := res.dataSource.AddDirectory(directory)
	if err != nil {
		panic(err)
	}
	return &directoryResolver{
		dataSource: res.dataSource,
		data:       &directory,
	}
}

type getDirectoryArgs struct {
	Id gql.ID
}

func (res *Resolver) GetDirectory(args getDirectoryArgs) (*directoryResolver, error) {
	localID, err := fromGlobalID("directory", string(args.Id))
	if err != nil {
		return nil, err
	}
	directory, err := res.
		dataSource.
		GetDirectory(localID)
	if err != nil {
		return nil, err
	}
	return &directoryResolver{
		dataSource: res.dataSource,
		data:       &directory,
	}, nil
}

func (res *Resolver) GetDirectories() *[]*directoryResolver {
	var resolverList []*directoryResolver
	directories, err := res.dataSource.GetDirectoryList()
	if err != nil {
		panic(err)
	}
	for index := range directories {
		resolverList = append(
			resolverList,
			&directoryResolver{
				dataSource: res.dataSource,
				data:       &directories[index],
			},
		)
	}
	return &resolverList
}

func (res *Resolver) GetRootDirectories() *[]*directoryResolver {
	var resolverList []*directoryResolver
	directories, err := res.dataSource.GetRootDirectoryList()
	if err != nil {
		panic(err)
	}
	for index := range directories {
		resolverList = append(
			resolverList,
			&directoryResolver{
				dataSource: res.dataSource,
				data:       &directories[index],
			},
		)
	}
	return &resolverList
}

type updateDirectoryArgs struct {
	ID    gql.ID
	Input struct {
		Name     *string
		ParentID *gql.ID
	}
}

func (res *Resolver) UpdateDirectory(args updateDirectoryArgs) (
	*directoryResolver, error,
) {
	localID, err := fromGlobalID("directory", string(args.ID))
	var localParentID *bson.ObjectId
	if err != nil {
		return nil, err
	}
	if args.Input.ParentID != nil {
		tempID, err := fromGlobalID("directory", string(*args.Input.ParentID))
		localParentID = &tempID
		if err != nil {
			return nil, err
		}
	}
	err = res.dataSource.UpdateDirectory(localID, core.DirectoryInput{
		Name:     args.Input.Name,
		ParentID: localParentID,
	})
	if err != nil {
		return nil, err
	}
	data, err := res.dataSource.GetDirectory(localID)
	if err != nil {
		return nil, err
	}
	return &directoryResolver{
		dataSource: res.dataSource,
		data:       &data,
	}, nil
}

type removeDirectoryArgs struct {
	Id string
}

func (res *Resolver) RemoveDirectory(args removeDirectoryArgs) (bool, error) {
	localID, err := fromGlobalID("directory", args.Id)
	if err != nil {
		return false, err
	}
	err = res.dataSource.RemoveDirectory(localID)
	if err != nil {
		return false, err
	}
	return true, nil
}
