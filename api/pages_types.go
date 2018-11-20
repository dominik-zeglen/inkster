package api

import (
	"context"

	"github.com/dominik-zeglen/inkster/core"
	gql "github.com/graph-gophers/graphql-go"
)

type pageResolver struct {
	dataSource core.AbstractDataContext
	data       *core.Page
}

func (res *pageResolver) ID() gql.ID {
	globalID := toGlobalID("page", res.data.ID)
	return gql.ID(globalID)
}

func (res *pageResolver) CreatedAt() string {
	return res.data.CreatedAt
}

func (res *pageResolver) UpdatedAt() string {
	return res.data.UpdatedAt
}

func (res *pageResolver) Name() string {
	return res.data.Name
}

func (res *pageResolver) Slug() string {
	return res.data.Slug
}

func (res *pageResolver) IsPublished() bool {
	return res.data.IsPublished
}

func (res *pageResolver) Fields() *[]*pageFieldResolver {
	var resolverList []*pageFieldResolver
	fields := res.data.Fields
	if fields == nil {
		return nil
	}
	for i := range fields {
		resolverList = append(
			resolverList,
			&pageFieldResolver{
				dataSource: res.dataSource,
				data:       &fields[i],
			},
		)
	}
	return &resolverList
}

func (res *pageResolver) Parent(ctx context.Context) (*directoryResolver, error) {
	parent := core.Directory{}
	parent.ID = res.data.ParentID
	err := dataSource.
		DB().
		Model(&parent).
		WherePK().
		Select()

	if err != nil {
		return nil, err
	}

	if !parent.IsPublished && !checkPermission(ctx) {
		return nil, errNoPermissions
	}

	return &directoryResolver{
		dataSource: res.dataSource,
		data:       &parent,
	}, nil
}

type pageFieldResolver struct {
	dataSource core.AbstractDataContext
	data       *core.PageField
}

func (res *pageFieldResolver) ID() gql.ID {
	globalID := toGlobalID("pageField", res.data.ID)
	return gql.ID(globalID)
}

func (res *pageFieldResolver) Name() string {
	return res.data.Name
}

func (res *pageFieldResolver) Type() string {
	return res.data.Type
}

func (res *pageFieldResolver) Value() *string {
	return &res.data.Value
}
