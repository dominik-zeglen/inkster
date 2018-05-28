package core

import (
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

// FieldTypes holds all allowed template field type names
var FieldTypes = []string{
	"container",
	"file",
	"image",
	"longText",
	"page",
	"text",
	"unique",
}
