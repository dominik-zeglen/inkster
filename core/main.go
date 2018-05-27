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

// UpdateContainerArguments is transactional model of an update properties
type UpdateContainerArguments struct {
	Name     string        `bson:",omitempty"`
	ParentID bson.ObjectId `bson:"parentId,omitempty"`
}

func (container Container) String() string {
	return fmt.Sprintf("Container<%s>", container.Name)
}

// Template allows user to quickly create new pages without repeatedly
// adding the same fields each time
type Template struct {
	ID     bson.ObjectId   `bson:"_id,omitempty" json:"id"`
	Name   string          `json:"name"`
	Fields []TemplateField `json:"fields"`
}

// UpdateTemplateArguments is transactional model of an update properties
type UpdateTemplateArguments struct {
	Name string `bson:",omitempty"`
}

func (template Template) String() string {
	return fmt.Sprintf("Template<%s>", template.Name)
}

// TemplateFieldTypes holds all allowed template field type names
var TemplateFieldTypes = []string{
	"container",
	"file",
	"image",
	"longText",
	"page",
	"text",
	"unique",
}

// TemplateField represents a single field in template
type TemplateField struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

// ErrNoEmpty returns error informing about missing field
func ErrNoEmpty(key string) error {
	return fmt.Errorf("Field cannot be omitted: %s", key)
}
