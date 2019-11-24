package main

import (
	"github.com/go-pg/migrations"
	"github.com/go-pg/pg/orm"
)

func init() {
	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			_, err := orm.
				NewQuery(db, nil).
				Exec("ALTER TABLE pages DROP CONSTRAINT pages_slug_key")

			if err != nil {
				return err
			}

			return nil
		},
	)
}
