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

// RemoveUser removes user from database
func (adapter Adapter) RemoveUser(userID bson.ObjectId) error {
	session := adapter.Session.Copy()
	session.SetSafe(&mgo.Safe{})
	defer session.Close()

	collection := session.DB(adapter.DBName).C("users")
	return collection.RemoveId(userID)
}
