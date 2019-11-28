package main

import (
	"encoding/json"
	"time"

	"github.com/go-pg/migrations"
	"github.com/go-pg/pg/orm"
)

type pageType struct {
	tableName   struct{}  `sql:"pages"`
	ID          int       `sql:",pk,autoincrement" json:"id"`
	CreatedAt   time.Time `sql:",notnull" json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time `sql:",notnull" json:"updatedAt" bson:"updatedAt"`
	AuthorID    int       `sql:",notnull" json:"authorId" validate:"required"`
	Name        string    `sql:",notnull" json:"name" validate:"required,min=3"`
	Slug        string    `sql:",notnull" json:"slug" validate:"omitempty,slug,min=3"`
	ParentID    int       `sql:",notnull" json:"parentId" validate:"required"`
	IsPublished bool      `sql:",notnull" json:"isPublished"`
}
type pageFieldType struct {
	tableName struct{}  `sql:"page_fields"`
	ID        int       `sql:",pk,autoincrement" json:"id"`
	CreatedAt time.Time `sql:",notnull" json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `sql:",notnull" json:"updatedAt" bson:"updatedAt"`
	PageID    int       `sql:",notnull,on_delete:CASCADE" json:"-"`
	Type      string    `json:"type" validate:"required,oneof=text longText image file"`
	Name      string    `json:"name" validate:"required,min=3"`
	Value     string    `json:"value"`
}
type jsonPageFieldType struct {
	Type  string `json:"type" validate:"required,oneof=text longText image file"`
	Slug  string `json:"name" validate:"required,slug"`
	Value string `json:"value"`
}

func init() {
	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			var pageFields []pageFieldType

			err := orm.
				NewQuery(db, &pageFields).
				Select()

			if err != nil {
				return err
			}

			fieldMap := make(map[int]([]jsonPageFieldType))

			for _, field := range pageFields {
				fieldMap[field.PageID] = append(
					fieldMap[field.PageID],
					jsonPageFieldType{
						Type:  field.Type,
						Slug:  field.Name,
						Value: field.Value,
					})
			}

			_, err = orm.
				NewQuery(db, nil).
				Exec("ALTER TABLE pages ADD COLUMN fields jsonb not null default '[]'::jsonb")

			if err != nil {
				return err
			}

			for id, fields := range fieldMap {
				jsonFields, _ := json.Marshal(fields)
				_, err := orm.
					NewQuery(db, nil).
					Exec(
						"UPDATE pages SET fields = ? WHERE id = ?",
						string(jsonFields),
						id,
					)

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
