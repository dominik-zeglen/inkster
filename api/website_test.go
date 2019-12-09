package api

import (
	"testing"

	"github.com/bradleyjkemp/cupaloy"
)

func TestWebsiteAPI(t *testing.T) {
	getWebsite := `
		query GetWebsite {
			website {
				name
				description
				domain
			}
		}
	`
	updateWebsite := `
		mutation UpdateWebsite($input: WebsiteUpdateInput!) {
			updateWebsite(input: $input) {
				errors {
					code
					field
					message
				}
				website {
					name
					description
					domain
				}
			}
		}
	`
	t.Run("Mutations", func(t *testing.T) {
		t.Run("Update website", func(t *testing.T) {
			defer resetDatabase()
			variables := `{
				"input": {
					"name": "Updated name",
					"description": "Updated description",
					"domain": "http://updateddomain.example.com"
				}
			}`
			result, err := execQuery(updateWebsite, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Update website with incorrect name", func(t *testing.T) {
			defer resetDatabase()
			variables := `{
				"input": {
					"name": "A"
				}
			}`
			result, err := execQuery(updateWebsite, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Update website with incorrect domain", func(t *testing.T) {
			defer resetDatabase()
			variables := `{
				"input": {
					"domain": "updateddomain"
				}
			}`
			result, err := execQuery(updateWebsite, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
	})
	t.Run("Queries", func(t *testing.T) {
		t.Run("Get website data", func(t *testing.T) {
			result, err := execQuery(getWebsite, "{}", nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
	})
}
