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
