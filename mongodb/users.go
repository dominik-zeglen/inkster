package mongodb

import (
	"github.com/dominik-zeglen/inkster/core"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// AddUser puts user in the database
func (adapter Adapter) AddUser(user core.User) (core.User, error) {
	err := user.Validate()
	if err != nil {
		return core.User{}, err
	}
	session := adapter.Session.Copy()
	session.SetSafe(&mgo.Safe{})
	defer session.Close()

	collection := session.DB(adapter.DBName).C("users")
	found, err := collection.
		Find(bson.M{"email": user.Email}).
		Count()
	if found > 0 {
		return core.User{}, core.ErrUserExists(user.Email)
	}
	if user.ID == "" {
		user.ID = bson.NewObjectId()
	}
	user.CreatedAt = adapter.GetCurrentTime()
	user.UpdatedAt = adapter.GetCurrentTime()

	err = collection.Insert(user)
	return user, err
}

// AuthenticateUser checks if given credentials are valid, then returns User object
func (adapter Adapter) AuthenticateUser(email string, password string) (core.User, error) {
	session := adapter.Session.Copy()
	defer session.Close()

	collection := session.DB(adapter.DBName).C("users")

	var user core.User
	err := collection.
		Find(bson.M{"email": email}).
		One(&user)
	if err != nil {
		if err == mgo.ErrNotFound {
			return core.User{}, core.ErrBadCredentials
		}
		return core.User{}, err
	}

	if user.AuthPassword(password) {
		return user, nil
	}
	return core.User{}, core.ErrBadCredentials
}

// GetUser allows user to fetch user from database
func (adapter Adapter) GetUser(userID bson.ObjectId) (core.User, error) {
	session := adapter.Session.Copy()
	session.SetSafe(&mgo.Safe{})
	defer session.Close()

	collection := session.DB(adapter.DBName).C("users")
	var user core.User
	err := collection.
		FindId(userID).
		One(&user)
	return user, err
}

// GetUserByEmail allows user to fetch user from database using his email address
func (adapter Adapter) GetUserByEmail(email string) (core.User, error) {
	session := adapter.Session.Copy()
	session.SetSafe(&mgo.Safe{})
	defer session.Close()

	collection := session.DB(adapter.DBName).C("users")
	var user core.User
	err := collection.
		Find(bson.M{
			"email": email,
		}).
		One(&user)
	return user, err
}

// GetUserList allows user to fetch all users from database
func (adapter Adapter) GetUserList() ([]core.User, error) {
	session := adapter.Session.Copy()
	session.SetSafe(&mgo.Safe{})
	defer session.Close()

	collection := session.DB(adapter.DBName).C("users")
	var users []core.User
	err := collection.
		Find(bson.M{}).
		All(&users)
	return users, err
}

type UserUpdateData struct {
	Active    *bool   `bson:",omitempty"`
	Email     *string `bson:",omitempty"`
	Password  *string `bson:",omitempty"`
	Salt      *string `bson:",omitempty"`
	UpdatedAt string  `bson:"updatedAt"`
}

// UpdateUser allows user to update his properties
func (adapter Adapter) UpdateUser(userID bson.ObjectId, data core.UserInput) (core.User, error) {
	session := adapter.Session.Copy()
	session.SetSafe(&mgo.Safe{})
	defer session.Close()

	userUpdateData := UserUpdateData{
		UpdatedAt: adapter.GetCurrentTime(),
	}

	if data.Password != nil {
		dummy := core.User{}
		dummy.CreatePassword(*data.Password)
		userUpdateData.Password = &dummy.Password
		userUpdateData.Salt = &dummy.Salt
	}

	userUpdateData.Email = data.Email
	userUpdateData.Active = data.Active

	user := core.User{}
	collection := session.DB(adapter.DBName).C("users")
	_, err := collection.FindId(userID).Apply(
		mgo.Change{
			Update: bson.M{
				"$set": userUpdateData,
			},
			ReturnNew: true,
		},
		&user,
	)
	return user, err
}

// RemoveUser removes user from database
func (adapter Adapter) RemoveUser(userID bson.ObjectId) error {
	session := adapter.Session.Copy()
	session.SetSafe(&mgo.Safe{})
	defer session.Close()

	collection := session.DB(adapter.DBName).C("users")
	return collection.RemoveId(userID)
}
