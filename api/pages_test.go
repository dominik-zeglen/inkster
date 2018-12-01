package api

import (
	"fmt"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
)

func TestPageAPI(t *testing.T) {
	t.Run("Mutations", func(t *testing.T) {
		createPage := `
			mutation CreatePage(
				$name: String!,
				$parentId: ID!,
				$fields: [PageFieldCreateInput!],
			) {
				createPage(
					input: {
						name: $name,
						parentId: $parentId,
						fields: $fields
					}
				) {
					errors {
						code
						field
						message
					}
					page {
						createdAt
						updatedAt
						name
						slug
						isPublished
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
		updatePage := `
				mutation UpdatePage(
					$id: ID!
					$input: PageUpdateInput
					$addFields: [PageFieldCreateInput!]
					$updateFields: [PageFieldUpdate!]
					$removeFields: [String!]
				) {
					updatePage(
					id: $id 
					input: $input
					addFields: $addFields
					updateFields: $updateFields
					removeFields: $removeFields
				) {
					errors {
						code
						field
						message
					}
					page {
						id
						createdAt
						updatedAt
						name
						slug
						isPublished
						fields {
							name
							type
							value
						}
					}
				}
			}`
		t.Run("Create page", func(t *testing.T) {
			defer resetDatabase()

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
						"type": "text",
						"value": "Value 2"
					}
				]
			}`, toGlobalID("directory", 1))
			result, err := execQuery(createPage, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Create page without fields", func(t *testing.T) {
			defer resetDatabase()

			variables := fmt.Sprintf(`{
				"name": "New Page",
				"parentId": "%s"
			}`, toGlobalID("directory", 1))
			result, err := execQuery(createPage, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Update page properties", func(t *testing.T) {
			defer resetDatabase()

			id := toGlobalID("page", 1)
			variables := fmt.Sprintf(`{
				"id": "%s",
				"input": {	
					"name": "Updated name",
					"isPublished": true
				}
			}`, id)
			result, err := execQuery(updatePage, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Add page fields", func(t *testing.T) {
			defer resetDatabase()

			id := toGlobalID("page", 1)
			variables := fmt.Sprintf(`{
				"id": "%s",
				"addFields": [
					{
						"name": "Field 3",
						"type": "text",
						"value": "some value"
					}
				]
			}`, id)
			result, err := execQuery(updatePage, variables, nil)
			if err != nil {
				t.Fatal(err)
			}

			cupaloy.SnapshotT(t, result)
		})
		t.Run("Update page fields", func(t *testing.T) {
			defer resetDatabase()

			id := toGlobalID("page", 1)
			pageFieldID := toGlobalID("pageField", 100)
			variables := fmt.Sprintf(`{
				"id": "%s",
				"updateFields": [
					{
						"id": "%s",
						"input": {
							"name": "Updated name",
							"value": "Updated value"
						}
					}
				]
			}`, id, pageFieldID)
			result, err := execQuery(updatePage, variables, nil)
			if err != nil {
				t.Fatal(err)
			}

			cupaloy.SnapshotT(t, result)
		})
		t.Run("Remove page fields", func(t *testing.T) {
			defer resetDatabase()

			id := toGlobalID("page", 1)
			pageFieldID := toGlobalID("pageField", 100)
			variables := fmt.Sprintf(`{
				"id": "%s",
				"removeFields": ["%s"]
			}`, id, pageFieldID)
			result, err := execQuery(updatePage, variables, nil)
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
			id := toGlobalID("page", 1)
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
		getPage := `
			query getPage($id: ID!){
				page(id: $id) {
					id
					createdAt
					updatedAt
					name
					slug
					isPublished
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
		t.Run("Get page", func(t *testing.T) {
			id := toGlobalID("page", 1)
			variables := fmt.Sprintf(`{
				"id": "%s"
			}`, id)
			result, err := execQuery(getPage, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Get page that does not exist", func(t *testing.T) {
			id := toGlobalID("page", 99)
			variables := fmt.Sprintf(`{
				"id": "%s"
			}`, id)
			result, err := execQuery(getPage, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Get page with bad ID", func(t *testing.T) {
			variables := fmt.Sprintf(`{
				"id": "%s"
			}`, "lorem")
			result, err := execQuery(getPage, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Get page list", func(t *testing.T) {
			query := `query Pages{
				pages {
					id
					createdAt
					updatedAt
					name
					slug
					isPublished
					fields {
						name
						type
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
