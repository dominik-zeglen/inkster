package api

import (
	"fmt"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	test "github.com/dominik-zeglen/inkster/testing"
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
					createdAt
					updatedAt
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
					createdAt
					updatedAt
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
		t.Run("Update template", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation TemplateUpdate(
				$id: ID!
				$name: String!
			){
				templateUpdate(
					id: $id, 
					input: {
						name: $name 
					}
				) {
					userErrors {
						field
						message
					}
					template {
						id
						createdAt
						updatedAt
						name
						fields {
							name
							type
						}
					}
				}
			}`
			id := toGlobalID("template", test.Templates[0].ID)
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
		t.Run("Update template without name", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation TemplateUpdate(
				$id: ID!
				$name: String!
			){
				templateUpdate(
					id: $id, 
					input: {
						name: $name 
					}
				) {
					userErrors {
						field
						message
					}
					template {
						id
						createdAt
						updatedAt
						name
						fields {
							name
							type
						}
					}
				}
			}`
			id := toGlobalID("template", test.Templates[0].ID)
			variables := fmt.Sprintf(`{
				"id": "%s",
				"name": ""
			}`, id)
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Add field to template", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation CreateTemplateField(
				$id: ID!
				$name: String!
				$type: String!
			){
				createTemplateField(
					id: $id, 
					input: {
						name: $name,
						type: $type
					}
				) {
					userErrors {
						field
						message
					}
					template {
						id
						createdAt
						updatedAt
						name
						fields {
							name
							type
						}
					}
				}
			}`
			id := toGlobalID("template", test.Templates[0].ID)
			variables := fmt.Sprintf(`{
				"id": "%s",
				"name": "New field",
				"type": "text"
			}`, id)
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Add field to template without name", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation CreateTemplateField(
				$id: ID!
				$name: String!
				$type: String!
			){
				createTemplateField(
					id: $id, 
					input: {
						name: $name,
						type: $type
					}
				) {
					userErrors {
						field
						message
					}
					template {
						id
						createdAt
						updatedAt
						name
						fields {
							name
							type
						}
					}
				}
			}`
			id := toGlobalID("template", test.Templates[0].ID)
			variables := fmt.Sprintf(`{
				"id": "%s",
				"name": "",
				"type": "text"
			}`, id)
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Remove field from template", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation RemoveTemplateField(
				$id: ID!
				$name: String!
			){
				removeTemplateField(
					id: $id, 
					input: {
						name: $name
					}
				) {
					userErrors {
						field
						message
					}
					template {
						id
						createdAt
						updatedAt
						name
						fields {
							name
							type
						}
					}
				}
			}`
			id := toGlobalID("template", test.Templates[0].ID)
			variables := fmt.Sprintf(`{
				"id": "%s",
				"name": "%s"
			}`, id, test.Templates[0].Fields[0].Name)
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Remove template", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation RemoveTemplate(
				$id: ID!
			){
				removeTemplate(id: $id) {
					userErrors {
						field
						message
					}
					removedObjectId
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
	})
	t.Run("Queries", func(t *testing.T) {
		resetDatabase()
		t.Run("Get template", func(t *testing.T) {
			query := `query getTemplate($id: ID!){
				template(id: $id) {
					id
					createdAt
					updatedAt
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
					createdAt
					updatedAt
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
		t.Run("Get template list", func(t *testing.T) {
			query := `query Templates{
				templates {
					id
					createdAt
					updatedAt
					name
					fields {
						name
						type
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
