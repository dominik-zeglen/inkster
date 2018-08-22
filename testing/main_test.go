package testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/dominik-zeglen/inkster/core"
	"github.com/dominik-zeglen/inkster/mock"
	"github.com/dominik-zeglen/inkster/mongodb"
	"github.com/globalsign/mgo"
)

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

var dataSources = []core.Adapter{
	mongodb.Adapter{},
	mock.Adapter{
		GetTime: func() string { return "2017-07-07T10:00:00.000Z" },
	},
}
var dataSource = dataSources[0]

func resetDatabase() {
	dataSource.ResetMockDatabase(CreateDirectories(), CreateTemplates(), CreatePages())
}

func TestMain(t *testing.T) {
	mongoAdapter := mongodb.Adapter{
		DBName:  os.Getenv("INKSTER_DB_NAME") + "_test",
		GetTime: func() string { return "2017-07-07T10:00:00.000Z" },
	}
	session, err := mgo.Dial(os.Getenv("INKSTER_DB_URI"))
	if err != nil {
		t.Fatal(err)
	}
	defer session.Close()
	mongoAdapter.Session = session
	dataSources[0] = mongoAdapter
	for i := range dataSources {
		dataSource = dataSources[i]
		resetDatabase()
		t.Log(fmt.Sprintf("Testing adapter %s...\n", dataSource.String()))
		t.Run("Test directories", testDirectories)
		t.Run("Test templates", testTemplates)
		t.Run("Test pages", testPages)
	}
}
