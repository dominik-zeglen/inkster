package testing

import (
	"github.com/dominik-zeglen/inkster/core"
)

func createUser(
	user core.User,
	id string,
	createdAt string,
	updatedAt string,
	password string,
) core.User {
	user.CreatePassword(password)
	user.ID = id
	user.CreatedAt = createdAt
	user.UpdatedAt = updatedAt

	return user
}

func CreateUsers() []core.User {
	users := []core.User{
		core.User{Email: "user1@xample.com", Active: true},
		core.User{Email: "user2@example.com", Active: false},
		core.User{Email: "user3@example.com", Active: true},
		core.User{Email: "user4@example.com", Active: false},
	}

	users[0] = createUser(
		users[0],
		"000000000001",
		"2007-07-07T10:00:00.000Z",
		"2007-07-07T10:00:00.000Z",
		"user1passwd",
	)
	users[1] = createUser(
		users[1],
		"000000000002",
		"2007-07-07T11:00:00.000Z",
		"2007-07-07T11:00:00.000Z",
		"user2passwd",
	)
	users[2] = createUser(
		users[2],
		"000000000003",
		"2007-07-07T12:00:00.000Z",
		"2007-07-07T12:00:00.000Z",
		"user3passwd",
	)
	users[3] = createUser(
		users[3],
		"000000000004",
		"2007-07-07T13:00:00.000Z",
		"2007-07-07T13:00:00.000Z",
		"user4passwd",
	)

	return users
}
