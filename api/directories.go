package api

import (
	"github.com/dominik-zeglen/inkster/core"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

func resolveDirectories(
	dataSource core.AbstractDataContext,
	sort *Sort,
	paginationInput Paginate,
	where *func(*orm.Query) *orm.Query,
) (*directoryConnectionResolver, error) {
	directories := []core.Directory{}

	query := dataSource.
		DB().
		Model(&directories)

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
		if int(*paginationInput.First) < len(directories) {
			directories = directories[:len(directories)-1]
			pageInfo.hasNextPage = true
		}
	} else if paginationInput.Last != nil {
		if int(*paginationInput.Last) < len(directories) {
			if paginationInput.Before != nil {
				if offset > 0 {
					pageInfo.hasPreviousPage = true
				}
			} else {
				directories = directories[1:]
				pageInfo.hasPreviousPage = true
			}
		}
	}

	if len(directories) > 0 {
		endCursor := Cursor(offset + len(directories) - 1)
		pageInfo.endCursor = &endCursor

		startCursor := Cursor(offset)
		pageInfo.startCursor = &startCursor
	}

	return &directoryConnectionResolver{
		dataSource: dataSource,
		data:       directories,
		pageInfo:   pageInfo,
	}, nil
}

func sortDirectories(query *orm.Query, sort *Sort) *orm.Query {
	if sort != nil {
		switch sort.Field {
		case "CREATED_AT":
			return query.OrderExpr("created_at " + sort.Order)

		case "IS_PUBLISHED":
			return query.OrderExpr("is_published " + sort.Order)

		case "NAME":
			return query.OrderExpr("name " + sort.Order)

		case "UPDATED_AT":
			return query.OrderExpr("updated_at " + sort.Order)
		}

	}

	return query.OrderExpr("created_at ASC")
}
