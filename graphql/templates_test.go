package graphql

import (
	"fmt"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	test "github.com/dominik-zeglen/ecoknow/testing"
)

func TestTemplateAPI(t *testing.T) {
	t.Run("Mutations", func(t *testing.T) {
		resetDatabase()
		t.Run("Create template", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation CreateTemplate(
				$name: String!,
				$fields: [TemplateFieldCreateInput!],
			){
				createTemplate(input: {
					name: $name,
					fields: $fields
				}) {
					name
					fields {
						name
						type
					}
				}
			}`
			variables := `{
				"name": "New Template",
				"fields": [
					{
						"name": "Field 1",
						"type": "text"
					},
					{
						"name": "Field 2",
						"type": "page"
					}
				]
			}`
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Create template without fields", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation CreateTemplate(
				$name: String!,
				$fields: [TemplateFieldCreateInput!],
			){
				createTemplate(input: {
					name: $name,
					fields: $fields
				}) {
					name
					fields {
						name
						type
					}
				}
			}`
			variables := `{
				"name": "New Template"
			}`
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
	})
	t.Run("Queries", func(t *testing.T) {
		resetDatabase()
		t.Run("Get template", func(t *testing.T) {
			query := `query getTemplate($id: ID!){
				template(id: $id) {
					id
					name
					fields {
						name
						type
					}
				}
			}`
			id := toGlobalID("template", test.Templates[0].ID)
			variables := fmt.Sprintf(`{
				"id": "%s"
			}`, id)
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Get template that does not exist", func(t *testing.T) {
			query := `query getTemplate($id: ID!){
				template(id: $id) {
					id
					name
					fields {
						name
						type
					}
				}
			}`
			id := toGlobalID("template", "000000000099")
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
}
