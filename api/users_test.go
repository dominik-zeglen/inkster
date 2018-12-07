package api

import (
	"fmt"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
)

func TestUserAPI(t *testing.T) {
	createUser := `
		mutation CreateUser($input: UserCreateInput!) {
			createUser(input: $input) {
				errors {
					code
					field
					message
				}
				user {
					createdAt
					updatedAt
					email
					isActive
					pages {
						id
						name
					}
				}
			}
		}`

	updateUser := `
		mutation UpdateUser(
			$id: ID!
			$input: UserUpdateInput!
		) {
			updateUser(id: $id, input: $input) {
				errors {
					code
					field
					message
				}
				user {
					id
					email
					isActive
					pages {
						id
						name
					}
				}
			}
		}`
	removeUser := `
		mutation RemoveUser($id: ID!){
			removeUser(id: $id) {
				removedObjectId
			}
		}`
	t.Run("Mutations", func(t *testing.T) {
		t.Run("Create user", func(t *testing.T) {
			defer resetDatabase()
			variables := `{
				"input": {
					"email": "new_user@example.com",
					"password": "examplepassword"
				}
			}`
			result, err := execQuery(createUser, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Create user without password", func(t *testing.T) {
			defer resetDatabase()
			variables := `{
				"input": {
					"email": "new_user@example.com"
				}
			}`
			result, err := execQuery(createUser, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Create user with invalid e-mail", func(t *testing.T) {
			defer resetDatabase()
			variables := `{
				"input": {
					"email": "invalidemail"
				}
			}`
			result, err := execQuery(createUser, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Remove user", func(t *testing.T) {
			defer resetDatabase()
			id := toGlobalID("user", 2)
			variables := fmt.Sprintf(`{
				"id": "%s"
			}`, id)
			result, err := execQuery(removeUser, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Remove user using his own token", func(t *testing.T) {
			defer resetDatabase()
			id := toGlobalID("user", 1)
			variables := fmt.Sprintf(`{
				"id": "%s"
			}`, id)
			result, err := execQuery(removeUser, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Update user", func(t *testing.T) {
			defer resetDatabase()
			id := toGlobalID("user", 1)
			variables := fmt.Sprintf(`{
				"id": "%s",
				"input": {
					"email": "new_email@example.com",
					"isActive": false
				}
			}`, id)
			result, err := execQuery(updateUser, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Update user with invalid e-mail", func(t *testing.T) {
			defer resetDatabase()
			id := toGlobalID("user", 1)
			variables := fmt.Sprintf(`{
				"id": "%s",
				"input": {
					"email": "invalidemail"
				}
			}`, id)
			result, err := execQuery(updateUser, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Update user with existing e-mail", func(t *testing.T) {
			defer resetDatabase()
			id := toGlobalID("user", 1)
			variables := fmt.Sprintf(`{
				"id": "%s",
				"input": {
					"email": "%s"
				}
			}`, id, "user2@example.com")
			result, err := execQuery(updateUser, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Update user with the same e-mail", func(t *testing.T) {
			defer resetDatabase()
			id := toGlobalID("user", 1)
			variables := fmt.Sprintf(`{
				"id": "%s",
				"input": {
					"email": "%s"
				}
			}`, id, "user1@example.com")
			result, err := execQuery(updateUser, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Reset password", func(t *testing.T) {
			defer resetDatabase()
			query := `mutation SendResetPasswordEmail(
				$email: String!
			) {
				sendUserPasswordResetToken(email: $email) 
			}`
			variables := fmt.Sprintf(`{
				"email": "%s"
			}`, "user1@example.com")
			r, err := execQuery(query, variables, nil)
			if err != nil || r == "" {
				t.Fatal(err)
			}

			query = `mutation ResetPassword(
				$token: String!,
				$password: String!
			) {
				resetUserPassword(token: $token, password: $password)
			}`
			variables = fmt.Sprintf(`{
				"password": "examplePassword",
				"token": "%s"
			}`, mailClient.Last())
			result, err := execQuery(query, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
	})
	t.Run("Queries", func(t *testing.T) {
		getUser := `
			query getUser($id: ID!){
				user(id: $id) {
					id
					createdAt
					updatedAt
					isActive
					pages {
						id
						name
					}
				}
			}`
		t.Run("Get user by ID", func(t *testing.T) {
			id := toGlobalID("user", 1)
			variables := fmt.Sprintf(`{
				"id": "%s"
			}`, id)
			result, err := execQuery(getUser, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Get user that does not exists", func(t *testing.T) {
			id := toGlobalID("user", 99)
			variables := fmt.Sprintf(`{
				"id": "%s"
			}`, id)
			result, err := execQuery(getUser, variables, nil)
			if err != nil {
				t.Fatal(err)
			}
			cupaloy.SnapshotT(t, result)
		})
		t.Run("Get user with bad ID", func(t *testing.T) {
			variables := fmt.Sprintf(`{
				"id": "%s"
			}`, "lorem")
			result, err := execQuery(getUser, variables, nil)
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
					pages {
						id
						name
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
