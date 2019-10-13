package api

import (
	"context"

	"github.com/dominik-zeglen/inkster/core"
	"github.com/dominik-zeglen/inkster/middleware"
	"github.com/gosimple/slug"
	gql "github.com/graph-gophers/graphql-go"
)

type pageCreateResult struct {
	page             *core.Page
	validationErrors []core.ValidationError
}

type pageCreateResultResolver struct {
	dataSource core.AbstractDataContext
	data       pageCreateResult
}

func (res *pageCreateResultResolver) Page() *pageResolver {
	return &pageResolver{
		dataSource: res.dataSource,
		data:       res.data.page,
	}
}
func (res *pageCreateResultResolver) Errors() []inputErrorResolver {
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
) (*pageCreateResultResolver, error) {
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

	return &pageCreateResultResolver{
		dataSource: res.dataSource,
		data: pageCreateResult{
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

func cleanUpdatePageInput(
	id int,
	input UpdatePageInput,
	dataSource core.AbstractDataContext,
) (core.PageInput, []core.ValidationError, error) {
	validationErrors := []core.ValidationError{}
	pageInput := core.PageInput{}

	if input.ParentID != nil {
		localID, err := fromGlobalID("page", *input.ParentID)
		if err != nil {
			return pageInput, validationErrors, err
		}
		pageInput.ParentID = &localID
	}
	pageInput.Name = input.Name
	pageInput.IsPublished = input.IsPublished

	validationErrors = append(validationErrors, pageInput.Validate()...)

	return pageInput, validationErrors, nil
}

func (res *Resolver) UpdatePage(
	ctx context.Context,
	args UpdatePageArgs,
) (*pageCreateResultResolver, error) {
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

	_, validationErrors, err := cleanUpdatePageInput(
		localID,
		args.Input,
		res.dataSource,
	)

	if err != nil {
		return nil, err
	}

	if len(validationErrors) > 0 {
		return &pageCreateResultResolver{
			dataSource: res.dataSource,
			data: pageCreateResult{
				page:             nil,
				validationErrors: validationErrors,
			},
		}, nil
	}

	page.ID = localID
	page.UpdatedAt = res.
		dataSource.
		GetCurrentTime()

	query := res.
		dataSource.
		DB().
		Model(&page).
		Column("updated_at")

	if args.Input.IsPublished != nil {
		page.IsPublished = *args.Input.IsPublished
		query = query.Column("is_published")
	}
	if args.Input.Name != nil {
		page.Name = *args.Input.Name
		query = query.Column("name")
	}
	if args.Input.ParentID != nil {
		localParentID, err := fromGlobalID("directory", *args.Input.ParentID)
		if err != nil {
			return nil, err
		}
		page.ParentID = localParentID
		query = query.Column("parent_id")
	}
	if args.Input.Slug != nil {
		page.Slug = *args.Input.Slug
		query = query.Column("slug")
	}
	if args.Input.Fields != nil {
		page.Fields = *args.Input.Fields
		query = query.Column("fields")
	}

	_, err = query.
		WherePK().
		Update()

	if err != nil {
		return nil, err
	}

	err = res.
		dataSource.
		DB().
		Model(&page).
		WherePK().
		Select()

	if err != nil {
		return nil, err
	}

	return &pageCreateResultResolver{
		dataSource: res.dataSource,
		data: pageCreateResult{
			page: &page,
		},
	}, nil
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

	_, err = res.
		dataSource.
		DB().
		Exec("DELETE FROM pages WHERE id = ?", localID)

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
