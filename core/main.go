package core

// Adapter interface provides abstraction over different data sources
type Adapter interface {
	AddDirectory(Directory) (Directory, error)
	GetDirectory(int) (Directory, error)
	GetDirectoryList() ([]Directory, error)
	GetRootDirectoryList() ([]Directory, error)
	GetDirectoryChildrenList(int) ([]Directory, error)
	UpdateDirectory(int, DirectoryInput) error
	RemoveDirectory(int) error

	AddTemplate(Template) (Template, error)
	AddTemplateField(int, TemplateField) error
	GetTemplate(int) (Template, error)
	GetTemplateList() ([]Template, error)
	UpdateTemplate(int, TemplateInput) error
	RemoveTemplate(int) error
	RemoveTemplateField(int, string) error

	AddPage(Page) (Page, error)
	AddPageFromTemplate(PageInput, int) (Page, error)
	AddPageField(int, PageField) error
	GetPage(int) (Page, error)
	GetPageBySlug(string) (Page, error)
	GetPagesFromDirectory(int) ([]Page, error)
	UpdatePage(int, PageInput) error
	RemovePage(int) error
	RemovePageField(int) error

	AddUser(User) (User, error)
	AuthenticateUser(string, string) (User, error)
	GetUser(int) (User, error)
	GetUserByEmail(string) (User, error)
	GetUserList() ([]User, error)
	UpdateUser(int, UserInput) (User, error)
	RemoveUser(int) error

	GetCurrentTime() string
	ResetMockDatabase(
		directories []Directory,
		templates []Template,
		pages []Page,
		users []User,
	) error
	String() string
}

// BaseModel is an abstraction over that all models
// should be composed of, providing most basic
// fields to keep order and consistency within code
type BaseModel struct {
	ID        int    `sql:",pk,autoincrement" json:"id"`
	CreatedAt string `sql:",notnull" json:"createdAt" bson:"createdAt"`
	UpdatedAt string `sql:",notnull" json:"updatedAt" bson:"updatedAt"`
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
