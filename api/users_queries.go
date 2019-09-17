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

type UsersArgs struct {
	Sort     *Sort
	Paginate PaginationInput
}

func (res *Resolver) Users(
	ctx context.Context,
	args UsersArgs,
) (*userConnectionResolver, error) {
	if !checkPermission(ctx) {
		return nil, errNoPermissions
	}
	return resolveUsers(
		res.dataSource,
		args.Sort,
		getPaginationData(args.Paginate),
		nil,
	)
}

func (res *Resolver) Viewer(ctx context.Context) (*userResolver, error) {
	viewer, ok := ctx.Value(middleware.UserContextKey).(*core.User)
	if ok {
		return &userResolver{
			data:       viewer,
			dataSource: res.dataSource,
		}, nil
	}
	return nil, nil
}
