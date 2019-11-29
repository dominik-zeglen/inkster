package api

import (
	"context"

	"github.com/dominik-zeglen/inkster/core"
	"github.com/dominik-zeglen/inkster/middleware"
	"github.com/gosimple/slug"
	gql "github.com/graph-gophers/graphql-go"
)

type pageOperationResult struct {
	page             *core.Page
	validationErrors []core.ValidationError
}

type pageOperationResultResolver struct {
	dataSource core.AbstractDataContext
	data       pageOperationResult
}

func (res *pageOperationResultResolver) Page() *pageResolver {
	if res.data.page == nil {
		return nil
	}

	return &pageResolver{
		dataSource: res.dataSource,
		data:       *res.data.page,
	}
}
func (res *pageOperationResultResolver) Errors() []inputErrorResolver {
	return createInputErrorResolvers(res.data.validationErrors)
}

type pageRemoveResult struct {
	removedObjectID gql.ID
}

type pageRemoveResultResolver struct {
	dataSource core.AbstractDataContext
	data       pageRemoveResult
}

func (res *pageRemoveResultResolver) RemovedObjectID() gql.ID {
	return res.data.removedObjectID
}

type createPageArgsInput struct {
	Name        string
	ParentID    string
	Slug        *string
	IsPublished *bool
	Fields      *[]core.PageField
}
type createPageArgs struct {
	Input createPageArgsInput
}

func (res *Resolver) CreatePage(
	ctx context.Context,
	args createPageArgs,
) (*pageOperationResultResolver, error) {
	if !checkPermission(ctx) {
		return nil, errNoPermissions
	}

	input := args.Input

	user := ctx.Value(middleware.UserContextKey).(*core.User)
	localID, err := fromGlobalID("directory", input.ParentID)
	if err != nil {
		return nil, err
	}

	page := core.Page{
		Name:     input.Name,
		ParentID: localID,
	}
	page.AuthorID = user.ID

	if input.Slug != nil {
		page.Slug = *input.Slug
	} else {
		page.Slug = slug.Make(page.Name)
	}

	if input.IsPublished != nil {
		page.IsPublished = *input.IsPublished
	}

	if input.Fields != nil {
		page.Fields = *input.Fields
	}

	insertedPage, validationErrors, err := core.CreatePage(
		page,
		res.dataSource,
	)

	return &pageOperationResultResolver{
		dataSource: res.dataSource,
		data: pageOperationResult{
			validationErrors: validationErrors,
			page:             insertedPage,
		},
	}, err
}

type UpdatePageInput struct {
	Fields      *[]core.PageField
	Name        *string
	Slug        *string
	ParentID    *string
	IsPublished *bool
}
type UpdatePageArgs struct {
	ID    gql.ID
	Input UpdatePageInput
}

func (res *Resolver) UpdatePage(
	ctx context.Context,
	args UpdatePageArgs,
) (*pageOperationResultResolver, error) {
	if !checkPermission(ctx) {
		return nil, errNoPermissions
	}

	localID, err := fromGlobalID("page", string(args.ID))
	if err != nil {
		return nil, err
	}

	page := core.Page{}
	page.ID = localID

	err = res.
		dataSource.
		DB().
		Model(&page).
		WherePK().
		Select()

	if err != nil {
		return nil, err
	}

	if args.Input.IsPublished != nil {
		page.IsPublished = *args.Input.IsPublished
	}
	if args.Input.Name != nil {
		page.Name = *args.Input.Name
	}
	if args.Input.Fields != nil {
		page.Fields = *args.Input.Fields
	}
	if args.Input.Slug != nil {
		page.Slug = *args.Input.Slug
	}
	if args.Input.ParentID != nil {
		parentID, err := fromGlobalID("directory", string(*args.Input.ParentID))
		if err != nil {
			return nil, err
		}
		page.ParentID = parentID
	}
	updatedPage, validationErrors, err := core.UpdatePage(
		page,
		res.dataSource,
	)

	return &pageOperationResultResolver{
		data: pageOperationResult{
			page:             updatedPage,
			validationErrors: validationErrors,
		},
		dataSource: res.dataSource,
	}, err
}

type removePageArgs struct {
	ID gql.ID
}

func (res *Resolver) RemovePage(
	ctx context.Context,
	args removePageArgs,
) (*pageRemoveResultResolver, error) {
	if !checkPermission(ctx) {
		return nil, errNoPermissions
	}

	localID, err := fromGlobalID("page", string(args.ID))
	if err != nil {
		return nil, err
	}

	err = core.RemovePage(localID, res.dataSource)
	if err != nil {
		return nil, err
	}

	return &pageRemoveResultResolver{
		dataSource: res.dataSource,
		data: pageRemoveResult{
			removedObjectID: args.ID,
		},
	}, nil
}
