package api

import (
	"context"

	"github.com/dominik-zeglen/inkster/core"
	"github.com/globalsign/mgo/bson"
	gql "github.com/graph-gophers/graphql-go"
)

type pageCreateResult struct {
	userErrors *[]userError
	page       core.Page
}

type pageCreateResultResolver struct {
	dataSource core.Adapter
	data       pageCreateResult
}

func (res *pageCreateResultResolver) Page() *pageResolver {
	return &pageResolver{
		dataSource: res.dataSource,
		data:       &res.data.page,
	}
}
func (res *pageCreateResultResolver) UserErrors() *[]*userErrorResolver {
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

type createPageArgs struct {
	Input struct {
		Name        string
		ParentID    string
		IsPublished *bool
		Fields      *[]*struct {
			Name  string
			Type  string
			Value string
		}
	}
}

func (res *Resolver) CreatePage(
	ctx context.Context,
	args createPageArgs,
) (*pageCreateResultResolver, error) {
	if !checkPermission(ctx) {
		return nil, errNoPermissions
	}

	if args.Input.Name == "" {
		return &pageCreateResultResolver{
			dataSource: res.dataSource,
			data: pageCreateResult{
				userErrors: &[]userError{
					userError{
						field:   "name",
						message: errNoEmpty("name").Error(),
					},
				},
			},
		}, nil
	}
	if args.Input.ParentID == "" {
		return &pageCreateResultResolver{
			dataSource: res.dataSource,
			data: pageCreateResult{
				userErrors: &[]userError{
					userError{
						field:   "parentId",
						message: errNoEmpty("parentId").Error(),
					},
				},
			},
		}, nil
	}
	localID, err := fromGlobalID("directory", args.Input.ParentID)
	if err != nil {
		return nil, err
	}
	page := core.Page{
		Name:     args.Input.Name,
		ParentID: bson.ObjectId(localID),
	}
	if args.Input.IsPublished != nil {
		page.IsPublished = *args.Input.IsPublished
	}
	if args.Input.Fields != nil {
		fields := *args.Input.Fields
		page.Fields = make([]core.PageField, len(fields))
		for i := range fields {
			page.Fields[i] = core.PageField{
				Name:  fields[i].Name,
				Type:  fields[i].Type,
				Value: fields[i].Value,
			}
		}
	}
	result, err := res.dataSource.AddPage(page)
	if err != nil {
		return nil, err
	}
	return &pageCreateResultResolver{
		dataSource: res.dataSource,
		data: pageCreateResult{
			page: result,
		},
	}, nil
}

type PageField2Input struct {
	Name   string
	Update struct {
		Name  *string
		Value *string
	}
}
type updatePageArgs struct {
	ID    gql.ID
	Input *struct {
		Name        *string
		Slug        *string
		ParentID    *string
		IsPublished *bool
		Fields      *[]PageField2Input
	}
	AddFields    *[]core.PageField
	RemoveFields *[]string
}

func (res *Resolver) UpdatePage(
	ctx context.Context,
	args updatePageArgs,
) (*pageCreateResultResolver, error) {
	if !checkPermission(ctx) {
		return nil, errNoPermissions
	}

	localID, err := fromGlobalID("page", string(args.ID))
	if err != nil {
		return nil, err
	}
	if args.Input != nil || args.AddFields != nil || args.RemoveFields != nil {
		page, err := res.dataSource.GetPage(localID)
		if err != nil {
			return nil, err
		}
		pageInput := core.PageInput{}
		if args.Input != nil {
			if args.Input.Name != nil {
				if *args.Input.Name == "" {
					return &pageCreateResultResolver{
						dataSource: res.dataSource,
						data: pageCreateResult{
							userErrors: &[]userError{
								userError{
									field:   "name",
									message: errNoEmpty("name").Error(),
								},
							},
						},
					}, nil
				}
				pageInput.Name = args.Input.Name
			}
			if args.Input.Slug != nil {
				if *args.Input.Slug == "" {
					return &pageCreateResultResolver{
						dataSource: res.dataSource,
						data: pageCreateResult{
							userErrors: &[]userError{
								userError{
									field:   "slug",
									message: errNoEmpty("slug").Error(),
								},
							},
						},
					}, nil
				}
				pageInput.Slug = args.Input.Slug
			}
			if args.Input.ParentID != nil {
				parentID := args.Input.ParentID
				parentObjectID := bson.ObjectId(*parentID)
				pageInput.ParentID = &parentObjectID
			}
			if args.Input.IsPublished != nil {
				pageInput.IsPublished = args.Input.IsPublished
			}
			if args.Input.Fields != nil {
				fields := *args.Input.Fields
				for inputFieldIndex := range fields {
					found := false
					for pageFieldIndex := range page.Fields {
						if page.Fields[pageFieldIndex].Name == fields[inputFieldIndex].Name {
							found = true
							if fields[inputFieldIndex].Update.Name != nil {
								page.Fields[pageFieldIndex].Name = *fields[inputFieldIndex].Update.Name
							}
							if fields[inputFieldIndex].Update.Value != nil {
								page.Fields[pageFieldIndex].Value = *fields[inputFieldIndex].Update.Value
							}
						}
					}
					if !found {
						return &pageCreateResultResolver{
							dataSource: res.dataSource,
							data: pageCreateResult{
								userErrors: &[]userError{
									userError{
										field:   fields[inputFieldIndex].Name,
										message: errNoEmpty(fields[inputFieldIndex].Name).Error(),
									},
								},
							},
						}, nil
					}
				}
				pageInput.Fields = &page.Fields
			}
		}
		if args.AddFields != nil {
			if pageInput.Fields == nil {
				fields := make([]core.PageField, len(page.Fields))
				copy(fields, page.Fields)
				pageInput.Fields = &fields
			}
			tmp := append(*pageInput.Fields, *args.AddFields...)
			pageInput.Fields = &tmp
		}
		if args.RemoveFields != nil {
			if pageInput.Fields == nil {
				fields := make([]core.PageField, len(page.Fields))
				copy(fields, page.Fields)
				pageInput.Fields = &fields
			}
			fields := *pageInput.Fields
			removeFields := *args.RemoveFields
			for inputFieldIndex := range removeFields {
				found := false
				for pageFieldIndex := range page.Fields {
					if fields[pageFieldIndex].Name == removeFields[inputFieldIndex] {
						found = true
					}
				}
				if !found {
					return &pageCreateResultResolver{
						dataSource: res.dataSource,
						data: pageCreateResult{
							userErrors: &[]userError{
								userError{
									field:   fields[inputFieldIndex].Name,
									message: errNoEmpty(fields[inputFieldIndex].Name).Error(),
								},
							},
						},
					}, nil
				}
			}
			fields = []core.PageField{}
			pageFields := *pageInput.Fields
			for pageFieldIndex := range pageFields {
				found := false
				for inputFieldIndex := range removeFields {
					if pageFields[pageFieldIndex].Name == removeFields[inputFieldIndex] {
						found = true
					}
				}
				if !found {
					fields = append(fields, pageFields[pageFieldIndex])
				}
			}
			pageInput.Fields = &fields
		}
		err = res.dataSource.UpdatePage(localID, pageInput)
		if err != nil {
			return nil, err
		}
	}
	page, err := res.dataSource.GetPage(localID)
	if err != nil {
		return nil, err
	}
	return &pageCreateResultResolver{
		dataSource: res.dataSource,
		data: pageCreateResult{
			page: page,
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
