package mongodb

import (
	"github.com/dominik-zeglen/inkster/core"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// AddDirectory puts directory in the database
func (adapter Adapter) AddDirectory(directory core.Directory) (core.Directory, error) {
	if directory.Name == "" {
		return core.Directory{}, core.ErrNoEmpty("Name")
	}
	session := adapter.Session.Copy()
	session.SetSafe(&mgo.Safe{})
	defer session.Close()

	collection := session.DB(adapter.DBName).C("directories")
	if directory.ID == "" {
		directory.ID = bson.NewObjectId()
	}
	directory.CreatedAt = adapter.GetCurrentTime()
	directory.UpdatedAt = adapter.GetCurrentTime()

	err := collection.Insert(directory)
	return directory, err
}

// GetDirectory gets directory from the database
func (adapter Adapter) GetDirectory(id bson.ObjectId) (core.Directory, error) {
	session := adapter.Session.Copy()
	defer session.Close()

	var directory core.Directory
	err := session.
		DB(adapter.DBName).
		C("directories").
		FindId(id).
		One(&directory)
	return directory, err
}

// GetDirectoryList gets all directories from the database
func (adapter Adapter) GetDirectoryList() ([]core.Directory, error) {
	session := adapter.Session.Copy()
	defer session.Close()

	var directories []core.Directory
	err := session.
		DB(adapter.DBName).
		C("directories").
		Find(bson.M{}).
		All(&directories)
	return directories, err
}

// GetRootDirectoryList gets only directories from a pg database that don't have parent
func (adapter Adapter) GetRootDirectoryList() ([]core.Directory, error) {
	session := adapter.Session.Copy()
	defer session.Close()

	var directories []core.Directory
	err := session.
		DB(adapter.DBName).
		C("directories").
		Find(bson.M{"parentId": bson.M{"$not": bson.M{"$type": 7}}}).
		All(&directories)
	return directories, err
}

// GetDirectoryChildrenList gets directories from a pg database which have same parent
func (adapter Adapter) GetDirectoryChildrenList(id bson.ObjectId) ([]core.Directory, error) {
	session := adapter.Session.Copy()
	defer session.Close()

	var directories []core.Directory
	err := session.
		DB(adapter.DBName).
		C("directories").
		Find(bson.M{"parentId": id}).
		All(&directories)
	return directories, err
}

type directoryUpdateInput struct {
	Data      core.DirectoryInput `bson:",inline"`
	UpdatedAt string              `bson:"updatedAt"`
}

// UpdateDirectory allows directory properties updaing
func (adapter Adapter) UpdateDirectory(
	directoryID bson.ObjectId,
	data core.DirectoryInput,
) error {
	session := adapter.Session.Copy()
	defer session.Close()

	collection := session.
		DB(adapter.DBName).
		C("directories")
	return collection.UpdateId(directoryID, bson.M{
		"$set": directoryUpdateInput{
			Data:      data,
			UpdatedAt: adapter.GetCurrentTime(),
		},
	})
}

// RemoveDirectory removes directory from a pg database
func (adapter Adapter) RemoveDirectory(id bson.ObjectId) error {
	session := adapter.Session.Copy()
	defer session.Close()

	err := session.
		DB(adapter.DBName).
		C("directories").
		RemoveId(id)
	return err
}
