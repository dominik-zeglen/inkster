package testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/dominik-zeglen/ecoknow/core"
	"github.com/dominik-zeglen/ecoknow/mock"
	"github.com/dominik-zeglen/ecoknow/mongodb"
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
	mongodb.Adapter{
		ConnectionURI: os.Getenv("FOXXY_DB_URI"),
		DBName:        os.Getenv("FOXXY_DB_NAME") + "_test",
	},
	mock.Adapter{},
}
var dataSource = dataSources[0]

func resetDatabase() {
	dataSource.ResetMockDatabase(Directories, Templates, Pages)
}

func TestMain(t *testing.T) {
	for i := range dataSources {
		dataSource = dataSources[i]
		t.Log(fmt.Sprintf("Testing adapter %s...\n", dataSource.String()))
		t.Run("Test directories", testDirectories)
		t.Run("Test templates", testTemplates)
		t.Run("Test pages", testPages)
	}
}
