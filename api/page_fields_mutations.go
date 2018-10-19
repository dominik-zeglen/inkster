package api

import (
	"context"

	"github.com/dominik-zeglen/inkster/core"
	"github.com/globalsign/mgo/bson"
	gql "github.com/graph-gophers/graphql-go"
)

type pageFieldOperationResult struct {
	userErrors *[]userError
	pageID     bson.ObjectId
	page       *core.Page
}
type pageFieldOperationResultResolver struct {
	dataSource core.Adapter
	data       pageFieldOperationResult
}

func (res *pageFieldOperationResultResolver) Page(ctx context.Context) (*pageResolver, error) {
	result, err := res.dataSource.GetPage(res.data.pageID)
	if err != nil {
		return nil, err
	}

	return &pageResolver{
		dataSource: res.dataSource,
		data:       &result,
	}, nil
}

func (res *pageFieldOperationResultResolver) UserErrors() *[]*userErrorResolver {
	var resolverList []*userErrorResolver
	if res.data.userErrors == nil {
		return nil
	}
	userErrors := *res.data.userErrors
	for i := range userErrors {
		resolverList = append(
			resolverList,
			&userErrorResolver{
				data: userErrors[i],
			},
		)
	}
	return &resolverList
}

type createPageFieldArgs struct {
	ID    gql.ID
	Input struct {
		Name  string
		Type  string
		Value string
	}
}

func (res *Resolver) CreatePageField(
	ctx context.Context,
	args createPageFieldArgs,
) (*pageFieldOperationResultResolver, error) {
	if !checkPermission(ctx) {
		return nil, errNoPermissions
	}

	localID, err := fromGlobalID("page", string(args.ID))
	if err != nil {
		return nil, err
	}
	if len(args.Input.Name) == 0 {
		return &pageFieldOperationResultResolver{
			dataSource: res.dataSource,
			data: pageFieldOperationResult{
				userErrors: &[]userError{
					userError{
						field:   "name",
						message: errNoEmpty("name").Error(),
					},
				},
			},
		}, nil
	}
	if len(args.Input.Type) == 0 {
		return &pageFieldOperationResultResolver{
			dataSource: res.dataSource,
			data: pageFieldOperationResult{
				userErrors: &[]userError{
					userError{
						field:   "type",
						message: errNoEmpty("type").Error(),
					},
				},
			},
		}, nil
	}
	field := core.PageField{
		Name:  args.Input.Name,
		Type:  args.Input.Type,
		Value: args.Input.Value,
	}
	err = res.dataSource.AddPageField(localID, field)
	if err != nil {
		return nil, err
	}
	return &pageFieldOperationResultResolver{
		dataSource: res.dataSource,
		data: pageFieldOperationResult{
			pageID: localID,
		},
	}, nil
}

type renamePageFieldArgs struct {
	ID    gql.ID
	Input struct {
		Name     string
		RenameTo string
	}
}

func (res *Resolver) RenamePageField(
	ctx context.Context,
	args renamePageFieldArgs,
) (*pageFieldOperationResultResolver, error) {
	if !checkPermission(ctx) {
		return nil, errNoPermissions
	}

	localID, err := fromGlobalID("page", string(args.ID))
	if err != nil {
		return nil, err
	}
	if len(args.Input.RenameTo) == 0 {
		return &pageFieldOperationResultResolver{
			dataSource: res.dataSource,
			data: pageFieldOperationResult{
				userErrors: &[]userError{
					userError{
						field:   "renameTo",
						message: errNoEmpty("renameTo").Error(),
					},
				},
				pageID: localID,
			},
		}, nil
	}
	page, err := res.dataSource.GetPage(localID)
	if err != nil {
		return nil, err
	}
	found := false
	var field core.PageField
	for fieldIndex := range page.Fields {
		if page.Fields[fieldIndex].Name == args.Input.Name {
			found = true
			field = page.Fields[fieldIndex]
		}
	}
	if !found {
		return nil, core.ErrNoField(args.Input.Name)
	}
	field.Name = args.Input.RenameTo
	err = res.dataSource.AddPageField(localID, field)
	if err != nil {
		return nil, err
	}
	err = res.dataSource.RemovePageField(localID, args.Input.Name)
	if err != nil {
		return nil, err
	}
	return &pageFieldOperationResultResolver{
		dataSource: res.dataSource,
		data: pageFieldOperationResult{
			pageID: localID,
		},
	}, nil
}

type updatePageFieldArgs struct {
	ID    gql.ID
	Input struct {
		Name  string
		Value string
	}
}

func (res *Resolver) UpdatePageField(
	ctx context.Context,
	args updatePageFieldArgs,
) (*pageFieldOperationResultResolver, error) {
	if !checkPermission(ctx) {
		return nil, errNoPermissions
	}

	localID, err := fromGlobalID("page", string(args.ID))
	if err != nil {
		return nil, err
	}
	err = res.dataSource.UpdatePageField(localID, args.Input.Name, args.Input.Value)
	if err != nil {
		return nil, err
	}
	return &pageFieldOperationResultResolver{
		dataSource: res.dataSource,
		data: pageFieldOperationResult{
			pageID: localID,
		},
	}, nil
}

type removePageFieldArgs struct {
	ID    gql.ID
	Input struct {
		Name string
	}
}

func (res *Resolver) RemovePageField(
	ctx context.Context,
	args removePageFieldArgs,
) (*pageFieldOperationResultResolver, error) {
	if !checkPermission(ctx) {
		return nil, errNoPermissions
	}

	localID, err := fromGlobalID("page", string(args.ID))
	if err != nil {
		return nil, err
	}
	err = res.dataSource.RemovePageField(localID, args.Input.Name)
	if err != nil {
		return nil, err
	}
	return &pageFieldOperationResultResolver{
		dataSource: res.dataSource,
		data: pageFieldOperationResult{
			pageID: localID,
		},
	}, nil
}
