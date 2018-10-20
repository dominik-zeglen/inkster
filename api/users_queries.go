package api

import (
	"context"

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
		return nil, err
	}
	result, err := res.dataSource.GetUser(localID)
	if err != nil {
		return nil, err
	}
	return &userResolver{
		dataSource: res.dataSource,
		data:       &result,
	}, nil
}

func (res *Resolver) Users(ctx context.Context) (*[]*userResolver, error) {
	if !checkPermission(ctx) {
		return nil, errNoPermissions
	}
	var resolverList []*userResolver
	result, err := res.dataSource.GetUserList()
	if err != nil {
		return nil, err
	}
	for index := range result {
		resolverList = append(
			resolverList,
			&userResolver{
				dataSource: res.dataSource,
				data:       &result[index],
			},
		)
	}
	return &resolverList, nil
}
