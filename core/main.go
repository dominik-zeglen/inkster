package core

import (
	"github.com/globalsign/mgo/bson"
)

// Adapter interface provides abstraction over different data sources
type Adapter interface {
	AddDirectory(Directory) (Directory, error)
	GetDirectory(bson.ObjectId) (Directory, error)
	GetDirectoryList() ([]Directory, error)
	GetRootDirectoryList() ([]Directory, error)
	GetDirectoryChildrenList(bson.ObjectId) ([]Directory, error)
	UpdateDirectory(bson.ObjectId, DirectoryInput) error
	RemoveDirectory(bson.ObjectId) error

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
	GetPagesFromDirectory(bson.ObjectId) ([]Page, error)
	UpdatePage(bson.ObjectId, PageInput) error
	UpdatePageField(bson.ObjectId, string, string) error
	RemovePage(bson.ObjectId) error
	RemovePageField(bson.ObjectId, string) error

	AddUser(User) (User, error)
	AuthenticateUser(string, string) (User, error)
	GetUser(bson.ObjectId) (User, error)
	GetUserByEmail(string) (User, error)
	GetUserList() ([]User, error)
	UpdateUser(bson.ObjectId, UserInput) (User, error)
	RemoveUser(bson.ObjectId) error

	GetCurrentTime() string
	ResetMockDatabase(
		directories []Directory,
		templates []Template,
		pages []Page,
		users []User,
	)
	String() string
}

// BaseModel is an abstraction over that all models
// should be composed of, providing most basic
// fields to keep order and consistency within code
type BaseModel struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	CreatedAt string        `json:"createdAt" bson:"createdAt"`
	UpdatedAt string        `json:"updatedAt" bson:"updatedAt"`
}

// FieldTypes holds all allowed template field type names
var FieldTypes = []string{
	"directory",
	"file",
	"image",
	"longText",
	"page",
	"text",
	"unique",
}
