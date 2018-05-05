package mock

import (
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/dominik-zeglen/ecoknow/core"
)

var dataSource = Adapter{}

func TestAddContainer(t *testing.T) {
	defer dataSource.ResetMockDatabase()
	container := core.Container{
		Name:     "New Container",
		ParentID: 2,
	}
	_, err := dataSource.AddContainer(container)
	if err != nil {
		t.Error(err)
	}
}

func TestAddContainerWithExplicitID(t *testing.T) {
	defer dataSource.ResetMockDatabase()
	container := core.Container{
		ID:       5,
		Name:     "New Container",
		ParentID: 2,
	}
	_, err := dataSource.AddContainer(container)
	if err != nil {
		t.Error(err)
	}
}

func TestAddContainerWithTooSmallExplicitID(t *testing.T) {
	defer dataSource.ResetMockDatabase()
	container := core.Container{
		ID:       3,
		Name:     "New Container",
		ParentID: 2,
	}
	_, err := dataSource.AddContainer(container)
	if err == nil {
		t.Error("Did not return error")
	}
}

// Test if mock function is able to retrieve existing container
func TestGetContainer(t *testing.T) {
	container, err := dataSource.GetContainer(1)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, container.Json())
}

// Test if mock function is able to throw error if getting non-existing container
func TestGetNonExistingContainer(t *testing.T) {
	_, err := dataSource.GetContainer(5)
	if err == nil {
		t.Error("Did not return error")
	}
}

func TestGetRootContainerList(t *testing.T) {
	containers, _ := dataSource.GetRootContainerList()
	cupaloy.SnapshotT(t, containers)
}

func TestGetContainerChildren(t *testing.T) {
	containers, _ := dataSource.GetContainerChildrenList(1)
	cupaloy.SnapshotT(t, containers)
}

func TestGetContainerList(t *testing.T) {
	containers, _ := dataSource.GetContainerList()
	cupaloy.SnapshotT(t, containers)
}
