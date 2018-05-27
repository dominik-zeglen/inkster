package mongodb

import (
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/dominik-zeglen/ecoknow/core"
)

func init() {
	resetDatabase()
}

func TestAddContainer(t *testing.T) {
	container := core.Container{
		Name:     "New Container",
		ParentID: "000000000001",
	}
	result, err := dataSource.AddContainer(container)
	id := result.ID
	result.ID = ""
	if err != nil {
		t.Fatal(err)
	}
	data, err := ToJSON(result)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, data)
	dataSource.RemoveContainer(id)
}

func TestAddContainerWithoutName(t *testing.T) {
	container := core.Container{
		ParentID: "000000000001",
	}
	_, err := dataSource.AddContainer(container)
	if err == nil {
		t.Error(ErrNoError)
	}
}

func TestBuildTree(t *testing.T) {
	parent := core.Container{
		ID:   "200000000000",
		Name: "Parent",
	}
	result, err := dataSource.AddContainer(parent)
	parentID := result.ID
	if err != nil {
		t.Fatal(err)
	}
	data, err := ToJSON(result)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, data)

	child := core.Container{
		ID:       "200000000001",
		Name:     "Child",
		ParentID: parentID,
	}
	result, err = dataSource.AddContainer(child)
	childID := result.ID
	if err != nil {
		t.Fatal(err)
	}
	data, err = ToJSON(result)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotMulti("Child", data)

	resultList, err := dataSource.GetContainerChildrenList(parentID)
	if err != nil {
		t.Fatal(err)
	}
	data, err = ToJSON(resultList)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotMulti("Parent's children", data)

	dataSource.RemoveContainer(childID)
	dataSource.RemoveContainer(parentID)
}

func TestGetContainer(t *testing.T) {
	result, err := dataSource.GetContainer("000000000001")
	if err != nil {
		t.Error(err)
	}
	data, err := ToJSON(result)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, data)
}

func TestGetNonExistingContainer(t *testing.T) {
	_, err := dataSource.GetContainer("000000000099")
	if err == nil {
		t.Error(ErrNoError)
	}
}

func TestGetRootContainerList(t *testing.T) {
	result, _ := dataSource.GetRootContainerList()
	data, err := ToJSON(result)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, data)
}

func TestGetContainerChildren(t *testing.T) {
	result, _ := dataSource.GetContainerChildrenList("000000000001")
	data, err := ToJSON(result)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, data)
}

func TestGetContainerList(t *testing.T) {
	result, _ := dataSource.GetContainerList()
	data, err := ToJSON(result)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, data)
}

func TestUpdateContainer(t *testing.T) {
	defer resetDatabase()
	err := dataSource.UpdateContainer("000000000001", core.UpdateContainerArguments{
		Name:     "Updated container name",
		ParentID: "000000000002",
	})
	if err != nil {
		t.Error(err)
	}
	result, err := dataSource.GetContainer("000000000001")
	if err != nil {
		t.Error(err)
	}
	data, err := ToJSON(result)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, data)
}

func TestUpdateContainerName(t *testing.T) {
	defer resetDatabase()
	err := dataSource.UpdateContainer("000000000004", core.UpdateContainerArguments{
		Name: "Updated container name",
	})
	if err != nil {
		t.Error(err)
	}
	result, err := dataSource.GetContainer("000000000004")
	if err != nil {
		t.Error(err)
	}
	data, err := ToJSON(result)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, data)
}

func TestRemoveContainer(t *testing.T) {
	err := dataSource.RemoveContainer("000000000004")
	if err != nil {
		t.Error(err)
	}
}

func TestRemoveNonExistingContainer(t *testing.T) {
	err := dataSource.RemoveContainer("000000000099")
	if err == nil {
		t.Error(ErrNoError)
	}
}
