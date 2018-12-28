package api

import (
	"context"
	"time"

	"github.com/dominik-zeglen/inkster/core"
	gql "github.com/graph-gophers/graphql-go"
)

type pageResolver struct {
	dataSource core.AbstractDataContext
	data       *core.Page
}

func (res *pageResolver) ID() gql.ID {
	globalID := toGlobalID(gqlPage, res.data.ID)
	return gql.ID(globalID)
}

func (res *pageResolver) Author() (*userResolver, error) {
	if res.data.Author == nil {
		author := core.User{}
		author.ID = res.data.AuthorID

		err := res.
			dataSource.
			DB().
			Model(&author).
			WherePK().
			Select()

		return &userResolver{
			data:       &author,
			dataSource: res.dataSource,
		}, err
	}
	return &userResolver{
		data:       res.data.Author,
		dataSource: res.dataSource,
	}, nil
}

func (res *pageResolver) CreatedAt() string {
	return res.data.CreatedAt.UTC().Format(time.RFC3339)
}

func (res *pageResolver) UpdatedAt() string {
	return res.data.UpdatedAt.UTC().Format(time.RFC3339)
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

func (res *pageResolver) Fields() (*[]*pageFieldResolver, error) {
	var resolverList []*pageFieldResolver
	fields := res.data.Fields
	if fields == nil {
		err := res.
			dataSource.
			DB().
			Model(&fields).
			Where("page_id = ?", res.data.ID).
			OrderExpr("id ASC").
			Select()

		if err != nil {
			return nil, err
		}

		res.data.Fields = fields
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
	return &resolverList, nil
}

func (res *pageResolver) Parent(ctx context.Context) (*directoryResolver, error) {
	parent := core.Directory{}
	parent.ID = res.data.ParentID
	err := res.
		dataSource.
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
	globalID := toGlobalID(gqlPageField, res.data.ID)
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

type pageConnectionResolver struct {
	dataSource core.AbstractDataContext
	data       []core.Page
	pageInfo   PageInfo
	offset     int
}

func (res pageConnectionResolver) Edges() []pageConnectionEdgeResolver {
	resolvers := make([]pageConnectionEdgeResolver, len(res.data))
	for resolverIndex := range resolvers {
		resolvers[resolverIndex] = pageConnectionEdgeResolver{
			dataSource: &res.dataSource,
			data:       res.data[resolverIndex],
			cursor:     pageCursor(resolverIndex + res.offset),
		}
	}
	return resolvers
}

func (res pageConnectionResolver) PageInfo() pageInfoResolver {
	return pageInfoResolver{
		pageInfo: res.pageInfo,
	}
}

type pageConnectionEdgeResolver struct {
	dataSource *core.AbstractDataContext
	data       core.Page
	cursor     pageCursor
}

func (res pageConnectionEdgeResolver) Cursor() string {
	return toGlobalID(gqlCursor, int(res.cursor))
}

func (res pageConnectionEdgeResolver) Node() *pageResolver {
	return &pageResolver{
		dataSource: *res.dataSource,
		data:       &res.data,
	}
}
