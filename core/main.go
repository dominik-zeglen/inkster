package core

import (
	"time"

	"github.com/go-pg/pg"
)

// Adapter interface provides abstraction over different data sources
type AbstractDataContext interface {
	GetCurrentTime() string
	DB() *pg.DB
}

func CreateDirectory(dataContext AbstractDataContext) Directory {
	directory := Directory{}
	directory.CreatedAt = dataContext.GetCurrentTime()
	directory.UpdatedAt = dataContext.GetCurrentTime()

	return directory
}

// BaseModel is an abstraction over that all models
// should be composed of, providing most basic
// fields to keep order and consistency within code
type BaseModel struct {
	ID        int    `sql:",pk,autoincrement" json:"id"`
	CreatedAt string `sql:",notnull" json:"createdAt" bson:"createdAt"`
	UpdatedAt string `sql:",notnull" json:"updatedAt" bson:"updatedAt"`
}

type DataContext struct {
	AbstractDataContext

	Session *pg.DB
}

func (_ DataContext) GetCurrentTime() string {
	return time.Now().UTC().Format(time.RFC3339)
}

func (dataContext DataContext) DB() *pg.DB {
	return dataContext.Session
}

type MockContext struct {
	AbstractDataContext

	Session *pg.DB
}

func (_ MockContext) GetCurrentTime() string {
	return "2017-07-07T10:00:00.000Z"
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
