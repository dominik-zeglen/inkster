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
	DBName:        os.Getenv("FOXXY_DB_NAME"),
}

func init() {
	containers := []core.Container{
		core.Container{ID: "000000000001", Name: "Container 1"},
		core.Container{ID: "000000000002", Name: "Container 2"},
		core.Container{ID: "000000000003", Name: "Container 3"},
		core.Container{ID: "000000000004", Name: "Container 1.1", ParentID: "000000000001"},
	}
	for id := range containers {
		db, err := mgo.Dial(dataSource.ConnectionURI)
		db.SetSafe(&mgo.Safe{})
		collection := db.DB(dataSource.DBName).C("containers")
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

func TestAddContainerWithExplicitID(t *testing.T) {
	container := core.Container{
		ID:       "000000000005",
		Name:     "New Container",
		ParentID: "000000000001",
	}
	result, err := dataSource.AddContainer(container)
	if err != nil {
		t.Error(err)
	}
	data, err := json.Marshal(&result)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, data)
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
	err := dataSource.RemoveContainer("000000000005")
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
