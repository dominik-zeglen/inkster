package api

import (
	"context"

	"github.com/go-pg/pg"
	gql "github.com/graph-gophers/graphql-go"
)

type getDirectoryArgs struct {
	Id gql.ID
}

func (res *Resolver) GetDirectory(
	ctx context.Context,
	args getDirectoryArgs,
) (*directoryResolver, error) {
	localID, err := fromGlobalID("directory", string(args.Id))
	if err != nil {
		return nil, err
	}
	directory, err := res.
		dataSource.
		GetDirectory(localID)
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	if !directory.IsPublished && !checkPermission(ctx) {
		return nil, errNoPermissions
	}
	return &directoryResolver{
		dataSource: res.dataSource,
		data:       &directory,
	}, nil
}

func (res *Resolver) GetDirectories() (*[]*directoryResolver, error) {
	var resolverList []*directoryResolver
	directories, err := res.dataSource.GetDirectoryList()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	for index := range directories {
		resolverList = append(
			resolverList,
			&directoryResolver{
				dataSource: res.dataSource,
				data:       &directories[index],
			},
		)
	}
	return &resolverList, nil
}

func (res *Resolver) GetRootDirectories() (*[]*directoryResolver, error) {
	var resolverList []*directoryResolver
	directories, err := res.dataSource.GetRootDirectoryList()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	for index := range directories {
		resolverList = append(
			resolverList,
			&directoryResolver{
				dataSource: res.dataSource,
				data:       &directories[index],
			},
		)
	}
	return &resolverList, nil
}
