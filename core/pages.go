package core

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
)

// Page is a object representing site content
type Page struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name     string        `json:"name"`
	Slug     string        `json:"slug"`
	ParentID bson.ObjectId `bson:"parentId" json:"parentId"`
	Fields   []PageField   `json:"fields"`
}

func (page Page) String() string {
	return fmt.Sprintf("Page<%s: %s>", page.Name, page.Slug)
}

// Validate checks if page can be put into database
func (page Page) Validate() error {
	if page.Name == "" {
		return ErrNoEmpty("Name")
	}
	if page.ParentID == "" {
		return ErrNoEmpty("ParentID")
	}
	if page.Fields != nil {
		for fieldIndex, field := range page.Fields {
			for comparisonFieldIndex, comparisonField := range page.Fields {
				if fieldIndex == comparisonFieldIndex {
					continue
				}
				if field.Name == comparisonField.Name {
					return ErrFieldExists(field.Name)
				}
			}
		}
	}
	return nil
}

// UpdatePageArguments is transactional model of an update properties
type UpdatePageArguments struct {
	Name string `bson:",omitempty"`
	Slug string `bson:",omitempty"`
}

// PageField represents a single field in template
type PageField struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Validate checks if field can be put into database
func (field PageField) Validate() error {
	if field.Name == "" {
		return ErrNoEmpty("Name")
	}
	if field.Type == "" {
		return ErrNoEmpty("Type")
	}
	found := false
	for fieldTypeID := range FieldTypes {
		if FieldTypes[fieldTypeID] == field.Type {
			found = true
		}
	}
	if !found {
		return ErrNoFieldType(field.Type)
	}
	return nil
}

func (field PageField) String() string {
	return fmt.Sprintf("PageField<%s: %s (%s)>",
		field.Name,
		field.Type,
		field.Value[:20],
	)
}

// PageInput is transactional model of an creation properties
type PageInput struct {
	Name     *string        `bson:",omitempty"`
	Slug     *string        `bson:",omitempty"`
	ParentID *bson.ObjectId `bson:",omitempty"`
	Fields   *[]PageField   `bson:",omitempty"`
}

// UpdatePageFieldArguments is transactional model of an update properties
type UpdatePageFieldArguments struct {
	Name  string `bson:",omitempty"`
	Value string `bson:",omitempty"`
}
