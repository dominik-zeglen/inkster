package mongodb

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/dominik-zeglen/ecoknow/core"
	"github.com/globalsign/mgo"
)

var dataSource = Adapter{
	ConnectionURI: os.Getenv("FOXXY_DB_URI"),
	DBName:        os.Getenv("FOXXY_DB_NAME") + "_test",
}

func init() {
	containers := []core.Container{
		core.Container{ID: "000000000001", Name: "Container 1"},
		core.Container{ID: "000000000002", Name: "Container 2"},
		core.Container{ID: "000000000003", Name: "Container 3"},
		core.Container{ID: "000000000004", Name: "Container 1.1", ParentID: "000000000001"},
	}
	session, err := mgo.Dial(dataSource.ConnectionURI)
	session.SetSafe(&mgo.Safe{})
	db := session.DB(dataSource.DBName)
	collection := db.C("containers")
	collection.DropCollection()
	for id := range containers {
		err = collection.Insert(containers[id])
		if err != nil {
			panic(err)
		}
	}
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
	data, err := json.Marshal(&result)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, data)
	dataSource.RemoveContainer(id)
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
	data, err := json.Marshal(&result)
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
	data, err = json.Marshal(&result)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotMulti("Child", data)

	resultList, err := dataSource.GetContainerChildrenList(parentID)
	if err != nil {
		t.Fatal(err)
	}
	data, err = json.Marshal(&resultList)
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
	data, err := json.Marshal(&result)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, data)
}

func TestGetNonExistingContainer(t *testing.T) {
	_, err := dataSource.GetContainer("000000000099")
	if err == nil {
		t.Error("Did not return error")
	}
}

func TestGetRootContainerList(t *testing.T) {
	result, _ := dataSource.GetRootContainerList()
	data, err := json.Marshal(&result)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, data)
}

func TestGetContainerChildren(t *testing.T) {
	result, _ := dataSource.GetContainerChildrenList("000000000001")
	data, err := json.Marshal(&result)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, data)
}

func TestGetContainerList(t *testing.T) {
	result, _ := dataSource.GetContainerList()
	data, err := json.Marshal(&result)
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
		t.Error("Did not return error")
	}
}
