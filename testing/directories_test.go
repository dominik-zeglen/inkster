package testing

import (
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/dominik-zeglen/inkster/core"
)

func testDirectories(t *testing.T) {
	dummyID := 0
	t.Run("Test setters", func(t *testing.T) {
		t.Run("Add directory", func(t *testing.T) {
			defer resetDatabase()
			directory := core.Directory{
				Name:     "New Directory",
				ParentID: Directories[0].ID,
			}
			result, err := dataSource.AddDirectory(directory)
			result.ID = dummyID
			if err != nil {
				t.Fatal(err)
			}
			data, err := ToJSON(result)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotT(t, data)
		})
		t.Run("Add directory without name", func(t *testing.T) {
			defer resetDatabase()
			directory := core.Directory{
				ParentID: Directories[0].ID,
			}
			_, err := dataSource.AddDirectory(directory)
			if err == nil {
				t.Error(ErrNoError)
			}
		})
		t.Run("Update directory", func(t *testing.T) {
			defer resetDatabase()
			name := "Updated directory name"
			parentID := Directories[1].ID
			isPublished := true
			err := dataSource.UpdateDirectory(Directories[0].ID, core.DirectoryInput{
				Name:        &name,
				ParentID:    &parentID,
				IsPublished: &isPublished,
			})
			if err != nil {
				t.Error(err)
			}
			result, err := dataSource.GetDirectory(Directories[0].ID)
			if err != nil {
				t.Error(err)
			}
			data, err := ToJSON(result)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotT(t, data)
		})
		t.Run("Update directory's name", func(t *testing.T) {
			defer resetDatabase()
			name := "Updated directory name"
			err := dataSource.UpdateDirectory(Directories[3].ID, core.DirectoryInput{
				Name: &name,
			})
			if err != nil {
				t.Error(err)
			}
			result, err := dataSource.GetDirectory(Directories[3].ID)
			if err != nil {
				t.Error(err)
			}
			data, err := ToJSON(result)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotT(t, data)
		})
		t.Run("Remove directory", func(t *testing.T) {
			defer resetDatabase()
			err := dataSource.RemoveDirectory(Directories[3].ID)
			if err != nil {
				t.Error(err)
			}
		})
	})
	t.Run("Test getters", func(t *testing.T) {
		t.Run("Get directory", func(t *testing.T) {
			result, err := dataSource.GetDirectory(Directories[0].ID)
			if err != nil {
				t.Error(err)
			}
			data, err := ToJSON(result)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotT(t, data)
		})
		t.Run("Get directory that does not exist", func(t *testing.T) {
			_, err := dataSource.GetDirectory(99)
			if err == nil {
				t.Error(ErrNoError)
			}
		})
		t.Run("Get root directory list", func(t *testing.T) {
			result, _ := dataSource.GetRootDirectoryList()
			data, err := ToJSON(result)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotT(t, data)
		})
		t.Run("Get directory's children", func(t *testing.T) {
			result, _ := dataSource.GetDirectoryChildrenList(Directories[0].ID)
			data, err := ToJSON(result)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotT(t, data)
		})
		t.Run("Get all directories", func(t *testing.T) {
			result, _ := dataSource.GetDirectoryList()
			data, err := ToJSON(result)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotT(t, data)
		})
	})
	t.Run("Test complex behaviour", func(t *testing.T) {
		t.Run("Build directory tree", func(t *testing.T) {
			defer resetDatabase()
			parent := core.Directory{
				Name: "Parent",
			}
			result, err := dataSource.AddDirectory(parent)
			parentID := result.ID
			if err != nil {
				t.Fatal(err)
			}
			result.ID = 0
			data, err := ToJSON(result)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotT(t, data)

			child := core.Directory{
				Name:     "Child",
				ParentID: parentID,
			}
			result, err = dataSource.AddDirectory(child)
			if err != nil {
				t.Fatal(err)
			}
			result.ID = 0
			data, err = ToJSON(result)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotMulti("Child", data)

			resultList, err := dataSource.GetDirectoryChildrenList(parentID)
			if err != nil {
				t.Fatal(err)
			}
			for idx := range resultList {
				resultList[idx].ID = 0
			}
			data, err = ToJSON(resultList)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotMulti("Parent's children", data)
		})
	})
}
