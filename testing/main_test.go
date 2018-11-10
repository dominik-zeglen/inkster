package testing

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"os"
	"regexp"
	"testing"

	"github.com/dominik-zeglen/inkster/core"
	"github.com/dominik-zeglen/inkster/postgres"
	"github.com/go-pg/pg"
	"github.com/go-testfixtures/testfixtures"
)

var ErrNoError = fmt.Errorf("Did not return error")
var fixtures *testfixtures.Context

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
	if err := fixtures.Load(); err != nil {
		panic(err)
	}
}

var dataSource core.Adapter

func TestMain(t *testing.T) {
	dbHost := os.Getenv("POSTGRES_HOST")
	pgOptions, err := pg.ParseURL(dbHost)
	if err != nil {
		panic(err)
	}
	if os.Getenv("CI") == "" {
		pgOptions.Database = "test_" + pgOptions.Database
		dbOptions, err := pq.ParseURL(os.Getenv("POSTGRES_HOST"))
		if err != nil {
			panic(err)
		}
		regex := regexp.MustCompile("dbname=")
		replace := "${1}dbname=test_$2"
		dbHost = regex.ReplaceAllString(dbOptions, replace)
	}

	pgSession := pg.Connect(pgOptions)
	pgAdapter := postgres.Adapter{
		GetTime: func() string { return "2017-07-07T10:00:00.000Z" },
		Session: pgSession,
	}
	dataSource = pgAdapter

	db, err := sql.Open("postgres", dbHost)
	if err != nil {
		panic(err)
	}

	fixtures, err = testfixtures.NewFolder(
		db,
		&testfixtures.PostgreSQL{},
		"../fixtures",
	)
	if err != nil {
		panic(err)
	}

	resetDatabase()

	t.Log(fmt.Sprintf("Testing adapter %s...\n", dataSource.String()))
	t.Run("Test directories", testDirectories)
	t.Run("Test pages", testPages)
	t.Run("Test users", testUsers)
}
