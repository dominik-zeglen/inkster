package core

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
)

// Template allows user to quickly create new pages without repeatedly
// adding the same fields each time
type Template struct {
	ID     bson.ObjectId   `bson:"_id,omitempty" json:"id"`
	Name   string          `json:"name"`
	Fields []TemplateField `json:"fields"`
}

// UpdateTemplateArguments is transactional model of an update properties
type UpdateTemplateArguments struct {
	Name string `bson:",omitempty"`
}

func (template Template) String() string {
	return fmt.Sprintf("Template<%s>", template.Name)
}

// Validate checks if template can be put into database
func (template Template) Validate() error {
	if template.Name == "" {
		return ErrNoEmpty("Name")
	}
	for fieldID := range template.Fields {
		err := template.Fields[fieldID].Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

// TemplateField represents a single field in template
type TemplateField struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

func (field TemplateField) String() string {
	return fmt.Sprintf("TemplateField<%s: %s>", field.Name, field.Type)
}

// Validate checks if field can be put into database
func (field TemplateField) Validate() error {
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
