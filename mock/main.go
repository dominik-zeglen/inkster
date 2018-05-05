package mock

import "github.com/dominik-zeglen/ecoknow/core"

// Adapter is an abstraction over database connection mock
type Adapter struct {
	core.Adapter
}

// ResetMockDatabase sets in-memory array to its initial state
func (adapter *Adapter) ResetMockDatabase() {
	containers = []core.Container{
		core.Container{ID: 1, Name: "Container 1"},
		core.Container{ID: 2, Name: "Container 2"},
		core.Container{ID: 3, Name: "Container 3"},
		core.Container{ID: 4, Name: "Container 1.1", ParentID: 1},
	}
	containerIDCounter = int32(5)
}
