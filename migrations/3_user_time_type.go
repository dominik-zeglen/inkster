package main

import (
	"fmt"

	"github.com/go-pg/migrations"
	"github.com/go-pg/pg/orm"
)

func init() {
	query := `
	ALTER TABLE %s 
  		ALTER COLUMN %s 
   			TYPE TIMESTAMP WITH TIME ZONE 
     			USING to_timestamp(%s, 'YYYY-MM-DD HH24:MI:SS');`
	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			fields := []string{"created_at", "updated_at"}
			tables := []string{"pages", "directories", "users"}

			for _, field := range fields {
				for _, table := range tables {
					_, err := orm.
						NewQuery(db, nil).
						Exec(fmt.Sprintf(
							query,
							table,
							field,
							field,
						))
					if err != nil {
						return err
					}
				}
			}

			return nil
		},
	)
}
