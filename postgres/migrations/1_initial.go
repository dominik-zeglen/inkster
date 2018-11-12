package main

import (
	"github.com/dominik-zeglen/inkster/core"
	"github.com/go-pg/migrations"
	"github.com/go-pg/pg/orm"
)

func init() {
	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			for _, model := range []interface{}{
				(*core.Directory)(nil),
				(*core.Page)(nil),
				(*core.PageField)(nil),
				(*core.Template)(nil),
				(*core.User)(nil),
			} {
				query := orm.NewQuery(db, model)
				err := query.CreateTable(&orm.CreateTableOptions{
					FKConstraints: true,
				})
				if err != nil {
					return err
				}
			}

			_, err := orm.
				NewQuery(db, nil).
				Exec("ALTER TABLE page_fields ADD CONSTRAINT unique_page_fields_page_id_name UNIQUE (page_id, name)")
			if err != nil {
				return err
			}

			return nil
		},
	)
}
