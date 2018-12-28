package api

import (
	"context"
	"time"

	"github.com/dominik-zeglen/inkster/core"
	"github.com/go-pg/pg/orm"
	gql "github.com/graph-gophers/graphql-go"
)

// Type resolvers
type directoryResolver struct {
	dataSource core.AbstractDataContext
	data       *core.Directory
}

func (res *directoryResolver) ID() gql.ID {
	globalID := toGlobalID(gqlDirectory, res.data.ID)
	return gql.ID(globalID)
}
func (res *directoryResolver) CreatedAt() string {
	return res.data.CreatedAt.UTC().Format(time.RFC3339)
}
func (res *directoryResolver) UpdatedAt() string {
	return res.data.UpdatedAt.UTC().Format(time.RFC3339)
}
func (res *directoryResolver) Name() string {
	return res.data.Name
}
func (res *directoryResolver) IsPublished() bool {
	return res.data.IsPublished
}
func (res *directoryResolver) Parent() *directoryResolver {
	if res.data.ParentID == 0 {
		return nil
	}
	parent := core.Directory{}
	parent.ID = res.data.ParentID
	err := res.
		dataSource.
		DB().
		Model(&parent).
		WherePK().
		Select()

	if err != nil {
		panic(err)
	}
	return &directoryResolver{
		dataSource: res.dataSource,
		data:       &parent,
	}
}

type DirectoryChildrenArgs struct {
	Sort     *Sort
	Paginate PaginationInput
}

func (res *directoryResolver) Children(
	args DirectoryChildrenArgs,
) (*directoryConnectionResolver, error) {
	where := func(query *orm.Query) *orm.Query {
		return query.Where("parent_id = ?", res.data.ID)
	}
	return resolveDirectories(
		res.dataSource,
		args.Sort,
		getPaginationData(args.Paginate),
		&where,
	)
}

type DirectoryPagesArgs struct {
	Sort     *Sort
	Paginate Paginate
}

func (res *directoryResolver) Pages(
	ctx context.Context,
	args DirectoryPagesArgs,
) (*pageConnectionResolver, error) {
	where := func(query *orm.Query) *orm.Query {
		return query.Where("parent_id = ?", res.data.ID)
	}
	return resolvePages(res.dataSource, args.Sort, args.Paginate, &where)
}

type directoryConnectionResolver struct {
	dataSource core.AbstractDataContext
	data       []core.Directory
	pageInfo   PageInfo
	offset     int
}

func (res directoryConnectionResolver) Edges() []directoryConnectionEdgeResolver {
	resolvers := make([]directoryConnectionEdgeResolver, len(res.data))
	for resolverIndex := range resolvers {
		resolvers[resolverIndex] = directoryConnectionEdgeResolver{
			dataSource: &res.dataSource,
			data:       res.data[resolverIndex],
			cursor:     pageCursor(resolverIndex + res.offset),
		}
	}
	return resolvers
}

func (res directoryConnectionResolver) PageInfo() pageInfoResolver {
	return pageInfoResolver{
		pageInfo: res.pageInfo,
	}
}

type directoryConnectionEdgeResolver struct {
	dataSource *core.AbstractDataContext
	data       core.Directory
	cursor     pageCursor
}

func (res directoryConnectionEdgeResolver) Cursor() string {
	return toGlobalID(gqlCursor, int(res.cursor))
}

func (res directoryConnectionEdgeResolver) Node() *directoryResolver {
	return &directoryResolver{
		dataSource: *res.dataSource,
		data:       &res.data,
	}
}
