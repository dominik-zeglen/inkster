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

type pagesArgs struct{}

func (res *Resolver) Pages(
	ctx context.Context,
) (*[]*pageResolver, error) {
	pages := []core.Page{}

	err := res.
		dataSource.
		DB().
		Model(&pages).
		OrderExpr("id ASC").
		Relation("Fields", func(query *orm.Query) (*orm.Query, error) {
			return query.OrderExpr("id ASC"), nil
		}).
		Relation("Author").
		Select()
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
