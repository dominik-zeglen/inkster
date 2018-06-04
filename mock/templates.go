package mock

import (
	"fmt"

	"github.com/dominik-zeglen/ecoknow/core"
	"github.com/globalsign/mgo/bson"
)

func (adapter Adapter) findTemplate(id *bson.ObjectId, name *string) (int, error) {
	if id != nil {
		for index := range adapter.templates {
			if adapter.templates[index].ID == *id {
				return index, nil
			}
		}
		return 0, fmt.Errorf("Template %s does not exist", id)
	}
	if name != nil {
		for index := range adapter.templates {
			if adapter.templates[index].Name == *name {
				return index, nil
			}
		}
		return 0, fmt.Errorf("Template %s does not exist", *name)
	}
	if id == nil && name == nil {
		return 0, fmt.Errorf("findTemplate() requires at least one argument")
	}
	return 0, fmt.Errorf("")
}
func (adapter Adapter) findTemplateField(id bson.ObjectId, name string) (int, int, error) {
	index, err := adapter.findTemplate(&id, nil)
	if err != nil {
		return 0, 0, err
	}
	for fieldIndex := range adapter.templates[index].Fields {
		if adapter.templates[index].Fields[fieldIndex].Name == name {
			return index, fieldIndex, nil
		}
	}
	return 0, 0, core.ErrNoField(name)
}

// AddTemplate puts template in the database
func (adapter Adapter) AddTemplate(template core.Template) (core.Template, error) {
	err := template.Validate()
	if err != nil {
		return core.Template{}, err
	}
	_, err = adapter.findTemplate(nil, &template.Name)
	if err == nil {
		return core.Template{}, core.ErrTemplateExists(template.Name)
	}
	if template.ID == "" {
		template.ID = bson.NewObjectId()
	} else {
		_, err = adapter.findTemplate(&template.ID, nil)
		if err == nil {
			return core.Template{}, core.ErrTemplateExists(template.ID.String())
		}
	}
	adapter.templates = append(adapter.templates, template)
	return template, nil
}

// AddTemplateField adds to template a new field at the end of it's field list
func (adapter Adapter) AddTemplateField(templateID bson.ObjectId, field core.TemplateField) error {
	err := field.Validate()
	if err != nil {
		return err
	}

	index, _, err := adapter.findTemplateField(templateID, field.Name)
	if err == nil {
		return core.ErrFieldExists(field.Name)
	}
	adapter.templates[index].Fields = append(
		adapter.templates[index].Fields,
		field,
	)
	return nil
}

// GetTemplate allows user to fetch template from database
func (adapter Adapter) GetTemplate(templateID bson.ObjectId) (core.Template, error) {
	index, err := adapter.findTemplate(&templateID, nil)
	return adapter.templates[index], err
}

// GetTemplateList allows user to fetch all templates from database
func (adapter Adapter) GetTemplateList() ([]core.Template, error) {
	return adapter.templates, nil
}

// UpdateTemplate allows user to update template properties
func (adapter Adapter) UpdateTemplate(templateID bson.ObjectId, data core.TemplateInput) error {
	index, err := adapter.findTemplate(nil, &data.Name)
	if err == nil {
		return core.ErrTemplateExists(data.Name)
	}
	adapter.templates[index].Name = data.Name
	return nil
}

// RemoveTemplate removes template from database
func (adapter Adapter) RemoveTemplate(templateID bson.ObjectId) error {
	index, err := adapter.findTemplate(&templateID, nil)
	if err != nil {
		return err
	}
	adapter.templates = append(adapter.templates[:index], adapter.templates[:index+1]...)
	return nil
}

// RemoveTemplateField removes field from template
func (adapter Adapter) RemoveTemplateField(templateID bson.ObjectId, templateFieldName string) error {
	index, fieldIndex, err := adapter.findTemplateField(templateID, templateFieldName)
	if err != nil {
		return err
	}
	adapter.templates[index].Fields = append(
		adapter.templates[index].Fields[:fieldIndex],
		adapter.templates[index].Fields[fieldIndex+1:]...,
	)
	return nil
}
