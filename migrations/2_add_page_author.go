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
				Exec("ALTER TABLE pages ADD COLUMN author_id BIGINT NOT NULL")
			if err != nil {
				return err
			}

			_, err = orm.
				NewQuery(db, nil).
				Exec("ALTER TABLE pages ADD CONSTRAINT pages_users_author_fk FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE")
			if err != nil {
				return err
			}

			return nil
		},
	)
}
