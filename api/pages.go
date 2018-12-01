package api

import (
	"context"

	"github.com/dominik-zeglen/inkster/core"
	"github.com/gosimple/slug"
	gql "github.com/graph-gophers/graphql-go"
)

type pageCreateResult struct {
	validationErrors []core.ValidationError
	page             *core.Page
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

func cleanCreatePageInput(
	input createPageArgsInput,
	dataSource core.AbstractDataContext,
) (
	*core.Page,
	error,
) {
	localID, err := fromGlobalID("directory", input.ParentID)
	if err != nil {
		return nil, err
	}

	page := core.Page{
		Name:     input.Name,
		ParentID: localID,
	}
	page.CreatedAt = dataSource.GetCurrentTime()
	page.UpdatedAt = dataSource.GetCurrentTime()

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

	return &page, nil
}

func (res *Resolver) CreatePage(
	ctx context.Context,
	args createPageArgs,
) (*pageCreateResultResolver, error) {
	if !checkPermission(ctx) {
		return nil, errNoPermissions
	}

	page, err := cleanCreatePageInput(args.Input, res.dataSource)
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

	_, err = res.
		dataSource.
		DB().
		Model(page).
		Insert()

	if err != nil {
		return nil, err
	}

	for fieldIndex, _ := range page.Fields {
		page.Fields[fieldIndex].PageID = page.ID
	}

	_, err = res.
		dataSource.
		DB().
		Model(&page.Fields).
		Insert()

	return &pageCreateResultResolver{
		dataSource: res.dataSource,
		data: pageCreateResult{
			validationErrors: errs,
			page:             page,
		},
	}, nil
}

type UpdatePageFieldsInput struct {
	Name  *string
	Value *string
}
type UpdatePageFields struct {
	ID    gql.ID
	Input UpdatePageFieldsInput
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
	UpdateFields *[]UpdatePageFields
	RemoveFields *[]string
}

func cleanUpdatePageInput(
	id int,
	input *UpdatePageInput,
	dataSource core.AbstractDataContext,
) (core.PageInput, []core.ValidationError, error) {
	validationErrors := []core.ValidationError{}
	pageInput := core.PageInput{}

	if input == nil {
		return pageInput, validationErrors, nil
	}

	if input.Slug != nil {
		foundPage := core.Page{}

		err := dataSource.
			DB().
			Model(&foundPage).
			Where("slug = ?", &input.Slug).
			Select()

		if err == nil {
			if foundPage.ID != id {
				validationErrors = append(
					validationErrors,
					core.ValidationError{
						Code:  core.ErrNotUnique,
						Field: "Slug",
						Param: input.Slug,
					},
				)
			}
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

func cleanUpdatePageAddFields(addFields []core.PageField) []core.ValidationError {
	validationErrors := []core.ValidationError{}

	for _, field := range addFields {
		validationErrors = append(
			validationErrors,
			field.Validate()...,
		)
	}

	return validationErrors
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

	err = res.
		dataSource.
		DB().
		Model(&page).
		Where("id = ?", localID).
		Relation("Fields").
		Select()

	if err != nil {
		return nil, err
	}

	if args.Input != nil ||
		args.AddFields != nil ||
		args.UpdateFields != nil ||
		args.RemoveFields != nil {
		_, validationErrors, err := cleanUpdatePageInput(
			localID,
			args.Input,
			res.dataSource,
		)

		if err != nil {
			return nil, err
		}

		if args.AddFields != nil {
			errs := cleanUpdatePageAddFields(*args.AddFields)
			validationErrors = append(validationErrors, errs...)
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

		if args.AddFields != nil {
			addPageFields := *args.AddFields
			for pageFieldIndex := range addPageFields {
				addPageFields[pageFieldIndex].CreatedAt = res.
					dataSource.
					GetCurrentTime()
				addPageFields[pageFieldIndex].UpdatedAt = res.
					dataSource.
					GetCurrentTime()
				addPageFields[pageFieldIndex].PageID = localID
			}
			_, err = res.
				dataSource.
				DB().
				Model(args.AddFields).
				Insert()

			if err != nil {
				return nil, err
			}
		}
		if args.UpdateFields != nil {
			for _, inputPageField := range *args.UpdateFields {
				localPageFieldID, err := fromGlobalID(
					"pageField",
					string(inputPageField.ID),
				)
				if err != nil {
					return nil, err
				}

				field := core.PageField{}
				field.ID = localPageFieldID
				field.UpdatedAt = res.
					dataSource.
					GetCurrentTime()

				query := res.
					dataSource.
					DB().
					Model(&field).
					Column("updated_at")

				if inputPageField.Input.Name != nil {
					field.Name = *inputPageField.Input.Name
					query = query.Column("name")
				}
				if inputPageField.Input.Value != nil {
					field.Value = *inputPageField.Input.Value
					query = query.Column("value")
				}

				_, err = query.
					WherePK().
					Update()

				if err != nil {
					return nil, err
				}
			}
		}
		if args.RemoveFields != nil {
			removePageFields := []core.PageField{}
			for _, pageFieldID := range *args.RemoveFields {
				localPageFieldID, err := fromGlobalID("pageField", pageFieldID)
				if err != nil {
					return nil, err
				}

				pageField := core.PageField{}
				pageField.ID = localPageFieldID
				removePageFields = append(removePageFields, pageField)
			}

			_, err = res.
				dataSource.
				DB().
				Model(&removePageFields).
				Delete()

			if err != nil {
				return nil, err
			}
		}
		if args.Input != nil {
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
		}

		_, err = query.
			WherePK().
			Update()

		if err != nil {
			return nil, err
		}
	}

	page.Fields = []core.PageField{}
	err = res.
		dataSource.
		DB().
		Model(&page).
		Where("id = ?", localID).
		Relation("Fields").
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
