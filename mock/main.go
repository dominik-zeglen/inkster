package mock

import (
	"github.com/dominik-zeglen/ecoknow/core"
)

// Adapter is an abstraction over database connection mock
type Adapter struct {
	core.Adapter
	containers []core.Container
	templates  []core.Template
	pages      []core.Page
}

// NewAdapter creates new Adapter
func NewAdapter(
	containers []core.Container,
	templates []core.Template,
	pages []core.Page,
) Adapter {
	return Adapter{
		containers: containers,
		templates:  templates,
		pages:      pages,
	}
}

func (adapter Adapter) String() string {
	return "In-memory mock"
}

// ResetMockDatabase sets in-memory array to its initial state
func (adapter Adapter) ResetMockDatabase(
	containers []core.Container,
	templates []core.Template,
	pages []core.Page,
) {
	copy(adapter.containers, containers)
	copy(adapter.templates, templates)
	copy(adapter.pages, pages)
}
