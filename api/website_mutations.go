package api

import (
	"context"

	"github.com/dominik-zeglen/inkster/core"
	"github.com/dominik-zeglen/inkster/middleware"
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
	if res.data.website == nil {
		return nil
	}

	return &websiteResolver{
		data: *res.data.website,
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

func (res *Resolver) UpdateWebsite(
	ctx context.Context,
	args updateWebsiteArgs,
) (*websiteOperationResultResolver, error) {
	if !checkPermission(ctx) {
		return nil, errNoPermissions
	}

	website := ctx.Value(middleware.WebsiteContextKey).(core.Website)

	if args.Input.Name != nil {
		website.Name = *args.Input.Name
	}
	if args.Input.Description != nil {
		website.Description = *args.Input.Description
	}
	if args.Input.Domain != nil {
		website.Domain = *args.Input.Domain
	}

	updatedWebsite, validationErrors, err := core.UpdateWebsite(
		website,
		res.dataSource,
	)

	return &websiteOperationResultResolver{
		data: websiteOperationResult{
			website:          updatedWebsite,
			validationErrors: validationErrors,
		},
		dataSource: res.dataSource,
	}, err
}
