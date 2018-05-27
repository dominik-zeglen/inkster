package mongodb

import (
	"github.com/dominik-zeglen/ecoknow/core"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// AddContainer puts container in the database
func (adapter *Adapter) AddContainer(container core.Container) (core.Container, error) {
	if container.Name == "" {
		return core.Container{}, core.ErrNoEmpty("Name")
	}
	db, err := mgo.Dial(adapter.ConnectionURI)
	db.SetSafe(&mgo.Safe{})
	defer db.Close()
	if err != nil {
		return core.Container{}, err
	}
	collection := db.DB(adapter.DBName).C("containers")
	if container.ID == "" {
		container.ID = bson.NewObjectId()
	}
	err = collection.Insert(container)
	return container, err
}

// GetContainer gets container from the database
func (adapter *Adapter) GetContainer(id bson.ObjectId) (core.Container, error) {
	db, err := mgo.Dial(adapter.ConnectionURI)
	defer db.Close()
	if err != nil {
		return core.Container{}, err
	}
	var container core.Container
	err = db.
		DB(adapter.DBName).
		C("containers").
		FindId(id).
		One(&container)
	return container, err
}

// GetContainerList gets all containers from the database
func (adapter *Adapter) GetContainerList() ([]core.Container, error) {
	db, err := mgo.Dial(adapter.ConnectionURI)
	defer db.Close()
	if err != nil {
		return nil, err
	}
	var containers []core.Container
	err = db.
		DB(adapter.DBName).
		C("containers").
		Find(bson.M{}).
		All(&containers)
	return containers, err
}

// GetRootContainerList gets only containers from a pg database that don't have parent
func (adapter *Adapter) GetRootContainerList() ([]core.Container, error) {
	db, err := mgo.Dial(adapter.ConnectionURI)
	defer db.Close()
	if err != nil {
		return nil, err
	}
	var containers []core.Container
	err = db.
		DB(adapter.DBName).
		C("containers").
		Find(bson.M{"parentId": bson.M{"$not": bson.M{"$type": 7}}}).
		All(&containers)
	return containers, err
}

// GetContainerChildrenList gets containers from a pg database which have same parent
func (adapter *Adapter) GetContainerChildrenList(id bson.ObjectId) ([]core.Container, error) {
	db, err := mgo.Dial(adapter.ConnectionURI)
	defer db.Close()
	if err != nil {
		return nil, err
	}
	var containers []core.Container
	err = db.
		DB(adapter.DBName).
		C("containers").
		Find(bson.M{"parentId": id}).
		All(&containers)
	return containers, err
}

// UpdateContainer allows container properties updaing
func (adapter *Adapter) UpdateContainer(
	containerID bson.ObjectId,
	data core.UpdateContainerArguments,
) error {
	db, err := mgo.Dial(adapter.ConnectionURI)
	defer db.Close()
	if err != nil {
		return err
	}
	collection := db.
		DB(adapter.DBName).
		C("containers")
	return collection.UpdateId(containerID, bson.M{
		"$set": data,
	})
}

// RemoveContainer removes container from a pg database
func (adapter *Adapter) RemoveContainer(id bson.ObjectId) error {
	db, err := mgo.Dial(adapter.ConnectionURI)
	defer db.Close()
	if err != nil {
		return err
	}
	err = db.
		DB(adapter.DBName).
		C("containers").
		RemoveId(id)
	return err
}
