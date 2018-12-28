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

type PaginationInput struct {
	After  *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	First  *int32  `json:"first,omitempty"`
	Last   *int32  `json:"last,omitempty"`
}
type PagesArgs struct {
	Sort     *Sort
	Paginate PaginationInput
}

func (res *Resolver) Pages(
	ctx context.Context,
	args PagesArgs,
) (*pageConnectionResolver, error) {
	paginate := Paginate{}

	if args.Paginate.First != nil {
		paginate.First = args.Paginate.First
	}

	if args.Paginate.Last != nil {
		paginate.Last = args.Paginate.Last
	}

	if args.Paginate.After != nil {
		cursor, err := fromGlobalCursor(*args.Paginate.After)
		if err == nil {
			paginate.After = &cursor
		}
	}

	if args.Paginate.Before != nil {
		cursor, err := fromGlobalCursor(*args.Paginate.Before)
		if err == nil {
			paginate.Before = &cursor
		}
	}

	return resolvePages(
		res.dataSource,
		args.Sort,
		paginate,
		nil,
	)
}
