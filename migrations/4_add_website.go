package main

import (
	"github.com/dominik-zeglen/inkster/app"
	"github.com/go-pg/migrations"
	"github.com/go-pg/pg/orm"
)

func init() {
	type Website struct {
		ID          string
		Name        string `sql:",notnull"`
		Description string
		Domain      string `sql:",notnull"`
	}

	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			for _, model := range []interface{}{
				(*Website)(nil),
			} {
				query := orm.NewQuery(db, model)
				err := query.CreateTable(&orm.CreateTableOptions{
					FKConstraints: true,
				})
				if err != nil {
					return err
				}

				config, err := app.LoadConfig("config.toml")
				if err != nil {
					return err
				}

				website := Website{
					ID:     "default",
					Domain: "http://localhost:" + config.Server.Port,
					Name:   "Inkster",
				}

				_, err = db.Model(&website).Insert()
				if err != nil {
					return err
				}
			}

			return nil
		},
	)
}
