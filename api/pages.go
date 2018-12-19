package api

import (
	"github.com/dominik-zeglen/inkster/core"
	"github.com/go-pg/pg/orm"
)

type PageSort struct {
	Field string
	Order string
}

func resolvePages(
	dataSource core.AbstractDataContext,
	sort *PageSort,
	where *func(*orm.Query) *orm.Query,
) (*[]*pageResolver, error) {
	pages := []core.Page{}

	query := dataSource.
		DB().
		Model(&pages)

	if where != nil {
		query = (*where)(query)
	}

	query = sortPages(query, sort)
	err := query.Select()

	if err != nil {
		return nil, err
	}

	resolvers := make([]*pageResolver, len(pages))
	for i := range pages {
		resolvers[i] = &pageResolver{
			dataSource: dataSource,
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
