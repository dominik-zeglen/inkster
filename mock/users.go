package mock

import (
	"fmt"

	"github.com/dominik-zeglen/inkster/core"
	"github.com/globalsign/mgo/bson"
)

func (adapter Adapter) findUser(id *bson.ObjectId, name *string) (int, error) {
	if id != nil {
		for index := range users {
			if users[index].ID == *id {
				return index, nil
			}
		}
		return 0, fmt.Errorf("User %s does not exist", id)
	}
	if name != nil {
		for index := range users {
			if users[index].Email == *name {
				return index, nil
			}
		}
		return 0, fmt.Errorf("User %s does not exist", *name)
	}
	if id == nil && name == nil {
		return 0, fmt.Errorf("findUser() requires at least one argument")
	}

	return 0, nil
}

// AddUser puts user in the database
func (adapter Adapter) AddUser(user core.User) (core.User, error) {
	errs := user.Validate()
	if len(errs) > 0 {
		return core.User{}, core.ErrNotValidated
	}
	_, err := adapter.findUser(nil, &user.Email)
	if err == nil {
		return core.User{}, core.ErrUserExists(user.Email)
	}
	if user.ID == "" {
		user.ID = bson.NewObjectId()
	} else {
		_, err = adapter.findUser(&user.ID, nil)
		if err == nil {
			return core.User{}, core.ErrUserExists(user.ID.String())
		}
	}
	user.CreatedAt = adapter.GetCurrentTime()
	user.UpdatedAt = adapter.GetCurrentTime()

	users = append(users, user)
	return user, nil
}

// AuthenticateUser checks if given credentials are valid, then returns User object
func (adapter Adapter) AuthenticateUser(email string, password string) (core.User, error) {
	userIndex, err := adapter.findUser(nil, &email)
	if err != nil {
		return core.User{}, err
	}
	user := users[userIndex]
	if user.AuthPassword(password) {
		return user, nil
	}
	return core.User{}, core.ErrBadCredentials
}

// GetUser allows user to fetch user from database
func (adapter Adapter) GetUser(userID bson.ObjectId) (core.User, error) {
	index, err := adapter.findUser(&userID, nil)
	return users[index], err
}

// GetUserByEmail allows user to fetch user from database using his email address
func (adapter Adapter) GetUserByEmail(email string) (core.User, error) {
	index, err := adapter.findUser(nil, &email)
	return users[index], err
}

// GetUserList allows user to fetch all users from database
func (adapter Adapter) GetUserList() ([]core.User, error) {
	return users, nil
}

// UpdateUser allows user to update his properties
func (adapter Adapter) UpdateUser(userID bson.ObjectId, data core.UserInput) (core.User, error) {
	index, err := adapter.findUser(&userID, nil)
	if err != nil {
		return core.User{}, err
	}
	users[index].UpdatedAt = adapter.GetCurrentTime()

	if data.Password != nil {
		dummy := core.User{}
		dummy.CreatePassword(*data.Password)
		users[index].Password = dummy.Password
		users[index].Salt = dummy.Salt
	}
	if data.Active != nil {
		users[index].Active = *data.Active
	}
	if data.Email != nil {
		users[index].Email = *data.Email
	}
	return users[index], nil
}

// RemoveUser removes user from database
func (adapter Adapter) RemoveUser(userID bson.ObjectId) error {
	index, err := adapter.findUser(&userID, nil)
	if err != nil {
		return err
	}
	users = append(users[:index], users[index+1:]...)
	return nil
}
