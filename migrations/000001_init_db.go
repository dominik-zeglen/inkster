package migrations

import (
	"fmt"

	"github.com/dominik-zeglen/ecoknow/core"
	"github.com/dominik-zeglen/ecoknow/postgres"
	"github.com/go-pg/pg"
)

func InitDB(adapter postgres.Adapter) error {
	db := pg.Connect(&adapter.ConnectionOptions)
	defer db.Close()

	applied := CheckMigrationIfApplied(adapter, 1)
	if applied {
		fmt.Println("Migration applied, skipping...")
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
