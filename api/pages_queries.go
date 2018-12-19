package api

import (
	"context"

	"github.com/dominik-zeglen/inkster/core"
	"github.com/go-pg/pg"
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
	page.ID = localID

	err = res.
		dataSource.
		DB().
		Model(&page).
		WherePK().
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

type PagesArgs struct {
	Sort *Sort
}

func (res *Resolver) Pages(
	ctx context.Context,
	args PagesArgs,
) (*[]*pageResolver, error) {
	return resolvePages(res.dataSource, args.Sort, nil)
}
