package postgres

import (
	"github.com/dominik-zeglen/ecoknow/core"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// AddContainer puts container in a pg database
func (adapter *Adapter) AddContainer(container core.Container) (core.Container, error) {
	db := pg.Connect(&adapter.ConnectionOptions)
	err := db.Insert(&container)
	defer db.Close()
	return container, err
}

// GetContainer gets container from a pg database
func (adapter *Adapter) GetContainer(id int32) (core.Container, error) {
	db := pg.Connect(&adapter.ConnectionOptions)
	defer db.Close()
	container := core.Container{ID: id}
	err := db.Select(&container)
	return container, err
}

// GetContainerList gets all containers from a pg database
func (adapter *Adapter) GetContainerList() ([]core.Container, error) {
	db := pg.Connect(&adapter.ConnectionOptions)
	defer db.Close()
	var containers []core.Container
	err := db.Model(&containers).Select()
	if err == pg.ErrNoRows {
		return containers, nil
	}
	return containers, err
}

// GetRootContainerList gets only containers from a pg database that don't have parent
func (adapter *Adapter) GetRootContainerList() ([]core.Container, error) {
	db := pg.Connect(&adapter.ConnectionOptions)
	defer db.Close()
	getRootContainerFilter := func(q *orm.Query) (*orm.Query, error) {
		return q.Where("parent_id IS NULL"), nil
	}
	var containers []core.Container
	err := db.Model(&containers).Apply(getRootContainerFilter).Select()
	if err == pg.ErrNoRows {
		return containers, nil
	}
	return containers, err
}

// GetContainerChildrenList gets containers from a pg database which have same parent
func (adapter *Adapter) GetContainerChildrenList(id int32) ([]core.Container, error) {
	db := pg.Connect(&adapter.ConnectionOptions)
	defer db.Close()
	getChildrenFilter := func(q *orm.Query) (*orm.Query, error) {
		return q.Where("parent_id = ?", id), nil
	}
	var containers []core.Container
	err := db.Model(&containers).Apply(getChildrenFilter).Select()
	if err == pg.ErrNoRows {
		return containers, nil
	}
	return containers, err
}

// RemoveContainer removes container from a pg database
func (adapter *Adapter) RemoveContainer(id int32) error {
	db := pg.Connect(&adapter.ConnectionOptions)
	defer db.Close()
	container := core.Container{ID: id}
	err := db.Delete(&container)
	return err
}
