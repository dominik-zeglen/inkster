package api

import (
	"fmt"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	test "github.com/dominik-zeglen/ecoknow/testing"
)

func TestPageAPI(t *testing.T) {
	t.Run("Mutations", func(t *testing.T) {
		resetDatabase()
		t.Run("Create page", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation CreatePage(
				$name: String!,
				$parentId: ID!,
				$fields: [PageFieldCreateInput!],
			){
				createPage(input: {
					name: $name,
					parentId: $parentId,
					fields: $fields
				}) {
					userErrors {
						field
						message
					}
					page {
						name
						slug
						fields {
							name
							type
						}
						parent {
							id
							name
						}
					}
				}
			}`
			variables := fmt.Sprintf(`{
				"name": "New Page",
				"parentId": "%s",
				"fields": [
					{
						"name": "Field 1",
						"type": "text",
						"value": "Value 1"
					},
					{
						"name": "Field 2",
						"type": "page",
						"value": "Value 2"
					}
				]
			}`, toGlobalID("directory", test.Directories[0].ID))
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Create page without fields", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation CreatePage(
				$name: String!,
				$parentId: ID!,
				$fields: [PageFieldCreateInput!],
			){
				createPage(input: {
					name: $name,
					parentId: $parentId,
					fields: $fields
				}) {
					userErrors {
						field
						message
					}
					page {
						name
						slug
						fields {
							name
							type
						}
						parent {
							id
							name
						}
					}
				}
			}`
			variables := fmt.Sprintf(`{
				"name": "New Page",
				"parentId": "%s"
			}`, toGlobalID("directory", test.Directories[0].ID))
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Add field to page", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation CreatePageField(
				$id: ID!
				$name: String!
				$type: String!
				$value: String!
			){
				createPageField(
					id: $id, 
					input: {
						name: $name,
						type: $type,
						value: $value
					}
				) {
					userErrors {
						field
						message
					}
					page {
						id
						name
						slug
						fields {
							name
							type
						}
					}
				}
			}`
			id := toGlobalID("page", test.Pages[0].ID)
			variables := fmt.Sprintf(`{
				"id": "%s",
				"name": "New field",
				"type": "text",
				"value": "value"
			}`, id)
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Add field to page without name", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation CreatePageField(
				$id: ID!
				$name: String!
				$type: String!
				$value: String!
			){
				createPageField(
					id: $id, 
					input: {
						name: $name,
						type: $type,
						value: $value
					}
				) {
					userErrors {
						field
						message
					}
					page {
						id
						name
						slug
						fields {
							name
							type
						}
					}
				}
			}`
			id := toGlobalID("page", test.Pages[0].ID)
			variables := fmt.Sprintf(`{
				"id": "%s",
				"name": "",
				"type": "text",
				"value": "value"
			}`, id)
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Rename field", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation RenamePageField(
				$id: ID!
				$name: String!
				$renameTo: String!
			){
				renamePageField(
					id: $id, 
					input: {
						name: $name,
						renameTo: $renameTo
					}
				) {
					userErrors {
						field
						message
					}
					page {
						id
						name
						slug
						fields {
							name
							type
							value
						}
					}
				}
			}`
			id := toGlobalID("page", test.Pages[0].ID)
			variables := fmt.Sprintf(`{
				"id": "%s",
				"name": "%s",
				"renameTo": "Renamed field"
			}`, id, test.Pages[0].Fields[0].Name)
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Rename field to empty string", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation RenamePageField(
				$id: ID!
				$name: String!
				$renameTo: String!
			){
				renamePageField(
					id: $id, 
					input: {
						name: $name,
						renameTo: $renameTo
					}
				) {
					userErrors {
						field
						message
					}
					page {
						id
						name
						slug
						fields {
							name
							type
							value
						}
					}
				}
			}`
			id := toGlobalID("page", test.Pages[0].ID)
			variables := fmt.Sprintf(`{
				"id": "%s",
				"name": "%s",
				"renameTo": ""
			}`, id, test.Pages[0].Fields[0].Name)
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Update page field's value", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation UpdatePageField(
				$id: ID!
				$name: String!
				$value: String!
			){
				updatePageField(
					id: $id, 
					input: {
						name: $name,
						value: $value
					}
				) {
					userErrors {
						field
						message
					}
					page {
						id
						name
						slug
						fields {
							name
							type
							value
						}
					}
				}
			}`
			id := toGlobalID("page", test.Pages[0].ID)
			variables := fmt.Sprintf(`{
				"id": "%s",
				"name": "%s",
				"value": "Updated value"
			}`, id, test.Pages[0].Fields[0].Name)
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Update page fields", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation UpdatePageFields($id: ID!, $fields: [PageFieldUpdate2Input!]) {
				updatePage(id: $id, input: { fields: $fields }) {
					userErrors {
						field
						message
					}
					page {
						id
						name
						slug
						fields {
							name
							type
							value
						}
					}
				}
			}`
			id := toGlobalID("page", test.Pages[0].ID)
			variables := fmt.Sprintf(`{
				"id": "%s",
				"fields": [{
					"name": "%s",
					"update": {
						"name": "Updated name",
						"value": "Updated value"
					}
				}]
			}`, id, test.Pages[0].Fields[0].Name)
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Update page fields and add one", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation UpdatePageFields($id: ID!, $fields: [PageFieldUpdate2Input!], $add: [PageFieldCreateInput!]) {
				updatePage(id: $id, input: { fields: $fields }, addFields: $add) {
					userErrors {
						field
						message
					}
					page {
						id
						name
						slug
						fields {
							name
							type
							value
						}
					}
				}
			}`
			id := toGlobalID("page", test.Pages[0].ID)
			variables := fmt.Sprintf(`{
				"id": "%s",
				"add": [{
					"name": "New field",
					"type": "unique",
					"value": "New value"
				}],
				"fields": [{
					"name": "%s",
					"update": {
						"name": "Updated name",
						"value": "Updated value"
					}
				}]
			}`, id, test.Pages[0].Fields[0].Name)
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Update page fields, add one and remove another", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation UpdatePageFields($id: ID!, $fields: [PageFieldUpdate2Input!], $add: [PageFieldCreateInput!], $remove: [String!]) {
				updatePage(id: $id, input: { fields: $fields }, addFields: $add, removeFields: $remove) {
					userErrors {
						field
						message
					}
					page {
						id
						name
						slug
						fields {
							name
							type
							value
						}
					}
				}
			}`
			id := toGlobalID("page", test.Pages[0].ID)
			variables := fmt.Sprintf(`{
				"id": "%s",
				"add": [{
					"name": "New field",
					"type": "unique",
					"value": "New value"
				}],
				"fields": [{
					"name": "%s",
					"update": {
						"name": "Updated name",
						"value": "Updated value"
					}
				}],
				"remove": ["%s"]
			}`, id, test.Pages[0].Fields[0].Name, test.Pages[0].Fields[1].Name)
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Remove two fields from page", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation UpdatePageFields($id: ID!, $remove: [String!]) {
				updatePage(id: $id, removeFields: $remove) {
					userErrors {
						field
						message
					}
					page {
						id
						name
						slug
						fields {
							name
							type
							value
						}
					}
				}
			}`
			id := toGlobalID("page", test.Pages[0].ID)
			variables := fmt.Sprintf(`{
				"id": "%s",
				"remove": ["%s", "%s"]
			}`, id, test.Pages[0].Fields[0].Name, test.Pages[0].Fields[1].Name)
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Update page properties", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation UpdatePageProperties($id: ID!, $name: String) {
				updatePage(id: $id, input: { name: $name }) {
					userErrors {
						field
						message
					}
					page {
						id
						name
						slug
						fields {
							name
							type
							value
						}
					}
				}
			}`
			id := toGlobalID("page", test.Pages[0].ID)
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
		t.Run("Remove field from page", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation RemovePageField(
				$id: ID!
				$name: String!
			){
				removePageField(
					id: $id, 
					input: {
						name: $name
					}
				) {
					userErrors {
						field
						message
					}
					page {
						id
						name
						slug
						fields {
							name
							type
						}
					}
				}
			}`
			id := toGlobalID("page", test.Pages[0].ID)
			variables := fmt.Sprintf(`{
				"id": "%s",
				"name": "%s"
			}`, id, test.Pages[0].Fields[0].Name)
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Remove page", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation RemovePage(
				$id: ID!
			){
				removePage(id: $id) {
					removedObjectId
				}
			}`
			id := toGlobalID("page", test.Pages[0].ID)
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
		t.Run("Get page", func(t *testing.T) {
			query := `query getPage($id: ID!){
				page(id: $id) {
					id
					name
					slug
					fields {
						name
						type
						value
					}
					parent {
						id
						name
					}
				}
			}`
			id := toGlobalID("page", test.Pages[0].ID)
			variables := fmt.Sprintf(`{
				"id": "%s"
			}`, id)
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Get page that does not exist", func(t *testing.T) {
			query := `query getPage($id: ID!){
				page(id: $id) {
					id
					name
					slug
					fields {
						name
						type
						value
					}
					parent {
						id
						name
					}
				}
			}`
			id := toGlobalID("page", "000000000099")
			variables := fmt.Sprintf(`{
				"id": "%s"
			}`, id)
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Get page list", func(t *testing.T) {
			query := `query Pages{
				pages {
					id
					name
					slug
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
