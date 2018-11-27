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
