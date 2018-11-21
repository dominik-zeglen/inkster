package api

import (
	"github.com/dominik-zeglen/inkster/core"
	gql "github.com/graph-gophers/graphql-go"
)

// Type resolvers
type directoryResolver struct {
	dataSource core.AbstractDataContext
	data       *core.Directory
}

func (res *directoryResolver) ID() gql.ID {
	globalID := toGlobalID("directory", res.data.ID)
	return gql.ID(globalID)
}
func (res *directoryResolver) CreatedAt() string {
	return res.data.CreatedAt
}
func (res *directoryResolver) UpdatedAt() string {
	return res.data.UpdatedAt
}
func (res *directoryResolver) Name() string {
	return res.data.Name
}
func (res *directoryResolver) IsPublished() bool {
	return res.data.IsPublished
}
func (res *directoryResolver) Parent() *directoryResolver {
	if res.data.ParentID == 0 {
		return nil
	}
	parent := core.Directory{}
	parent.ID = res.data.ParentID
	err := dataSource.
		DB().
		Model(&parent).
		WherePK().
		Select()

	if err != nil {
		panic(err)
	}
	return &directoryResolver{
		dataSource: res.dataSource,
		data:       &parent,
	}
}
func (res *directoryResolver) Children() *[]*directoryResolver {
	var resolverList []*directoryResolver
	directories := []core.Directory{}

	err := res.
		dataSource.
		DB().
		Model(&directories).
		Where("parent_id = ?", res.data.ID).
		Select()

	if err != nil {
		panic(err)
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
	return &resolverList
}
func (res *directoryResolver) Pages() (*[]*pageResolver, error) {
	var resolverList []*pageResolver
	pages := []core.Page{}

	err := res.
		dataSource.
		DB().
		Model(&pages).
		Where("parent_id = ?", res.data.ID).
		Relation("Fields").
		Select()

	if err != nil {
		return nil, err
	}
	for index := range pages {
		resolverList = append(
			resolverList,
			&pageResolver{
				dataSource: res.dataSource,
				data:       &pages[index],
			},
		)
	}
	return &resolverList, nil
}
