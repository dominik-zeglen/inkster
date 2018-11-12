package api

import (
	"fmt"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	test "github.com/dominik-zeglen/inkster/testing"
)

func TestDirectoryAPI(t *testing.T) {
	createDirectory := `
		mutation CreateDirectory($input: DirectoryCreateInput!) {
			createDirectory(input: $input) {
				errors {
					code
					field
					message
				}
				directory {
					createdAt
					updatedAt
					name
					isPublished
					parent {
						id
						name
					}
				}
			}
		}`
	updateDirectory := `
		mutation UpdateDirectory(
			$id: ID!
			$input: DirectoryUpdateInput!
		) {
			updateDirectory(
				id: $id, 
				input: $input
			) {
				errors {
					code
					field
					message
				}
				directory {
					id
					createdAt
					updatedAt
					name
					isPublished
					parent {
						id
						name
					}
					children {
						id
						name
					}
			}
			}
		}`
	t.Run("Mutations", func(t *testing.T) {
		t.Run("Create directory", func(t *testing.T) {
			defer resetDatabase()
			parentID := toGlobalID("directory", test.Directories[0].ID)
			variables := fmt.Sprintf(`{
				"input": {
					"name": "New Directory",
					"parentId": "%s"
				}
			}`, parentID)
			result, err := execQuery(createDirectory, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Create without parent directory", func(t *testing.T) {
			defer resetDatabase()
			variables := `{
				"input": {
					"name": "New Directory"
				}
			}`
			result, err := execQuery(createDirectory, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Create with too short name", func(t *testing.T) {
			defer resetDatabase()
			variables := `{
				"input": {
					"name": "a"
				}
			}`
			result, err := execQuery(createDirectory, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Create directory in parent that does not exist", func(t *testing.T) {
			defer resetDatabase()
			parentID := toGlobalID("directory", 5)
			variables := fmt.Sprintf(`{
				"input": {
					"name": "New Directory",
					"parentId": "%s"
				}
			}`, parentID)
			result, err := execQuery(createDirectory, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Update directory", func(t *testing.T) {
			defer resetDatabase()
			id := toGlobalID("directory", test.Directories[3].ID)
			parentID := toGlobalID("directory", test.Directories[1].ID)
			variables := fmt.Sprintf(`{
				"id": "%s",
				"input": {
					"name": "Updated name",
					"parentId": "%s",
					"isPublished": true
				}
			}`, id, parentID)
			result, err := execQuery(updateDirectory, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Update directory partially", func(t *testing.T) {
			defer resetDatabase()
			id := toGlobalID("directory", test.Directories[3].ID)
			variables := fmt.Sprintf(`{
				"id": "%s",
				"input": {
					"name": "Updated name"
				}
			}`, id)
			result, err := execQuery(updateDirectory, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Update directory with too short name", func(t *testing.T) {
			defer resetDatabase()
			id := toGlobalID("directory", test.Directories[3].ID)
			variables := fmt.Sprintf(`{
				"id": "%s",
				"input": {
					"name": "a"
				}
			}`, id)
			result, err := execQuery(updateDirectory, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Update directory with its own ID as parentID", func(t *testing.T) {
			defer resetDatabase()
			id := toGlobalID("directory", test.Directories[3].ID)
			variables := fmt.Sprintf(`{
				"id": "%s",
				"input": {
					"parentId": "%s"
				}
			}`, id, id)
			result, err := execQuery(updateDirectory, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Remove directory", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation RemoveDirectory($id: ID!){
				removeDirectory(id: $id)
			}`
			id := toGlobalID("directory", test.Directories[0].ID)
			variables := fmt.Sprintf(`{
				"id": "%s"
			}`, id)
			result, err := execQuery(query, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)

			_, err = dataSource.GetDirectory(test.Directories[0].ID)
			if err == nil {
				t.Fatal(ErrNoError)
			}
		})
		t.Run("Remove directory that does not exist", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation RemoveDirectory($id: ID!){
				removeDirectory(id: $id)
			}`
			id := toGlobalID("directory", test.Directories[0].ID)
			variables := fmt.Sprintf(`{
				"id": "%s"
			}`, id)
			result, err := execQuery(query, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
	})

	t.Run("Queries", func(t *testing.T) {
		t.Run("Get directory", func(t *testing.T) {
			query := `query GetDirectory($id: ID!){
				getDirectory(id: $id) {
					id
					createdAt
					updatedAt
					name
					isPublished
					parent {
						id
						name
					}
					children {
						id
						name
					}
					pages {
						id
						name
						slug
					}
				}
			}`
			id := toGlobalID("directory", test.Directories[0].ID)
			variables := fmt.Sprintf(`{
				"id": "%s"
			}`, id)
			result, err := execQuery(query, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Get directory that does not exist", func(t *testing.T) {
			query := `query GetDirectory($id: ID!){
				getDirectory(id: $id) {
					id
					createdAt
					updatedAt
					name
					isPublished
					parent {
						id
						name
					}
					children {
						id
						name
					}
					pages {
						id
						name
						slug
					}
				}
			}`
			id := toGlobalID("directory", 99)
			variables := fmt.Sprintf(`{
				"id": "%s"
			}`, id)
			result, err := execQuery(query, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Get directory list", func(t *testing.T) {
			query := `query GetDirectories{
				getDirectories {
					id
					createdAt
					updatedAt
					name
					isPublished
					parent {
						id
						name
					}
					children {
						id
						name
					}
					pages {
						id
						name
						slug
					}
				}
			}`
			result, err := execQuery(query, "{}", nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Get root directory list", func(t *testing.T) {
			query := `query GetRootDirectories{
				getRootDirectories {
					id
					createdAt
					updatedAt
					name
					isPublished
					parent {
						id
						name
					}
					children {
						id
						name
					}
					pages {
						id
						name
						slug
					}
				}
			}`
			result, err := execQuery(query, "{}", nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
	})
}
