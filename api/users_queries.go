package api

import (
	"context"

	"github.com/dominik-zeglen/inkster/core"
	"github.com/dominik-zeglen/inkster/middleware"
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

func (res *Resolver) Viewer(ctx context.Context) (*userResolver, error) {
	viewer, ok := ctx.Value("user").(*middleware.UserClaims)
	if ok {
		user := core.User{}
		user.ID = viewer.ID

		err := res.
			dataSource.
			DB().
			Model(&user).
			WherePK().
			Select()

		return &userResolver{
			data:       &user,
			dataSource: res.dataSource,
		}, err
	}
	return nil, nil
}
