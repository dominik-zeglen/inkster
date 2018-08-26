package mock

import (
	"time"

	"github.com/dominik-zeglen/inkster/core"
)

var directories = make([]core.Directory, 0)
var templates = make([]core.Template, 0)
var pages = make([]core.Page, 0)
var users = make([]core.User, 0)

// Adapter is an abstraction over database connection mock
type Adapter struct {
	core.Adapter

	GetTime func() string
}

func (adapter Adapter) GetCurrentTime() string {
	if adapter.GetTime == nil {
		return time.Now().UTC().Format(time.RFC3339)
	}
	return adapter.GetTime()
}

func (adapter Adapter) String() string {
	return "In-memory mock"
}

// ResetMockDatabase sets in-memory array to its initial state
func (adapter Adapter) ResetMockDatabase(
	dataDirectories []core.Directory,
	dataTemplates []core.Template,
	dataPages []core.Page,
	dataUsers []core.User,
) {
	directories = make([]core.Directory, len(dataDirectories))
	templates = make([]core.Template, len(dataTemplates))
	pages = make([]core.Page, len(dataPages))
	users = make([]core.User, len(dataUsers))
	copy(directories, dataDirectories)
	copy(templates, dataTemplates)
	copy(pages, dataPages)
	copy(users, dataUsers)

	for templateIndex := range templates {
		fields := make([]core.TemplateField, len(dataTemplates[templateIndex].Fields))
		copy(fields, dataTemplates[templateIndex].Fields)
		templates[templateIndex].Fields = fields
	}
	for pageIndex := range pages {
		fields := make([]core.PageField, len(dataPages[pageIndex].Fields))
		copy(fields, dataPages[pageIndex].Fields)
		pages[pageIndex].Fields = fields
	}
}
