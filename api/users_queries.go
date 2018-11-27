package api

import (
	"context"

	"github.com/dominik-zeglen/inkster/core"
	"github.com/go-pg/pg"
	gql "github.com/graph-gophers/graphql-go"
)

type UserQueryArgs struct {
	ID gql.ID
}

func (res *Resolver) User(
	ctx context.Context,
	args UserQueryArgs,
) (*userResolver, error) {
	if !checkPermission(ctx) {
		return nil, errNoPermissions
	}
	localID, err := fromGlobalID("user", string(args.ID))
	if err != nil {
		return nil, nil
	}

	user := core.User{}
	user.ID = localID

	err = res.
		dataSource.
		DB().
		Model(&user).
		WherePK().
		Select()

	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &userResolver{
		dataSource: res.dataSource,
		data:       &user,
	}, nil
}

func (res *Resolver) Users(ctx context.Context) (*[]*userResolver, error) {
	if !checkPermission(ctx) {
		return nil, errNoPermissions
	}
	var resolverList []*userResolver

	users := []core.User{}
	err := res.
		dataSource.
		DB().
		Model(&users).
		Select()

	if err != nil {
		return nil, err
	}
	for index := range users {
		resolverList = append(
			resolverList,
			&userResolver{
				dataSource: res.dataSource,
				data:       &users[index],
			},
		)
	}
	return &resolverList, nil
}
