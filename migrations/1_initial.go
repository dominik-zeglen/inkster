package main

import (
	"github.com/go-pg/migrations"
	"github.com/go-pg/pg/orm"
)

func init() {
	type BaseModel struct {
		ID        int    `sql:",pk,autoincrement"`
		CreatedAt string `sql:",notnull"`
		UpdatedAt string `sql:",notnull"`
	}

	type Directory struct {
		BaseModel
		Name        string `sql:",notnull"`
		ParentID    int    `sql:",on_delete:CASCADE"`
		Parent      *Directory
		IsPublished bool `sql:",notnull"`
	}

	type Page struct {
		BaseModel
		Name        string `sql:",notnull"`
		Slug        string `sql:",unique,notnull"`
		ParentID    int    `sql:",notnull,on_delete:CASCADE"`
		Parent      *Directory
		IsPublished bool `sql:",notnull"`
	}

	type PageField struct {
		BaseModel
		Page   *Page
		PageID int `sql:",notnull,on_delete:CASCADE"`
		Type   string
		Name   string
		Value  string
	}

	type User struct {
		BaseModel
		Active   bool   `sql:",notnull"`
		Email    string `sql:",notnull,unique"`
		Password []byte `sql:",notnull"`
		Salt     []byte `sql:",notnull"`
	}

	type TemplateField struct {
		Type string
		Name string
	}

	type Template struct {
		BaseModel
		Name   string `sql:",notnull,unique"`
		Fields []TemplateField
	}

	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			for _, model := range []interface{}{
				(*Directory)(nil),
				(*Page)(nil),
				(*PageField)(nil),
				(*Template)(nil),
				(*TemplateField)(nil),
				(*User)(nil),
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
