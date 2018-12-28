package api

import (
	"github.com/dominik-zeglen/inkster/core"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

func resolveUsers(
	dataSource core.AbstractDataContext,
	sort *Sort,
	paginationInput Paginate,
	where *func(*orm.Query) *orm.Query,
) (*userConnectionResolver, error) {
	users := []core.User{}

	query := dataSource.
		DB().
		Model(&users)

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
		if int(*paginationInput.First) < len(users) {
			users = users[:len(users)-1]
			pageInfo.hasNextPage = true
		}
	} else if paginationInput.Last != nil {
		if int(*paginationInput.Last) < len(users) {
			if paginationInput.Before != nil {
				if offset > 0 {
					pageInfo.hasPreviousPage = true
				}
			} else {
				users = users[1:]
				pageInfo.hasPreviousPage = true
			}
		}
	}

	if len(users) > 0 {
		endCursor := Cursor(offset + len(users) - 1)
		pageInfo.endCursor = &endCursor

		startCursor := Cursor(offset)
		pageInfo.startCursor = &startCursor
	}

	return &userConnectionResolver{
		dataSource: dataSource,
		data:       users,
		pageInfo:   pageInfo,
	}, nil
}

func sortUsers(query *orm.Query, sort *Sort) *orm.Query {
	if sort != nil {
		switch sort.Field {
		case "ACTIVE":
			return query.OrderExpr("active " + sort.Order)

		case "CREATED_AT":
			return query.OrderExpr("created_at " + sort.Order)

		case "EMAIL":
			return query.OrderExpr("email " + sort.Order)

		case "UPDATED_AT":
			return query.OrderExpr("updated_at " + sort.Order)
		}
	}

	return query.OrderExpr("created_at ASC")
}
