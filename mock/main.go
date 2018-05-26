package mock

import (
	"github.com/dominik-zeglen/ecoknow/core"
	"github.com/globalsign/mgo/bson"
)

// Adapter is an abstraction over database connection mock
type Adapter struct {
	core.Adapter
}

// ResetMockDatabase sets in-memory array to its initial state
func (adapter *Adapter) ResetMockDatabase() {
	containers = []core.Container{
		core.Container{ID: bson.ObjectId("000000000001"), Name: "Container 1"},
		core.Container{ID: bson.ObjectId("000000000002"), Name: "Container 2"},
		core.Container{ID: bson.ObjectId("000000000003"), Name: "Container 3"},
		core.Container{ID: bson.ObjectId("000000000004"), Name: "Container 1.1", ParentID: bson.ObjectId("000000000001")},
	}
}
