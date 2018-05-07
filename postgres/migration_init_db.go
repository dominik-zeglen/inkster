package postgres

import (
	"github.com/dominik-zeglen/ecoknow/core"
	"github.com/go-pg/pg"
)

func InitDB(adapter Adapter) error {
	db := pg.Connect(&adapter.ConnectionOptions)
	defer db.Close()

	applied := CheckMigrationIfApplied(adapter, 1)
	if applied {
		return nil
	}

	err := db.CreateTable(&core.Container{}, nil)
	if err != nil {
		return err
	}

	err = db.CreateTable(&core.Migration{}, nil)
	if err != nil {
		return err
	}

	err = ApplyMigration(adapter, 1, "init_db")
	if err != nil {
		return err
	}
	return nil
}
