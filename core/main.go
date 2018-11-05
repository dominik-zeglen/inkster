package core

// Adapter interface provides abstraction over different data sources
type Adapter interface {
	AddDirectory(Directory) (Directory, error)
	GetDirectory(string) (Directory, error)
	GetDirectoryList() ([]Directory, error)
	GetRootDirectoryList() ([]Directory, error)
	GetDirectoryChildrenList(string) ([]Directory, error)
	UpdateDirectory(string, DirectoryInput) error
	RemoveDirectory(string) error

	AddTemplate(Template) (Template, error)
	AddTemplateField(string, TemplateField) error
	GetTemplate(string) (Template, error)
	GetTemplateList() ([]Template, error)
	UpdateTemplate(string, TemplateInput) error
	RemoveTemplate(string) error
	RemoveTemplateField(string, string) error

	AddPage(Page) (Page, error)
	AddPageFromTemplate(PageInput, string) (Page, error)
	AddPageField(string, PageField) error
	GetPage(string) (Page, error)
	GetPageBySlug(string) (Page, error)
	GetPagesFromDirectory(string) ([]Page, error)
	UpdatePage(string, PageInput) error
	UpdatePageField(string, string, string) error
	RemovePage(string) error
	RemovePageField(string, string) error

	AddUser(User) (User, error)
	AuthenticateUser(string, string) (User, error)
	GetUser(string) (User, error)
	GetUserByEmail(string) (User, error)
	GetUserList() ([]User, error)
	UpdateUser(string, UserInput) (User, error)
	RemoveUser(string) error

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
	ID        string `bson:"_id,omitempty" json:"id"`
	CreatedAt string `json:"createdAt" bson:"createdAt"`
	UpdatedAt string `json:"updatedAt" bson:"updatedAt"`
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
