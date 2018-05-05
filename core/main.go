package core

import (
	"fmt"
)

// Adapter interface provides abstraction over different data sources
type Adapter interface {
	AddContainer(Container) (Container, error)
	GetContainer(int32) (Container, error)
	GetContainerList() ([]Container, error)
	GetRootContainerList() ([]Container, error)
	GetContainerChildrenList(int32) ([]Container, error)
	RemoveContainer(int32) error
}

// Container is used to create tree-like structures
type Container struct {
	ID       int32
	Name     string
	ParentID int32
}

func (container Container) String() string {
	return fmt.Sprintf("Container<%d %s>", container.ID, container.Name)
}

func (container Container) Json() string {
	return fmt.Sprintf(
		"{ ID: %d, Name: \"%s\", ParentID: %d }",
		container.ID,
		container.Name,
		container.ParentID,
	)
}

type Migration struct {
	ID   int32
	Name string
	Date int32
}
