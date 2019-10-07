package api

import (
	"context"

	"github.com/dominik-zeglen/inkster/core"
	gql "github.com/graph-gophers/graphql-go"
)

type directoryOperationResult struct {
	directory        *core.Directory
	validationErrors []core.ValidationError
}

type directoryOperationResultResolver struct {
	dataSource core.AbstractDataContext
	data       directoryOperationResult
}

func (res *directoryOperationResultResolver) Errors() []inputErrorResolver {
	return createInputErrorResolvers(res.data.validationErrors)
}

func (res *directoryOperationResultResolver) Directory() *directoryResolver {
	if res.data.directory == nil {
		return nil
	}
	return &directoryResolver{
		dataSource: res.dataSource,
		data:       res.data.directory,
	}
}

type createDirectoryInput struct {
	Name        string
	ParentID    *string
	IsPublished *bool
}
type createDirectoryArgs struct {
	Input createDirectoryInput
}

func (res *Resolver) CreateDirectory(
	ctx context.Context,
	args createDirectoryArgs,
) (*directoryOperationResultResolver, error) {
	if !checkPermission(ctx) {
		return nil, errNoPermissions
	}

	input := args.Input

	directory := core.Directory{}
	directory.Name = args.Input.Name

	if input.IsPublished != nil {
		directory.IsPublished = *input.IsPublished
	}

	if input.ParentID != nil {
		parentID, err := fromGlobalID("directory", *input.ParentID)
		if err != nil {
			return nil, err
		}
		directory.ParentID = &parentID
	}

	insertedDirectory, validationErrors, err := core.CreateDirectory(
		directory,
		res.dataSource,
	)

	return &directoryOperationResultResolver{
		data: directoryOperationResult{
			directory:        insertedDirectory,
			validationErrors: validationErrors,
		},
		dataSource: res.dataSource,
	}, err
}

type updateDirectoryInput struct {
	Name        *string
	ParentID    *string
	IsPublished *bool
}
type updateDirectoryArgs struct {
	ID    gql.ID
	Input updateDirectoryInput
}

func (res *Resolver) UpdateDirectory(
	ctx context.Context,
	args updateDirectoryArgs,
) (*directoryOperationResultResolver, error) {
	if !checkPermission(ctx) {
		return nil, errNoPermissions
	}

	localID, err := fromGlobalID("directory", string(args.ID))
	if err != nil {
		return nil, err
	}

	directory := core.Directory{}
	directory.ID = localID

	err = res.
		dataSource.
		DB().
		Model(&directory).
		Select()

	if args.Input.IsPublished != nil {
		directory.IsPublished = *args.Input.IsPublished
	}
	if args.Input.Name != nil {
		directory.Name = *args.Input.Name
	}
	if args.Input.ParentID != nil {
		parentID, err := fromGlobalID("directory", string(*args.Input.ParentID))
		if err != nil {
			return nil, err
		}
		directory.ParentID = &parentID
	}

	updatedDirectory, validationErrors, err := core.UpdateDirectory(
		directory,
		res.dataSource,
	)

	return &directoryOperationResultResolver{
		data: directoryOperationResult{
			directory:        updatedDirectory,
			validationErrors: validationErrors,
		},
		dataSource: res.dataSource,
	}, nil
}

type removeDirectoryArgs struct {
	ID string
}

func (res *Resolver) RemoveDirectory(
	ctx context.Context,
	args removeDirectoryArgs,
) (bool, error) {
	if !checkPermission(ctx) {
		return false, errNoPermissions
	}
	localID, err := fromGlobalID("directory", args.ID)
	if err != nil {
		return false, err
	}

	err = core.RemoveDirectory(localID, res.dataSource)

	if err != nil {
		return false, err
	}
	return true, nil
}
