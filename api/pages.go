package api

import (
	"github.com/dominik-zeglen/inkster/core"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

func resolvePages(
	dataSource core.AbstractDataContext,
	sort *Sort,
	paginationInput Paginate,
	where *func(*orm.Query) *orm.Query,
) (*pageConnectionResolver, error) {
	pages := []core.Page{}

	query := dataSource.
		DB().
		Model(&pages)

	if where != nil {
		query = (*where)(query)
	}

	query = sortPages(query, sort)
	query = query.OrderExpr("id ASC")
	query, offset := paginate(query, paginationInput)
	err := query.Select()

	if err != nil {
		if err != pg.ErrNoRows {
			return nil, err
		}
	}

	pageInfo := PageInfo{}

	if paginationInput.First != nil {
		if int(*paginationInput.First) < len(pages) {
			pages = pages[:len(pages)-1]
			pageInfo.hasNextPage = true
		}
	} else if paginationInput.Last != nil {
		if int(*paginationInput.Last) < len(pages) {
			if paginationInput.Before != nil {
				if offset > 0 {
					pageInfo.hasPreviousPage = true
				}
			} else {
				pages = pages[1:]
				pageInfo.hasPreviousPage = true
			}
		}
	}

	if len(pages) > 0 {
		endCursor := Cursor(offset + len(pages) - 1)
		pageInfo.endCursor = &endCursor

		startCursor := Cursor(offset)
		pageInfo.startCursor = &startCursor
	}

	return &pageConnectionResolver{
		dataSource: dataSource,
		data:       pages,
		pageInfo:   pageInfo,
		offset:     offset,
	}, nil
}

func sortPages(query *orm.Query, sort *Sort) *orm.Query {
	if sort != nil {
		orderColumn := "id"
		switch sort.Field {
		case "AUTHOR":
			query = query.Relation("Author")
			orderColumn = "Author.email"

		case "IS_PUBLISHED":
			orderColumn = "is_published"

		case "NAME":
			orderColumn = "name"

		case "SLUG":
			orderColumn = "slug"

		case "UPDATED_AT":
			orderColumn = "updated_at"
		}

		return query.OrderExpr(orderColumn + " " + sort.Order)
	}

	return query
}
