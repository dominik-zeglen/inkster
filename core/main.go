package core

import (
	"time"

	"github.com/go-pg/pg"
)

// Adapter interface provides abstraction over different data sources
type AbstractDataContext interface {
	GetCurrentTime() time.Time
	DB() *pg.DB
}

// BaseModel is an abstraction over that all models
// should be composed of, providing most basic
// fields to keep order and consistency within code
type BaseModel struct {
	ID        int       `sql:",pk,autoincrement" json:"id"`
	CreatedAt time.Time `sql:",notnull" json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `sql:",notnull" json:"updatedAt" bson:"updatedAt"`
}

type DataContext struct {
	AbstractDataContext

	Session *pg.DB
}

func (_ DataContext) GetCurrentTime() time.Time {
	return time.Now()
}

func (dataContext DataContext) DB() *pg.DB {
	return dataContext.Session
}

type MockContext struct {
	AbstractDataContext

	Session *pg.DB
}

func (_ MockContext) GetCurrentTime() time.Time {
	output, _ := time.Parse(
		"2006-01-02T15:04:05.000Z",
		"2017-07-07T10:00:00.000Z",
	)
	return output
}

func (dataContext MockContext) DB() *pg.DB {
	return dataContext.Session
}

// FieldTypes holds all allowed template field type names
var FieldTypes = []string{
	"directory",
	"file",
	"image",
	"longText",
	"page",
	"text",
	"unique",
}
