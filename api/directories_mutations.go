package api

import (
	"context"

	"github.com/dominik-zeglen/inkster/core"
	"github.com/go-pg/pg"
	gql "github.com/graph-gophers/graphql-go"
)

type directoryOperationResult struct {
	errors    []core.ValidationError
	directory *core.Directory
}

type directoryOperationResultResolver struct {
	dataSource core.Adapter
	data       directoryOperationResult
}

func (res *directoryOperationResultResolver) Errors() []*inputErrorResolver {
	var resolverList []*inputErrorResolver
	for i := range res.data.errors {
		resolverList = append(
			resolverList,
			&inputErrorResolver{
				err: res.data.errors[i],
			},
		)
	}
	return resolverList
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

type directoryAddInput struct {
	Name        string `validate:"min=3"`
	ParentID    *string
	IsPublished *bool
}
type createDirectoryArgs struct {
	Input directoryAddInput `validate:"dive"`
}

func (args createDirectoryArgs) validate(dataSource core.Adapter) (
	[]core.ValidationError,
	error,
) {
	validationErrors := core.ValidateModel(args)

	if args.Input.ParentID != nil {
		localID, err := fromGlobalID("directory", *args.Input.ParentID)
		if err != nil {
			return nil, err
		}
		_, err = dataSource.GetDirectory(localID)
		if err != nil {
			if err == pg.ErrNoRows {
				validationErrors = append(validationErrors, core.ValidationError{
					Code:  core.ErrDoesNotExist,
					Field: "ParentID",
					Param: args.Input.ParentID,
				})
			} else {
				return nil, err
			}
		}
	}

	return validationErrors, nil
}

func (res *Resolver) CreateDirectory(
	ctx context.Context,
	args createDirectoryArgs,
) (*directoryOperationResultResolver, error) {
	if !checkPermission(ctx) {
		return nil, errNoPermissions
	}

	var directory core.Directory
	input := args.Input
	if input.ParentID != nil {
		parentID, err := fromGlobalID("directory", *input.ParentID)
		if err != nil {
			return nil, err
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

	validationErrors, err := args.validate(res.dataSource)
	if err != nil {
		return nil, err
	}
	if len(validationErrors) > 0 {
		return &directoryOperationResultResolver{
			dataSource: res.dataSource,
			data: directoryOperationResult{
				errors:    validationErrors,
				directory: nil,
			},
		}, nil
	}

	directory, err = res.dataSource.AddDirectory(directory)
	if err != nil {
		return nil, err
	}
	return &directoryOperationResultResolver{
		data: directoryOperationResult{
			directory: &directory,
			errors:    validationErrors,
		},
		dataSource: res.dataSource,
	}, nil
}

type updateDirectoryArgs struct {
	ID    gql.ID
	Input struct {
		Name        *string `validate:"omitempty,min=3"`
		ParentID    *string
		IsPublished *bool
	}
}

func (args updateDirectoryArgs) validate(dataSource core.Adapter) (
	[]core.ValidationError,
	error,
) {
	validationErrors := core.ValidateModel(args)

	if args.Input.ParentID != nil {
		if string(args.ID) == *args.Input.ParentID {
			validationErrors = append(validationErrors, core.ValidationError{
				Code:  core.ErrNotEqual,
				Field: "ParentID",
				Param: args.Input.ParentID,
			})
		} else {
			localID, err := fromGlobalID("directory", *args.Input.ParentID)
			if err != nil {
				return nil, err
			}
			_, err = dataSource.GetDirectory(localID)
			if err != nil {
				if err == pg.ErrNoRows {
					validationErrors = append(validationErrors, core.ValidationError{
						Code:  core.ErrDoesNotExist,
						Field: "ParentID",
						Param: args.Input.ParentID,
					})
				} else {
					return nil, err
				}
			}
		}
	}

	return validationErrors, nil
}

func (res *Resolver) UpdateDirectory(
	ctx context.Context,
	args updateDirectoryArgs,
) (*directoryOperationResultResolver, error) {
	if !checkPermission(ctx) {
		return nil, errNoPermissions
	}

	validationErrors, err := args.validate(res.dataSource)
	if err != nil {
		return nil, err
	}
	if len(validationErrors) > 0 {
		return &directoryOperationResultResolver{
			data: directoryOperationResult{
				directory: nil,
				errors:    validationErrors,
			},
			dataSource: res.dataSource,
		}, nil
	}

	localID, err := fromGlobalID("directory", string(args.ID))
	var localParentID *int
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
	directory, err := res.dataSource.GetDirectory(localID)
	if err != nil {
		return nil, err
	}
	return &directoryOperationResultResolver{
		data: directoryOperationResult{
			directory: &directory,
			errors:    validationErrors,
		},
		dataSource: res.dataSource,
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
