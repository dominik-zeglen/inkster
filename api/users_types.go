package api

import (
	"github.com/dominik-zeglen/inkster/core"
	gql "github.com/graph-gophers/graphql-go"
)

type userResolver struct {
	dataSource core.AbstractDataContext
	data       *core.User
}

func (res *userResolver) ID() gql.ID {
	globalID := toGlobalID("user", res.data.ID)
	return gql.ID(globalID)
}

func (res *userResolver) CreatedAt() string {
	return res.data.CreatedAt
}

func (res *userResolver) UpdatedAt() string {
	return res.data.UpdatedAt
}

func (res *userResolver) Email() string {
	return res.data.Email
}

func (res *userResolver) IsActive() bool {
	return res.data.Active
}

func (res *userResolver) Pages() []*pageResolver {
	if res.data.Pages == nil {
		res.data.Pages = &[]core.Page{}
		err := res.
			dataSource.
			DB().
			Model(res.data.Pages).
			Where("author_id = ?", res.data.ID).
			OrderExpr("id ASC").
			Select()

		if err != nil {
			panic(err)
		}
	}
	resolvers := []*pageResolver{}
	pages := *res.data.Pages
	for index := range pages {
		resolvers = append(resolvers, &pageResolver{
			data:       &pages[index],
			dataSource: res.dataSource,
		})
	}
	return resolvers
}
