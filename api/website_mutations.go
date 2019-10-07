package api

import (
	"context"

	"github.com/dominik-zeglen/inkster/core"
)

type websiteOperationResult struct {
	website          *core.Website
	validationErrors []core.ValidationError
}

type websiteOperationResultResolver struct {
	dataSource core.AbstractDataContext
	data       websiteOperationResult
}

func (res *websiteOperationResultResolver) Errors() []inputErrorResolver {
	return createInputErrorResolvers(res.data.validationErrors)
}

func (res *websiteOperationResultResolver) Website() *websiteResolver {
	return &websiteResolver{
		data: res.data.website,
	}
}

type websiteUpdateInput struct {
	Name        *string
	Description *string
	Domain      *string
}
type updateWebsiteArgs struct {
	Input websiteUpdateInput
}

func (args updateWebsiteArgs) validate(dataSource core.AbstractDataContext) (
	[]core.ValidationError,
	*core.Website,
	error,
) {
	errors := []core.ValidationError{}

	website := core.Website{}
	website.ID = core.WEBSITE_DB_ID

	err := dataSource.
		DB().
		Model(&website).
		WherePK().
		Select()

	if err != nil {
		return errors, nil, err
	}

	cleanedWebsite := website
	if args.Input.Name != nil {
		cleanedWebsite.Name = *args.Input.Name
	}
	if args.Input.Description != nil {
		cleanedWebsite.Description = *args.Input.Description
	}
	if args.Input.Domain != nil {
		cleanedWebsite.Domain = *args.Input.Domain
	}

	errors = append(errors, cleanedWebsite.Validate()...)
	if len(errors) > 0 {
		return errors, &website, err
	}
	return errors, &cleanedWebsite, err
}

func (res *Resolver) UpdateWebsite(
	ctx context.Context,
	args updateWebsiteArgs,
) (*websiteOperationResultResolver, error) {
	if !checkPermission(ctx) {
		return nil, errNoPermissions
	}

	validationErrors, website, err := args.validate(res.dataSource)
	if err != nil {
		return nil, err
	}

	if len(validationErrors) > 0 {
		return &websiteOperationResultResolver{
			dataSource: res.dataSource,
			data: websiteOperationResult{
				validationErrors: validationErrors,
				website:          website,
			},
		}, nil
	}

	err = res.dataSource.DB().Update(website)

	return &websiteOperationResultResolver{
		dataSource: res.dataSource,
		data: websiteOperationResult{
			validationErrors: validationErrors,
			website:          website,
		},
	}, err
}
