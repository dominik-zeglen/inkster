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
						author {
							id
							email
						}
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
						author {
							id
							email
						}
						createdAt
						updatedAt
						name
						slug
						isPublished
						fields {
							id
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
			}`, toGlobalID(gqlDirectory, 1))
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
			}`, toGlobalID(gqlDirectory, 1))
			result, err := execQuery(createPage, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Update page properties", func(t *testing.T) {
			defer resetDatabase()

			id := toGlobalID(gqlPage, 1)
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

			id := toGlobalID(gqlPage, 1)
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

			id := toGlobalID(gqlPage, 1)
			pageFieldID := toGlobalID(gqlPageField, 100)
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

			id := toGlobalID(gqlPage, 1)
			pageFieldID := toGlobalID(gqlPageField, 100)
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
			id := toGlobalID(gqlPage, 1)
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
					author {
						id
						email
					}
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
		getPages := `
			query Pages($sort: PageSort){
				pages(sort: $sort, paginate: { first: 5 }) {
					edges {
						node {
							id
							author {
								id
								email
							}
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
					}
				}
			}`
		t.Run("Get page", func(t *testing.T) {
			id := toGlobalID(gqlPage, 1)
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
			id := toGlobalID(gqlPage, 99)
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
			result, err := execQuery(getPages, "{}", nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})

		testSortingFields := []string{
			"AUTHOR",
			"CREATED_AT",
			"IS_PUBLISHED",
			"NAME",
			"SLUG",
			"UPDATED_AT",
		}
		testSortingOrders := []string{"ASC", "DESC"}

		for _, field := range testSortingFields {
			for _, order := range testSortingOrders {
				t.Run(
					"Get page list using "+field+" "+order,
					func(t *testing.T) {
						variables := fmt.Sprintf(`{
							"sort": {
								"field": "%s",
								"order": "%s"
							}
						}`, field, order)
						result, err := execQuery(getPages, variables, nil)
						if err != nil {
							t.Fatal(err)
						}
						cupaloy.SnapshotT(t, result)
					},
				)
			}
		}
	})
}
