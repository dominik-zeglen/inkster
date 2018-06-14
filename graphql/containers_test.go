package graphql

import (
	"fmt"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	test "github.com/dominik-zeglen/ecoknow/testing"
)

func TestContainerAPI(t *testing.T) {
	t.Run("Mutations", func(t *testing.T) {
		resetDatabase()
		t.Run("Create container", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation CreateContainer(
				$name: String!,
				$parentId: ID, 
			){
				createContainer(input: {
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
			parentID := toGlobalID("container", test.Containers[0].ID)
			variables := fmt.Sprintf(`{
				"name": "New Container",
				"parentId": "%s"
			}`, parentID)
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Create without parent container", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation CreateContainer(
				$name: String!
			){
				createContainer(input: {
					name: $name 
				}) {
					name
				}
			}`
			variables := `{
				"name": "New Container"
			}`
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Update container", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation UpdateContainer(
				$id: ID!
				$name: String
				$parentId: ID
			){
				updateContainer(
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
			id := toGlobalID("container", test.Containers[3].ID)
			parentID := toGlobalID("container", test.Containers[1].ID)
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
		t.Run("Update container partially", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation UpdateContainer(
				$id: ID!
				$name: String
				$parentId: ID
			){
				updateContainer(
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
			id := toGlobalID("container", test.Containers[3].ID)
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
		t.Run("Remove container", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation RemoveContainer($id: ID!){
				removeContainer(id: $id)
			}`
			id := toGlobalID("container", test.Containers[0].ID)
			variables := fmt.Sprintf(`{
				"id": "%s"
			}`, id)
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)

			_, err = dataSource.GetContainer(test.Containers[0].ID)
			if err == nil {
				t.Fatal(ErrNoError)
			}
		})
		t.Run("Remove container that does not exist", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation RemoveContainer($id: ID!){
				removeContainer(id: $id)
			}`
			id := toGlobalID("container", test.Containers[0].ID)
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
		t.Run("Get container", func(t *testing.T) {
			query := `query GetContainer($id: ID!){
				getContainer(id: $id) {
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
			id := toGlobalID("container", test.Containers[0].ID)
			variables := fmt.Sprintf(`{
				"id": "%s"
			}`, id)
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Get container that does not exist", func(t *testing.T) {
			query := `query GetContainer($id: ID!){
				getContainer(id: $id) {
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
			id := toGlobalID("container", "000000000099")
			variables := fmt.Sprintf(`{
				"id": "%s"
			}`, id)
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Get container list", func(t *testing.T) {
			query := `query GetContainers{
				getContainers {
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
			result, err := execQuery(query, "{}")
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Get root container list", func(t *testing.T) {
			query := `query GetRootContainers{
				getRootContainers {
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
			result, err := execQuery(query, "{}")
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
	})
}
