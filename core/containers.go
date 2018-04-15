package core

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type Container struct {
	Id       int32
	Name     string
	ParentId int32
}

func (container Container) String() string {
	return fmt.Sprintf("Container<%d %s>", container.Id, container.Name)
}

func AddContainer(container Container) (Container, error) {
	db := pg.Connect(DbOptions)
	err := db.Insert(&container)
	defer db.Close()
	return container, err
}

func GetContainer(id int32) (Container, error) {
	db := pg.Connect(DbOptions)
	container := Container{Id: id}
	err := db.Select(&container)
	defer db.Close()
	return container, err
}

func GetContainerList() ([]Container, error) {
	db := pg.Connect(DbOptions)
	var containers []Container
	err := db.Model(&containers).Select()
	defer db.Close()
	return containers, err
}

func GetRootContainerList() ([]Container, error) {
	db := pg.Connect(DbOptions)
	defer db.Close()
	getRootContainerFilter := func(q *orm.Query) (*orm.Query, error) {
		return q.Where("parent_id IS NULL"), nil
	}
	var containers []Container
	err := db.Model(&containers).Apply(getRootContainerFilter).Select()
	return containers, err
}

func GetContainerChildrenList(id int32) ([]Container, error) {
	db := pg.Connect(DbOptions)
	getChildrenFilter := func(q *orm.Query) (*orm.Query, error) {
		return q.Where("parent_id = ?", id), nil
	}
	var containers []Container
	err := db.Model(&containers).Apply(getChildrenFilter).Select()
	defer db.Close()
	return containers, err
}

func RemoveContainer(id int32) error {
	db := pg.Connect(DbOptions)
	container := Container{Id: id}
	err := db.Delete(&container)
	defer db.Close()
	return err
}
