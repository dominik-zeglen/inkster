package api

import (
	"github.com/dominik-zeglen/inkster/core"
	"github.com/go-pg/pg/orm"
)

func resolveUsers(
	dataSource core.AbstractDataContext,
	sort *Sort,
	where *func(*orm.Query) *orm.Query,
) (*[]*userResolver, error) {
	users := []core.User{}

	query := dataSource.
		DB().
		Model(&users)

	if where != nil {
		query = (*where)(query)
	}

	query = sortUsers(query, sort).OrderExpr("created_at ASC")
	err := query.Select()

	if err != nil {
		return nil, err
	}

	resolvers := make([]*userResolver, len(users))
	for i := range users {
		resolvers[i] = &userResolver{
			dataSource: dataSource,
			data:       &users[i],
		}
	}

	return &resolvers, nil
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
