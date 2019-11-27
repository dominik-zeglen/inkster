package main

import (
	"time"

	"github.com/go-pg/migrations"
	"github.com/go-pg/pg/orm"
)

type Page struct {
	ID          int         `sql:",pk,autoincrement" json:"id"`
	CreatedAt   time.Time   `sql:",notnull" json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time   `sql:",notnull" json:"updatedAt" bson:"updatedAt"`
	AuthorID    int         `sql:",notnull" json:"authorId" validate:"required"`
	Name        string      `sql:",notnull" json:"name" validate:"required,min=3"`
	Slug        string      `sql:",notnull" json:"slug" validate:"omitempty,slug,min=3"`
	ParentID    int         `sql:",notnull" json:"parentId" validate:"required"`
	IsPublished bool        `sql:",notnull" json:"isPublished"`
	Fields      []PageField `json:"fields" sql:"type:jsonb" validate:"dive"`
}
type PageField struct {
	ID        int       `sql:",pk,autoincrement" json:"id"`
	CreatedAt time.Time `sql:",notnull" json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `sql:",notnull" json:"updatedAt" bson:"updatedAt"`
	PageID    int       `sql:",notnull,on_delete:CASCADE" json:"-"`
	Type      string    `json:"type" validate:"required,oneof=text longText image file"`
	Slug      string    `json:"name" validate:"required,slug"`
	Value     string    `json:"value"`
}

func init() {
	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			_, err := orm.
				NewQuery(db, nil).
				Exec("ALTER TABLE pages ADD COLUMN page_fields jsonb not null default '[]'::jsonb")

			if err != nil {
				return err
			}

			var pageFields []PageField

			_, err = orm.
				NewQuery(db, pageFields).
				Exec("SELECT * FROM page_fields")

			if err != nil {
				return err
			}

			fieldMap := make(map[int]([]PageField))

			for _, field := range pageFields {
				fieldMap[field.PageID] = append(fieldMap[field.PageID], field)
			}

			for id, fields := range fieldMap {
				page := Page{
					ID:     id,
					Fields: fields,
				}

				_, err := orm.
					NewQuery(db, nil).
					Model(&page).
					WherePK().
					UpdateNotNull()

				if err != nil {
					return err
				}
			}

			_, err = orm.
				NewQuery(db, nil).
				Exec("DROP TABLE page_fields")

			return err
		},
	)
}
