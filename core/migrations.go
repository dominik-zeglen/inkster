package core

import (
	"time"

	"github.com/go-pg/pg"
)

type Migration struct {
	Id   int32
	Name string
	Date int32
}

func ApplyMigration(id int32, name string) error {
	db := pg.Connect(DbOptions)
	defer db.Close()
	date := int32(time.Now().Unix())
	return db.Insert(&Migration{Id: id, Name: name, Date: date})
}

func CheckMigrationIfApplied(id int32) (bool, error) {
	db := pg.Connect(DbOptions)
	defer db.Close()
	migration := Migration{Id: id}
	err := db.Select(&migration)
	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			return true, nil
		}
		return false, err
	}
	return false, nil
}
