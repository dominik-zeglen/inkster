package postgres

import (
	"github.com/dominik-zeglen/inkster/core"
)

// AddUser puts user in the database
func (adapter Adapter) AddUser(user core.User) (core.User, error) {
	errs := user.Validate()
	if len(errs) > 0 {
		return core.User{}, core.ErrNotValidated
	}

	user.CreatedAt = adapter.GetCurrentTime()
	user.UpdatedAt = adapter.GetCurrentTime()
	_, err := adapter.
		Session.
		Model(&user).
		Returning("id").
		Insert()

	return user, err
}

// AuthenticateUser checks if given credentials are valid, then returns User object
func (adapter Adapter) AuthenticateUser(email string, password string) (core.User, error) {
	user, err := adapter.GetUserByEmail(email)

	if user.AuthPassword(password) {
		return user, err
	}
	return core.User{}, core.ErrBadCredentials
}

// GetUser allows user to fetch user from database
func (adapter Adapter) GetUser(id int) (core.User, error) {
	user := core.User{}
	user.ID = id
	err := adapter.
		Session.
		Select(&user)
	return user, err
}

// GetUserByEmail allows user to fetch user from database using his email address
func (adapter Adapter) GetUserByEmail(email string) (core.User, error) {
	var user core.User
	err := adapter.
		Session.
		Model(&user).
		Where("email = ?", email).
		Select()
	return user, err
}

// GetUserList allows user to fetch all users from database
func (adapter Adapter) GetUserList() ([]core.User, error) {
	var users []core.User
	err := adapter.
		Session.
		Model(&users).
		Select()
	return users, err
}

// UpdateUser allows user to update his properties
func (adapter Adapter) UpdateUser(id int, data core.UserInput) (core.User, error) {
	if len(data.Validate()) > 0 {
		return core.User{}, core.ErrNotValidated
	}

	user := core.User{}
	user.ID = id
	user.UpdatedAt = adapter.GetCurrentTime()

	query := adapter.
		Session.
		Model(&user).
		Column("updated_at")

	if data.Active != nil {
		user.Active = *data.Active
		query = query.Column("active")
	}
	if data.Email != nil {
		user.Email = *data.Email
		query = query.Column("email")
	}
	if data.Password != nil {
		user.CreatePassword(*data.Password)
		query = query.
			Column("password").
			Column("salt")
	}

	_, err := query.
		WherePK().
		Update()

	if err != nil {
		return core.User{}, err
	}

	err = adapter.
		Session.
		Model(&user).
		WherePK().
		Select()

	return user, err
}

// RemoveUser removes user from database
func (adapter Adapter) RemoveUser(id int) error {
	var user core.User
	_, err := adapter.
		Session.
		Model(&user).
		Where("id = ?", id).
		Delete()
	return err
}
