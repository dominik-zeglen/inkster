package api

import (
	"context"

	"github.com/dominik-zeglen/inkster/core"
	"github.com/go-pg/pg"
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

type updateDirectoryArgs struct {
	ID    gql.ID
	Input struct {
		Name        *string `validate:"omitempty,min=3"`
		ParentID    *string
		IsPublished *bool
	}
}

func (args updateDirectoryArgs) validate(dataSource core.AbstractDataContext) (
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
			directory := core.Directory{}
			directory.ID = localID
			err = dataSource.
				DB().
				Model(&directory).
				WherePK().
				Select()
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
				directory:        nil,
				validationErrors: validationErrors,
			},
			dataSource: res.dataSource,
		}, nil
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

	directory.UpdatedAt = res.
		dataSource.
		GetCurrentTime()

	query := res.
		dataSource.
		DB().
		Model(&directory).
		Column("updated_at")

	if args.Input.IsPublished != nil {
		directory.IsPublished = *args.Input.IsPublished
		query = query.Column("is_published")
	}
	if args.Input.Name != nil {
		directory.Name = *args.Input.Name
		query = query.Column("name")
	}
	if args.Input.ParentID != nil {
		parentID, err := fromGlobalID("directory", string(*args.Input.ParentID))
		if err != nil {
			return nil, err
		}
		directory.ParentID = &parentID
		query = query.Column("parent_id")
	}

	_, err = query.
		WherePK().
		Update()

	if err != nil {
		return nil, err
	}

	return &directoryOperationResultResolver{
		data: directoryOperationResult{
			directory:        &directory,
			validationErrors: validationErrors,
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

	_, err = res.
		dataSource.
		DB().
		Exec("DELETE FROM directories WHERE id = ?", localID)

	if err != nil {
		return false, err
	}
	return true, nil
}
