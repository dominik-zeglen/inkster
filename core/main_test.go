package core

import (
	"database/sql"
	"regexp"

	"github.com/dgrijalva/jwt-go"
	"github.com/dominik-zeglen/inkster/config"
	"github.com/go-pg/pg"
	"github.com/go-testfixtures/testfixtures"
	"github.com/lib/pq"
)

var conf config.Config
var dataSource MockContext
var fixtures *testfixtures.Context

func init() {
	conf = *config.Load("../")
	dbHost := conf.Postgres.URI
	pgOptions, err := pg.ParseURL(dbHost)
	if err != nil {
		panic(err)
	}
	if !conf.Miscellaneous.CI {
		pgOptions.Database = "test_" + pgOptions.Database
		dbOptions, err := pq.ParseURL(dbHost)
		if err != nil {
			panic(err)
		}
		regex := regexp.MustCompile("dbname=")
		replace := "${1}dbname=test_$2"
		dbHost = regex.ReplaceAllString(dbOptions, replace)
	}

	pgSession := pg.Connect(pgOptions)
	pgAdapter := MockContext{
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

	jwt.TimeFunc = dataSource.GetCurrentTime

	resetDatabase()
}

func resetDatabase() {
	if err := fixtures.Load(); err != nil {
		panic(err)
	}
}
