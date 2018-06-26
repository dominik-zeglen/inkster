package mongodb

import (
	"github.com/dominik-zeglen/ecoknow/core"
	"github.com/globalsign/mgo"
)

// Adapter is an abstraction over mongodb session
type Adapter struct {
	core.Adapter

	ConnectionURI string
	DBName        string
}

func (adapter Adapter) String() string {
	return "MongoDB"
}

func (adapter Adapter) ResetMockDatabase(
	directories []core.Directory,
	templates []core.Template,
	pages []core.Page,
) {
	session, err := mgo.Dial(adapter.ConnectionURI)
	session.SetSafe(&mgo.Safe{})
	db := session.DB(adapter.DBName)

	collection := db.C("directories")
	collection.DropCollection()
	for id := range directories {
		err = collection.Insert(directories[id])
		if err != nil {
			panic(err)
		}
	}

	collection = db.C("templates")
	collection.DropCollection()
	for id := range templates {
		err = collection.Insert(templates[id])
		if err != nil {
			panic(err)
		}
	}

	collection = db.C("pages")
	collection.DropCollection()
	for id := range pages {
		err = collection.Insert(pages[id])
		if err != nil {
			panic(err)
		}
	}
}
