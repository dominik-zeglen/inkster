package testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/dominik-zeglen/inkster/core"
	"github.com/dominik-zeglen/inkster/postgres"
	"github.com/go-pg/pg"
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

func resetDatabase() {
	dataSource.ResetMockDatabase(
		CreateDirectories(),
		CreateTemplates(),
		CreatePages(),
		CreateUsers(),
	)
}

var dataSource core.Adapter

func TestMain(t *testing.T) {
	pgSession := pg.Connect(&pg.Options{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: "test_" + os.Getenv("POSTGRES_DB"),
	})
	pgAdapter := postgres.Adapter{
		GetTime: func() string { return "2017-07-07T10:00:00.000Z" },
		Session: pgSession,
	}
	dataSource = pgAdapter

	resetDatabase()

	t.Log(fmt.Sprintf("Testing adapter %s...\n", dataSource.String()))
	t.Run("Test directories", testDirectories)
	t.Run("Test pages", testPages)
	t.Run("Test users", testUsers)
}
