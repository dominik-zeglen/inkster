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

// ResetMockDatabase sets in-memory array to its initial state
func (adapter Adapter) ResetMockDatabase(
	containers []core.Container,
	templates []core.Template,
	pages []core.Page,
) {
	adapter.containers = containers
	adapter.templates = templates
	adapter.pages = pages
}
