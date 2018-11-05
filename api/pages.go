package api

import (
	"context"

	"github.com/dominik-zeglen/inkster/core"
	"github.com/globalsign/mgo/bson"
	gql "github.com/graph-gophers/graphql-go"
)

type pageCreateResult struct {
	validationErrors []core.ValidationError
	page             *core.Page
}

type pageCreateResultResolver struct {
	dataSource core.Adapter
	data       pageCreateResult
}

func (res *pageCreateResultResolver) Page() *pageResolver {
	return &pageResolver{
		dataSource: res.dataSource,
		data:       res.data.page,
	}
}
func (res *pageCreateResultResolver) Errors() []inputErrorResolver {
	resolverList := []inputErrorResolver{}
	if res.data.validationErrors == nil {
		return nil
	}
	for i := range res.data.validationErrors {
		resolverList = append(
			resolverList,
			inputErrorResolver{
				err: res.data.validationErrors[i],
			},
		)
	}
	return resolverList
}

type pageRemoveResult struct {
	removedObjectID gql.ID
}

type pageRemoveResultResolver struct {
	dataSource core.Adapter
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

func cleanCreatePageInput(input createPageArgsInput) (
	*core.Page,
	error,
) {
	localID, err := fromGlobalID("directory", input.ParentID)
	if err != nil {
		return nil, err
	}

	page := core.Page{
		Name:     input.Name,
		ParentID: bson.ObjectId(localID),
	}

	if input.IsPublished != nil {
		page.IsPublished = *input.IsPublished
	}

	if input.Fields != nil {
		page.Fields = *input.Fields
	}

	return &page, nil
}

func (res *Resolver) CreatePage(
	ctx context.Context,
	args createPageArgs,
) (*pageCreateResultResolver, error) {
	if !checkPermission(ctx) {
		return nil, errNoPermissions
	}

	page, err := cleanCreatePageInput(args.Input)
	if err != nil {
		return nil, err
	}

	errs := page.Validate()
	if len(errs) > 0 {
		return &pageCreateResultResolver{
			dataSource: res.dataSource,
			data: pageCreateResult{
				validationErrors: errs,
				page:             nil,
			},
		}, nil
	}

	result, err := res.dataSource.AddPage(*page)
	if err != nil {
		return nil, err
	}
	return &pageCreateResultResolver{
		dataSource: res.dataSource,
		data: pageCreateResult{
			validationErrors: errs,
			page:             &result,
		},
	}, nil
}

type UpdatePageInput struct {
	Name        *string
	Slug        *string
	ParentID    *string
	IsPublished *bool
}
type UpdatePageArgs struct {
	ID           gql.ID
	Input        *UpdatePageInput
	AddFields    *[]core.PageField
	RemoveFields *[]string
}

func cleanUpdatePageInput(
	input *UpdatePageInput,
	dataSource core.Adapter,
) (core.PageInput, []core.ValidationError, error) {
	validationErrors := []core.ValidationError{}
	pageInput := core.PageInput{}

	if input == nil {
		return pageInput, validationErrors, nil
	}

	if input.Slug != nil {
		_, err := dataSource.GetPageBySlug(*input.Slug)
		if err == nil {
			validationErrors = append(
				validationErrors,
				core.ValidationError{
					Code:  core.ErrNotUnique,
					Field: "Slug",
					Param: input.Slug,
				},
			)
		}
		pageInput.Slug = input.Slug
	}

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

func cleanUpdatePageAddFields(
	addFields []core.PageField,
	fields []core.PageField,
) (
	[]core.PageField,
	[]core.ValidationError,
) {
	validationErrors := []core.ValidationError{}

	fields = append(fields, addFields...)

	for _, field := range fields {
		validationErrors = append(
			validationErrors,
			field.Validate()...,
		)
	}

	return fields, validationErrors
}

func cleanUpdatePageRemoveFields(
	removeFields []string,
	fields []core.PageField,
) []core.PageField {
	cleanedFields := []core.PageField{}

	for _, field := range fields {
		found := false

		for _, removeField := range removeFields {
			if field.Name == removeField {
				found = true
				break
			}
		}

		if !found {
			cleanedFields = append(cleanedFields, field)
		}
	}

	return cleanedFields
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
	page, err := res.dataSource.GetPage(localID)
	if err != nil {
		return nil, err
	}

	if args.Input != nil || args.AddFields != nil || args.RemoveFields != nil {
		pageInput, _, err := cleanUpdatePageInput(
			args.Input,
			res.dataSource,
		)

		if err != nil {
			return nil, err
		}

		if args.AddFields != nil || args.RemoveFields != nil {
			pageInput.Fields = &page.Fields

			if args.AddFields != nil {
				fields, _ := cleanUpdatePageAddFields(
					*args.AddFields,
					*pageInput.Fields,
				)
				pageInput.Fields = &fields
			}

			if args.RemoveFields != nil {
				fields := cleanUpdatePageRemoveFields(
					*args.RemoveFields,
					*pageInput.Fields,
				)
				pageInput.Fields = &fields
			}
		}

		err = res.dataSource.UpdatePage(localID, pageInput)
		if err != nil {
			return nil, err
		}
	}
	page, err = res.dataSource.GetPage(localID)
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
	err = res.dataSource.RemovePage(localID)
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
