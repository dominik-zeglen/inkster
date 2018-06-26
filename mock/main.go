package mock

import (
	"github.com/dominik-zeglen/ecoknow/core"
)

var directories = make([]core.Directory, 0)
var templates = make([]core.Template, 0)
var pages = make([]core.Page, 0)

// Adapter is an abstraction over database connection mock
type Adapter struct {
	core.Adapter
}

func (adapter Adapter) String() string {
	return "In-memory mock"
}

// ResetMockDatabase sets in-memory array to its initial state
func (adapter Adapter) ResetMockDatabase(
	dataDirectories []core.Directory,
	dataTemplates []core.Template,
	dataPages []core.Page,
) {
	directories = make([]core.Directory, len(dataDirectories))
	templates = make([]core.Template, len(dataTemplates))
	pages = make([]core.Page, len(dataPages))
	copy(directories, dataDirectories)
	copy(templates, dataTemplates)
	copy(pages, dataPages)
}
