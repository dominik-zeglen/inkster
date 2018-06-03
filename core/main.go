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
	UpdateContainer(bson.ObjectId, ContainerInput) error
	RemoveContainer(bson.ObjectId) error

	AddTemplate(Template) (Template, error)
	AddTemplateField(bson.ObjectId, TemplateField) error
	GetTemplate(bson.ObjectId) (Template, error)
	GetTemplateList() ([]Template, error)
	UpdateTemplate(bson.ObjectId, TemplateInput) error
	RemoveTemplate(bson.ObjectId) error
	RemoveTemplateField(bson.ObjectId, string) error

	AddPage(Page) (Page, error)
	AddPageFromTemplate(PageInput, bson.ObjectId) (Page, error)
	AddPageField(bson.ObjectId, PageField) error
	GetPage(bson.ObjectId) (Page, error)
	GetPageBySlug(string) (Page, error)
	GetPagesFromContainer(bson.ObjectId) ([]Page, error)
	UpdatePage(bson.ObjectId, UpdatePageArguments) error
	UpdatePageField(bson.ObjectId, string, string) error
	RemovePage(bson.ObjectId) error
	RemovePageField(bson.ObjectId, string) error
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
