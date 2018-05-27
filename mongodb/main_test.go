package mongodb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/dominik-zeglen/ecoknow/core"
	"github.com/globalsign/mgo"
)

var dataSource = Adapter{
	ConnectionURI: os.Getenv("FOXXY_DB_URI"),
	DBName:        os.Getenv("FOXXY_DB_NAME") + "_test",
}

var ErrNoError = fmt.Errorf("Did not return error")

// ToJSON is handy snippet for pretty-formatting json snapshots
func ToJSON(object interface{}) (string, error) {
	data, err := json.Marshal(&object)
	if err != nil {
		return "", err
	}
	var out bytes.Buffer
	json.Indent(&out, data, "", "    ")
	return out.String(), nil
}

func resetDatabase() {
	containers := []core.Container{
		core.Container{ID: "000000000001", Name: "Container 1"},
		core.Container{ID: "000000000002", Name: "Container 2"},
		core.Container{ID: "000000000003", Name: "Container 3"},
		core.Container{ID: "000000000004", Name: "Container 1.1", ParentID: "000000000001"},
	}
	session, err := mgo.Dial(dataSource.ConnectionURI)
	session.SetSafe(&mgo.Safe{})
	db := session.DB(dataSource.DBName)
	collection := db.C("containers")
	collection.DropCollection()
	for id := range containers {
		err = collection.Insert(containers[id])
		if err != nil {
			panic(err)
		}
	}

	templates := []core.Template{
		core.Template{
			ID:   "000000000001",
			Name: "Template 1",
			Fields: []core.TemplateField{
				core.TemplateField{Type: "text", Name: "Field 1"},
				core.TemplateField{Type: "image", Name: "Field 2"},
			},
		},
		core.Template{
			ID:   "000000000002",
			Name: "Template 2",
			Fields: []core.TemplateField{
				core.TemplateField{Type: "unique", Name: "Field 3"},
				core.TemplateField{Type: "text", Name: "Field 4"},
				core.TemplateField{Type: "file", Name: "Field 5"},
			},
		},
	}
	collection = db.C("templates")
	collection.DropCollection()
	for id := range templates {
		err = collection.Insert(templates[id])
		if err != nil {
			panic(err)
		}
	}
}
