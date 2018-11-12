package postgres

import (
	"github.com/dominik-zeglen/inkster/core"
)

// AddTemplate puts template in the database
func (adapter Adapter) AddTemplate(template core.Template) (core.Template, error) {
	err := template.Validate()
	if err != nil {
		return core.Template{}, err
	}

	template.CreatedAt = adapter.GetCurrentTime()
	template.UpdatedAt = adapter.GetCurrentTime()

	return template, nil
}

// AddTemplateField adds to template a new field at the end of it's field list
func (adapter Adapter) AddTemplateField(templateID int, field core.TemplateField) error {
	return nil
}

// GetTemplate allows user to fetch template from database
func (adapter Adapter) GetTemplate(templateID int) (core.Template, error) {
	template := core.Template{}
	return template, nil
}

// GetTemplateList allows user to fetch all templates from database
func (adapter Adapter) GetTemplateList() ([]core.Template, error) {
	templates := []core.Template{}
	return templates, nil
}

type templateUpdateInput struct {
	Data      core.TemplateInput `bson:",inline"`
	UpdatedAt string             `bson:"updatedAt"`
}

// UpdateTemplate allows user to update template properties
func (adapter Adapter) UpdateTemplate(templateID int, data core.TemplateInput) error {
	return core.ErrBadCredentials
}

// RemoveTemplate removes template from database
func (adapter Adapter) RemoveTemplate(templateID int) error {
	return core.ErrBadCredentials
}

// RemoveTemplateField removes field from template
func (adapter Adapter) RemoveTemplateField(templateID int, templateFieldName string) error {
	return nil
}
