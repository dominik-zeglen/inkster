package testing

import (
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/dominik-zeglen/ecoknow/core"
	"github.com/globalsign/mgo/bson"
)

func init() {
	resetDatabase()
}

func testContainers(t *testing.T) {
	t.Run("Test setters", func(t *testing.T) {
		t.Run("Add container", func(t *testing.T) {
			defer resetDatabase()
			container := core.Container{
				Name:     "New Container",
				ParentID: Containers[0].ID,
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
		})
		t.Run("Add container without name", func(t *testing.T) {
			container := core.Container{
				ParentID: Containers[0].ID,
			}
			_, err := dataSource.AddContainer(container)
			if err == nil {
				t.Error(ErrNoError)
			}
		})
		t.Run("Update container", func(t *testing.T) {
			defer resetDatabase()
			name := "Updated container name"
			parentID := bson.ObjectId(Containers[1].ID)
			err := dataSource.UpdateContainer(Containers[0].ID, core.ContainerInput{
				Name:     &name,
				ParentID: &parentID,
			})
			if err != nil {
				t.Error(err)
			}
			result, err := dataSource.GetContainer(Containers[0].ID)
			if err != nil {
				t.Error(err)
			}
			data, err := ToJSON(result)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotT(t, data)
		})
		t.Run("Update container's name", func(t *testing.T) {
			defer resetDatabase()
			name := "Updated container name"
			err := dataSource.UpdateContainer(Containers[3].ID, core.ContainerInput{
				Name: &name,
			})
			if err != nil {
				t.Error(err)
			}
			result, err := dataSource.GetContainer(Containers[3].ID)
			if err != nil {
				t.Error(err)
			}
			data, err := ToJSON(result)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotT(t, data)
		})
		t.Run("Remove container", func(t *testing.T) {
			defer resetDatabase()
			err := dataSource.RemoveContainer(Containers[3].ID)
			if err != nil {
				t.Error(err)
			}
		})
		t.Run("Remove container that does not exist", func(t *testing.T) {
			defer resetDatabase()
			err := dataSource.RemoveContainer("000000000099")
			if err == nil {
				t.Error(ErrNoError)
			}
		})
	})
	t.Run("Test getters", func(t *testing.T) {
		t.Run("Get container", func(t *testing.T) {
			result, err := dataSource.GetContainer(Containers[0].ID)
			if err != nil {
				t.Error(err)
			}
			data, err := ToJSON(result)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotT(t, data)
		})
		t.Run("Get container that does not exist", func(t *testing.T) {
			_, err := dataSource.GetContainer("000000000099")
			if err == nil {
				t.Error(ErrNoError)
			}
		})
		t.Run("Get root container list", func(t *testing.T) {
			result, _ := dataSource.GetRootContainerList()
			data, err := ToJSON(result)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotT(t, data)
		})
		t.Run("Get container's children", func(t *testing.T) {
			result, _ := dataSource.GetContainerChildrenList(Containers[0].ID)
			data, err := ToJSON(result)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotT(t, data)
		})
		t.Run("Get all containers", func(t *testing.T) {
			result, _ := dataSource.GetContainerList()
			data, err := ToJSON(result)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotT(t, data)
		})
	})
	t.Run("Test complex behaviour", func(t *testing.T) {
		t.Run("Build container tree", func(t *testing.T) {
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
		})
	})
}
