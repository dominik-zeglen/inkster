package mock

import (
	"encoding/json"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/dominik-zeglen/ecoknow/core"
	"github.com/globalsign/mgo/bson"
)

var dataSource = Adapter{}

func TestAddContainer(t *testing.T) {
	defer dataSource.ResetMockDatabase()
	container := core.Container{
		Name:     "New Container",
		ParentID: bson.ObjectId("000000000001"),
	}
	_, err := dataSource.AddContainer(container)
	if err != nil {
		t.Error(err)
	}
}

func TestAddContainerWithExplicitID(t *testing.T) {
	defer dataSource.ResetMockDatabase()
	container := core.Container{
		ID:       bson.ObjectId("000000000005"),
		Name:     "New Container",
		ParentID: bson.ObjectId("000000000001"),
	}
	_, err := dataSource.AddContainer(container)
	if err != nil {
		t.Error(err)
	}
}

func TestAddContainerWithTooSmallExplicitID(t *testing.T) {
	defer dataSource.ResetMockDatabase()
	container := core.Container{
		ID:       bson.ObjectId("000000000002"),
		Name:     "New Container",
		ParentID: bson.ObjectId("000000000001"),
	}
	_, err := dataSource.AddContainer(container)
	if err == nil {
		t.Error("Did not return error")
	}
}

// Test if mock function is able to retrieve existing container
func TestGetContainer(t *testing.T) {
	container, err := dataSource.GetContainer(bson.ObjectId("000000000001"))
	if err != nil {
		t.Error(err)
	}
	data, err := json.Marshal(container)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, data)
}

// Test if mock function is able to throw error if getting non-existing container
func TestGetNonExistingContainer(t *testing.T) {
	_, err := dataSource.GetContainer(bson.ObjectId("000000000099"))
	if err == nil {
		t.Error("Did not return error")
	}
}

func TestGetRootContainerList(t *testing.T) {
	containers, _ := dataSource.GetRootContainerList()
	cupaloy.SnapshotT(t, containers)
}

func TestGetContainerChildren(t *testing.T) {
	containers, _ := dataSource.GetContainerChildrenList(bson.ObjectId("000000000001"))
	cupaloy.SnapshotT(t, containers)
}

func TestGetContainerList(t *testing.T) {
	containers, _ := dataSource.GetContainerList()
	cupaloy.SnapshotT(t, containers)
}

func TestRemoveContainer(t *testing.T) {
	defer dataSource.ResetMockDatabase()
	err := dataSource.RemoveContainer(bson.ObjectId("000000000002"))
	if err != nil {
		t.Error(err)
	}
}

func TestRemoveNonExistingContainer(t *testing.T) {
	err := dataSource.RemoveContainer(bson.ObjectId("000000000099"))
	if err == nil {
		t.Error("Did not return error")
	}
}
