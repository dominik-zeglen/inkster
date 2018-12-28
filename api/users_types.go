package api

import (
	"context"
	"time"

	"github.com/dominik-zeglen/inkster/core"
	"github.com/go-pg/pg/orm"
	gql "github.com/graph-gophers/graphql-go"
)

type userResolver struct {
	dataSource core.AbstractDataContext
	data       *core.User
}

func (res *userResolver) ID() gql.ID {
	globalID := toGlobalID(gqlUser, res.data.ID)
	return gql.ID(globalID)
}

func (res *userResolver) CreatedAt() string {
	return res.data.CreatedAt.UTC().Format(time.RFC3339)
}

func (res *userResolver) UpdatedAt() string {
	return res.data.UpdatedAt.UTC().Format(time.RFC3339)
}

func (res *userResolver) Email() string {
	return res.data.Email
}

func (res *userResolver) IsActive() bool {
	return res.data.Active
}

type UserPagesArgs struct {
	Sort     *Sort
	Paginate Paginate
}

func (res *userResolver) Pages(
	ctx context.Context,
	args UserPagesArgs,
) (*pageConnectionResolver, error) {
	where := func(query *orm.Query) *orm.Query {
		return query.Where("author_id = ?", res.data.ID)
	}
	return resolvePages(res.dataSource, args.Sort, args.Paginate, &where)
}

type userConnectionResolver struct {
	dataSource core.AbstractDataContext
	data       []core.User
	pageInfo   PageInfo
	offset     int
}

func (res userConnectionResolver) Edges() []userConnectionEdgeResolver {
	resolvers := make([]userConnectionEdgeResolver, len(res.data))
	for resolverIndex := range resolvers {
		resolvers[resolverIndex] = userConnectionEdgeResolver{
			dataSource: &res.dataSource,
			data:       res.data[resolverIndex],
			cursor:     pageCursor(resolverIndex + res.offset),
		}
	}
	return resolvers
}

func (res userConnectionResolver) PageInfo() pageInfoResolver {
	return pageInfoResolver{
		pageInfo: res.pageInfo,
	}
}

type userConnectionEdgeResolver struct {
	dataSource *core.AbstractDataContext
	data       core.User
	cursor     pageCursor
}

func (res userConnectionEdgeResolver) Cursor() string {
	return toGlobalID(gqlCursor, int(res.cursor))
}

func (res userConnectionEdgeResolver) Node() *userResolver {
	return &userResolver{
		dataSource: *res.dataSource,
		data:       &res.data,
	}
}
