package api

import (
	"context"

	gql "github.com/graph-gophers/graphql-go"
)

type pageArgs struct {
	ID gql.ID
}

func (res *Resolver) Page(
	ctx context.Context,
	args pageArgs,
) (*pageResolver, error) {
	localID, err := fromGlobalID("page", string(args.ID))
	if err != nil {
		return nil, err
	}
	result, err := res.dataSource.GetPage(localID)
	if err != nil {
		return nil, err
	}

	if !result.IsPublished && !checkPermission(ctx) {
		return nil, errNoPermissions
	}

	return &pageResolver{
		dataSource: res.dataSource,
		data:       &result,
	}, nil
}
