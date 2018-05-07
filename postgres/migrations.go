package postgres

import (
	"fmt"
	"os"

	"github.com/dominik-zeglen/ecoknow/core"
	"github.com/go-pg/pg"
)

func CheckMigrationIfApplied(dataSource Adapter, migrationID int32) bool {
	db := pg.Connect(&dataSource.ConnectionOptions)
	migration := core.Migration{ID: migrationID}
	err := db.Select(&migration)
	defer db.Close()
	if err != nil {
		return false
	}
	return true
}

func ApplyMigration(dataSource Adapter, ID int32, name string) error {
	db := pg.Connect(&dataSource.ConnectionOptions)
	defer db.Close()
	migration := core.Migration{ID: ID, Name: name}
	err := db.Insert(&migration)
	if err != nil {
		return err
	}
	return nil
}

// ApplyMigrations writes migrations to database
func ApplyMigrations(dataSource Adapter) error {
	migrations := []func(Adapter) error{
		InitDB,
	}

	for i, migration := range migrations {
		var err error
		if i != 0 {
			if CheckMigrationIfApplied(dataSource, int32(i)) {
				continue
			}
		}

		err = migration(dataSource)
		if err != nil {
			panic(err)
		}
		if os.Getenv("FOXXY_DEBUG") == "1" {
			fmt.Printf("Migration %d applied\n", i+1)
		}
	}

	if os.Getenv("FOXXY_DEBUG") == "1" {
		fmt.Println("Migrations properly applied")
	}

	return nil
}
