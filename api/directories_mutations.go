package api

import (
	"context"

	"github.com/dominik-zeglen/inkster/core"
	"github.com/globalsign/mgo/bson"
	gql "github.com/graph-gophers/graphql-go"
)

type directoryAddInput struct {
	Name        string
	ParentID    *string
	IsPublished *bool
}
type createDirectoryArgs struct {
	Input directoryAddInput
}

func (res *Resolver) CreateDirectory(
	ctx context.Context,
	args createDirectoryArgs,
) (*directoryResolver, error) {
	if !checkPermission(ctx) {
		return nil, errNoPermissions
	}
	var directory core.Directory
	input := args.Input
	if input.ParentID != nil {
		parentID, err := fromGlobalID("directory", *input.ParentID)
		if err != nil {
			return nil, nil
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
	if input.IsPublished != nil {
		directory.IsPublished = *input.IsPublished
	}
	directory, err := res.dataSource.AddDirectory(directory)
	if err != nil {
		return nil, err
	}
	return &directoryResolver{
		dataSource: res.dataSource,
		data:       &directory,
	}, nil
}

type updateDirectoryArgs struct {
	ID    gql.ID
	Input struct {
		Name        *string
		ParentID    *gql.ID
		IsPublished *bool
	}
}

func (res *Resolver) UpdateDirectory(
	ctx context.Context,
	args updateDirectoryArgs,
) (*directoryResolver, error) {
	if !checkPermission(ctx) {
		return nil, errNoPermissions
	}
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
		Name:        args.Input.Name,
		ParentID:    localParentID,
		IsPublished: args.Input.IsPublished,
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

func (res *Resolver) RemoveDirectory(
	ctx context.Context,
	args removeDirectoryArgs,
) (bool, error) {
	if !checkPermission(ctx) {
		return false, errNoPermissions
	}
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
