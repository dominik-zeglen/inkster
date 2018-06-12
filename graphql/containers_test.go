package graphql

import (
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	test "github.com/dominik-zeglen/ecoknow/testing"
)

func TestContainerAPI(t *testing.T) {
	t.Run("Mutations", func(t *testing.T) {
		t.Run("Create container", func(t *testing.T) {
			query := `mutation CreateContainer(
				$name: String!,
				$parentId: String, 
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
			variables := fmt.Sprintf(`{
				"name": "New Container",
				"parentId": "%s"
			}`, test.containers[0].ID)
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Create without parent container", func(t *testing.T) {
			query := `mutation CreateContainer(
				$name: String!,
				$parentId: String, 
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
			variables := `{
				"name": "New Container"
			}`
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
	})
}
