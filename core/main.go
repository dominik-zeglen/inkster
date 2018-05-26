package core

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
)

// Adapter interface provides abstraction over different data sources
type Adapter interface {
	AddContainer(Container) (Container, error)
	GetContainer(bson.ObjectId) (Container, error)
	GetContainerList() ([]Container, error)
	GetRootContainerList() ([]Container, error)
	GetContainerChildrenList(bson.ObjectId) ([]Container, error)
	RemoveContainer(bson.ObjectId) error
}

// Container is used to create tree-like structures
type Container struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name     string        `json:"name"`
	ParentID bson.ObjectId `bson:"parentId,omitempty" json:"parentId"`
}

func (container Container) String() string {
	return fmt.Sprintf("Container<%s>", container.Name)
}

type PageField struct {
	ID    int32
	Type  string
	Value string
}

func (pageField PageField) String() string {
	return fmt.Sprintf(
		"PageField<#%d %s: %s>",
		pageField.ID,
		pageField.Type,
		pageField.Value,
	)
}

func (pageField PageField) Json() string {
	return fmt.Sprintf(
		"{ ID: %d, Type: \"%s\", Value: \"%s\" }",
		pageField.ID,
		pageField.Type,
		pageField.Value,
	)
}

type PageType struct {
	ID   int32
	Name string
}

func (pageType PageType) String() string {
	return fmt.Sprintf(
		"PageType<#%d %s>",
		pageType.ID,
		pageType.Name,
	)
}

func (pageType PageType) Json() string {
	return fmt.Sprintf(
		"{ ID: %d, Name: \"%s\" }",
		pageType.ID,
		pageType.Name,
	)
}

type Page struct {
	ID       int32
	Name     string
	ParentID int32
	TypeID   int32
}

type Migration struct {
	ID   int32
	Name string
	Date int32
}
