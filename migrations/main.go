package migrations

import (
	"fmt"
	"os"

	"github.com/dominik-zeglen/foxxy/core"
)

func ApplyMigrations() error {
	migrations := []func() error{
		InitDB,
	}

	for i, migration := range migrations {
		var err error
		if i != 0 {
			var applied bool

			applied, err = core.CheckMigrationIfApplied(int32(i))

			if err != nil {
				return err
			}
			if applied {
				continue
			}
		}

		err = migration()
		if err != nil {
			return err
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
