package api

import (
	"fmt"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	test "github.com/dominik-zeglen/inkster/testing"
)

func TestDirectoryAPI(t *testing.T) {
	t.Run("Mutations", func(t *testing.T) {
		resetDatabase()
		t.Run("Create directory", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation CreateDirectory(
				$name: String!,
				$parentId: ID, 
			){
				createDirectory(input: {
					name: $name, 
					parentId: $parentId
				}) {
					name
					parent {
						id
						name
					}
				}
			}`
			parentID := toGlobalID("directory", test.Directories[0].ID)
			variables := fmt.Sprintf(`{
				"name": "New Directory",
				"parentId": "%s"
			}`, parentID)
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Create without parent directory", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation CreateDirectory(
				$name: String!
			){
				createDirectory(input: {
					name: $name 
				}) {
					name
				}
			}`
			variables := `{
				"name": "New Directory"
			}`
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Update directory", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation UpdateDirectory(
				$id: ID!
				$name: String
				$parentId: ID
			){
				updateDirectory(
					id: $id, 
					input: {
						name: $name 
						parentId: $parentId
					}
				) {
					id
					name
					parent {
						id
						name
					}
					children {
						id
						name
					}
				}
			}`
			id := toGlobalID("directory", test.Directories[3].ID)
			parentID := toGlobalID("directory", test.Directories[1].ID)
			variables := fmt.Sprintf(`{
				"id": "%s",
				"name": "Updated name",
				"parentId": "%s"
			}`, id, parentID)
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Update directory partially", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation UpdateDirectory(
				$id: ID!
				$name: String
				$parentId: ID
			){
				updateDirectory(
					id: $id, 
					input: {
						name: $name 
						parentId: $parentId
					}
				) {
					id
					name
					parent {
						id
						name
					}
					children {
						id
						name
					}
				}
			}`
			id := toGlobalID("directory", test.Directories[3].ID)
			variables := fmt.Sprintf(`{
				"id": "%s",
				"name": "Updated name"
			}`, id)
			result, err := execQuery(query, variables)
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
			result, err := execQuery(query, variables)
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
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
	})

	t.Run("Queries", func(t *testing.T) {
		resetDatabase()
		t.Run("Get directory", func(t *testing.T) {
			query := `query GetDirectory($id: ID!){
				getDirectory(id: $id) {
					id
					name
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
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Get directory that does not exist", func(t *testing.T) {
			query := `query GetDirectory($id: ID!){
				getDirectory(id: $id) {
					id
					name
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
			id := toGlobalID("directory", "000000000099")
			variables := fmt.Sprintf(`{
				"id": "%s"
			}`, id)
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Get directory list", func(t *testing.T) {
			query := `query GetDirectories{
				getDirectories {
					id
					name
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
			result, err := execQuery(query, "{}")
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Get root directory list", func(t *testing.T) {
			query := `query GetRootDirectories{
				getRootDirectories {
					id
					name
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
			result, err := execQuery(query, "{}")
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
	})
}
