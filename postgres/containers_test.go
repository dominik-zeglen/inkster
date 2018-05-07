package postgres

import (
	"fmt"
	"os"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/dominik-zeglen/ecoknow/core"
	"github.com/go-pg/pg"
)

var connectionOptions = pg.Options{
	User:     os.Getenv("FOXXY_DB_USER"),
	Password: os.Getenv("FOXXY_DB_USER_PASSWORD"),
	Database: os.Getenv("FOXXY_DB_NAME"),
	Addr: fmt.Sprintf("%s:%s",
		os.Getenv("FOXXY_DB_ADDR"),
		os.Getenv("FOXXY_DB_PORT"),
	),
}
var dataSource = Adapter{ConnectionOptions: connectionOptions}

func init() {
	err := ApplyMigrations(dataSource)
	if err != nil {
		panic(err)
	}
	containers := []core.Container{
		core.Container{Name: "Container 1"},
		core.Container{Name: "Container 2"},
		core.Container{Name: "Container 3"},
		core.Container{Name: "Container 1.1", ParentID: 1},
	}
	for id := range containers {
		db := pg.Connect(&connectionOptions)
		err = db.Insert(&containers[id])
		if err != nil {
			panic(err)
		}
	}
}

func TestAddContainer(t *testing.T) {
	container := core.Container{
		Name:     "New Container",
		ParentID: 2,
	}
	result, err := dataSource.AddContainer(container)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, result.Json())
}

func TestAddContainerWithExplicitID(t *testing.T) {
	container := core.Container{
		ID:       6,
		Name:     "New Container",
		ParentID: 2,
	}
	result, err := dataSource.AddContainer(container)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, result.Json())
}

func TestAddContainerWithTooSmallExplicitID(t *testing.T) {
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

// Test if pg function is able to retrieve existing container
func TestGetContainer(t *testing.T) {
	container, err := dataSource.GetContainer(1)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, container.Json())
}

// Test if pg function is able to throw error if getting non-existing container
func TestGetNonExistingContainer(t *testing.T) {
	_, err := dataSource.GetContainer(99)
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

func TestRemoveContainer(t *testing.T) {
	err := dataSource.RemoveContainer(5)
	if err != nil {
		t.Error(err)
	}
}

func TestRemoveNonExistingContainer(t *testing.T) {
	err := dataSource.RemoveContainer(5)
	if err == nil {
		t.Error("Did not return error")
	}
}
