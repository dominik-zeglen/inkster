package api

import (
	"fmt"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	test "github.com/dominik-zeglen/inkster/testing"
)

func TestUserAPI(t *testing.T) {
	t.Run("Mutations", func(t *testing.T) {
		resetDatabase()
		t.Run("Create user", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation CreateUser($input: UserCreateInput!){
				createUser(input: $input) {
					errors {
						field
						message
					}
					user {
						createdAt
						updatedAt
						email
						isActive
					}
				}
			}`
			variables := `{
				"input": {
					"email": "new_user@example.com",
					"password": "examplepassword"
				}
			}`
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Create user without password", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation CreateUser($input: UserCreateInput!){
				createUser(input: $input) {
					errors {
						field
						message
					}
					user {
						createdAt
						updatedAt
						email
						isActive
					}
				}
			}`
			variables := `{
				"input": {
					"email": "new_user@example.com"
				}
			}`
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Remove user", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation RemoveUser($id: ID!){
				removeUser(id: $id) {
					removedObjectId
				}
			}`
			id := toGlobalID("user", test.Users[0].ID)
			variables := fmt.Sprintf(`{
				"id": "%s"
			}`, id)
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Update user", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation UpdateUser(
				$id: ID!
				$input: UserUpdateInput!
			) {
				updateUser(id: $id, input: $input) {
					errors {
						field
						message
					}
					user {
						id
						email
						isActive
					}
				}
			}`
			id := toGlobalID("user", test.Users[0].ID)
			variables := fmt.Sprintf(`{
				"id": "%s",
				"input": {
					"email": "new_email@example.com",
					"isActive": false
				}
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
		t.Run("Get user by ID", func(t *testing.T) {
			query := `query getUser($id: ID!){
				user(id: $id) {
					id
					createdAt
					updatedAt
					isActive
				}
			}`
			id := toGlobalID("user", test.Users[0].ID)
			variables := fmt.Sprintf(`{
				"id": "%s"
			}`, id)
			result, err := execQuery(query, variables)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Get user list", func(t *testing.T) {
			query := `query getUsers {
				users {
					id
					createdAt
					updatedAt
					isActive
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
