package api

import (
	"context"

	"github.com/dominik-zeglen/inkster/core"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
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
		return nil, nil
	}

	directory := core.Directory{}
	directory.ID = localID
	err = res.
		dataSource.
		DB().
		Model(&directory).
		WherePK().
		Select()

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

type DirectoriesArgs struct {
	Sort *Sort
}

func (res *Resolver) GetDirectories(
	args DirectoriesArgs,
) (*[]*directoryResolver, error) {
	return resolveDirectories(res.dataSource, args.Sort, nil)
}

type RootDirectoriesArgs struct {
	Sort *Sort
}

func (res *Resolver) GetRootDirectories(
	args RootDirectoriesArgs,
) (*[]*directoryResolver, error) {
	where := func(query *orm.Query) *orm.Query {
		return query.Where("parent_id IS NULL OR parent_id = 0")
	}
	return resolveDirectories(res.dataSource, args.Sort, &where)
}
