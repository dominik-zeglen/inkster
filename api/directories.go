package api

import (
	"github.com/dominik-zeglen/inkster/core"
	"github.com/go-pg/pg/orm"
)

func resolveDirectories(
	dataSource core.AbstractDataContext,
	sort *Sort,
	where *func(*orm.Query) *orm.Query,
) (*[]*directoryResolver, error) {
	directories := []core.Directory{}

	query := dataSource.
		DB().
		Model(&directories)

	if where != nil {
		query = (*where)(query)
	}

	query = sortDirectories(query, sort).OrderExpr("created_at ASC")
	err := query.Select()

	if err != nil {
		return nil, err
	}

	resolvers := make([]*directoryResolver, len(directories))
	for i := range directories {
		resolvers[i] = &directoryResolver{
			dataSource: dataSource,
			data:       &directories[i],
		}
	}

	return &resolvers, nil
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
