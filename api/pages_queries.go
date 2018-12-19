package api

import (
	"context"

	"github.com/dominik-zeglen/inkster/core"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	gql "github.com/graph-gophers/graphql-go"
)

type pageArgs struct {
	ID gql.ID
}

func (res *Resolver) Page(
	ctx context.Context,
	args pageArgs,
) (*pageResolver, error) {
	localID, err := fromGlobalID("page", string(args.ID))
	if err != nil {
		return nil, nil
	}

	page := core.Page{}

	err = res.
		dataSource.
		DB().
		Model(&page).
		Where("page.id = ?", localID).
		Relation("Fields", func(query *orm.Query) (*orm.Query, error) {
			return query.OrderExpr("id ASC"), nil
		}).
		Relation("Author").
		Select()

	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err

	}

	if !page.IsPublished && !checkPermission(ctx) {
		return nil, errNoPermissions
	}

	return &pageResolver{
		dataSource: res.dataSource,
		data:       &page,
	}, nil
}

type PageSort struct {
	Field string
	Order string
}
type PagesArgs struct {
	Sort *PageSort
}

func (res *Resolver) Pages(
	ctx context.Context,
	args PagesArgs,
) (*[]*pageResolver, error) {
	pages := []core.Page{}

	query := res.
		dataSource.
		DB().
		Model(&pages)
	query = sortPages(query, args.Sort)
	err := query.Select()

	if err != nil {
		return nil, err
	}

	resolvers := make([]*pageResolver, len(pages))
	for i := range pages {
		resolvers[i] = &pageResolver{
			dataSource: res.dataSource,
			data:       &pages[i],
		}
	}

	return &resolvers, nil
}

func sortPages(query *orm.Query, sort *PageSort) *orm.Query {
	if sort != nil {
		switch sort.Field {
		case "AUTHOR":
			query = query.Relation("Author")
			return query.OrderExpr("Author.email " + sort.Order)

		case "IS_PUBLISHED":
			return query.OrderExpr("is_published " + sort.Order)

		case "NAME":
			return query.OrderExpr("name " + sort.Order)

		case "SLUG":
			return query.OrderExpr("slug " + sort.Order)

		case "UPDATED_AT":
			return query.OrderExpr("updated_at " + sort.Order)
		}

	}

	return query.OrderExpr("created_at ASC")
}
