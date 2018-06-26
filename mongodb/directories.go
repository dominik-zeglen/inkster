package mongodb

import (
	"github.com/dominik-zeglen/ecoknow/core"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// AddDirectory puts directory in the database
func (adapter Adapter) AddDirectory(directory core.Directory) (core.Directory, error) {
	if directory.Name == "" {
		return core.Directory{}, core.ErrNoEmpty("Name")
	}
	db, err := mgo.Dial(adapter.ConnectionURI)
	db.SetSafe(&mgo.Safe{})
	defer db.Close()
	if err != nil {
		return core.Directory{}, err
	}
	collection := db.DB(adapter.DBName).C("directories")
	if directory.ID == "" {
		directory.ID = bson.NewObjectId()
	}
	err = collection.Insert(directory)
	return directory, err
}

// GetDirectory gets directory from the database
func (adapter Adapter) GetDirectory(id bson.ObjectId) (core.Directory, error) {
	db, err := mgo.Dial(adapter.ConnectionURI)
	defer db.Close()
	if err != nil {
		return core.Directory{}, err
	}
	var directory core.Directory
	err = db.
		DB(adapter.DBName).
		C("directories").
		FindId(id).
		One(&directory)
	return directory, err
}

// GetDirectoryList gets all directories from the database
func (adapter Adapter) GetDirectoryList() ([]core.Directory, error) {
	db, err := mgo.Dial(adapter.ConnectionURI)
	defer db.Close()
	if err != nil {
		return nil, err
	}
	var directories []core.Directory
	err = db.
		DB(adapter.DBName).
		C("directories").
		Find(bson.M{}).
		All(&directories)
	return directories, err
}

// GetRootDirectoryList gets only directories from a pg database that don't have parent
func (adapter Adapter) GetRootDirectoryList() ([]core.Directory, error) {
	db, err := mgo.Dial(adapter.ConnectionURI)
	defer db.Close()
	if err != nil {
		return nil, err
	}
	var directories []core.Directory
	err = db.
		DB(adapter.DBName).
		C("directories").
		Find(bson.M{"parentId": bson.M{"$not": bson.M{"$type": 7}}}).
		All(&directories)
	return directories, err
}

// GetDirectoryChildrenList gets directories from a pg database which have same parent
func (adapter Adapter) GetDirectoryChildrenList(id bson.ObjectId) ([]core.Directory, error) {
	db, err := mgo.Dial(adapter.ConnectionURI)
	defer db.Close()
	if err != nil {
		return nil, err
	}
	var directories []core.Directory
	err = db.
		DB(adapter.DBName).
		C("directories").
		Find(bson.M{"parentId": id}).
		All(&directories)
	return directories, err
}

// UpdateDirectory allows directory properties updaing
func (adapter Adapter) UpdateDirectory(
	directoryID bson.ObjectId,
	data core.DirectoryInput,
) error {
	db, err := mgo.Dial(adapter.ConnectionURI)
	defer db.Close()
	if err != nil {
		return err
	}
	collection := db.
		DB(adapter.DBName).
		C("directories")
	return collection.UpdateId(directoryID, bson.M{
		"$set": data,
	})
}

// RemoveDirectory removes directory from a pg database
func (adapter Adapter) RemoveDirectory(id bson.ObjectId) error {
	db, err := mgo.Dial(adapter.ConnectionURI)
	defer db.Close()
	if err != nil {
		return err
	}
	err = db.
		DB(adapter.DBName).
		C("directories").
		RemoveId(id)
	return err
}
