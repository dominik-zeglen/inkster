package migrations

import (
	"github.com/dominik-zeglen/foxxy/core"
	"github.com/go-pg/pg"
)

func InitDB() error {
	db := pg.Connect(core.DbOptions)
	defer db.Close()

	_, err := core.CheckMigrationIfApplied(1)
	if err == nil {
		return nil
	}

	err = db.CreateTable(&core.Container{}, nil)
	if err != nil {
		return err
	}

	err = db.CreateTable(&core.Migration{}, nil)
	if err != nil {
		return err
	}

	err = core.ApplyMigration(1, "init_db")
	if err != nil {
		return err
	}
	return nil
}
