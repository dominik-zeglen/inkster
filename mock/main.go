package mock

import (
	"github.com/dominik-zeglen/ecoknow/core"
)

var containers = make([]core.Container, 0)
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
	dataContainers []core.Container,
	dataTemplates []core.Template,
	dataPages []core.Page,
) {
	containers = make([]core.Container, len(dataContainers))
	templates = make([]core.Template, len(dataTemplates))
	pages = make([]core.Page, len(dataPages))
	copy(containers, dataContainers)
	copy(templates, dataTemplates)
	copy(pages, dataPages)
}
