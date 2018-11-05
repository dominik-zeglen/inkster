package testing

import (
	"github.com/dominik-zeglen/inkster/core"
)

func createTemplate(template core.Template, id string, createdAt string, updatedAt string) core.Template {
	template.ID = id
	template.CreatedAt = createdAt
	template.UpdatedAt = updatedAt

	return template
}

func CreateTemplates() []core.Template {
	templates := []core.Template{
		core.Template{
			Name: "Template 1",
			Fields: []core.TemplateField{
				core.TemplateField{Type: "text", Name: "Field 1"},
				core.TemplateField{Type: "image", Name: "Field 2"},
			},
		},
		core.Template{
			Name: "Template 2",
			Fields: []core.TemplateField{
				core.TemplateField{Type: "unique", Name: "Field 3"},
				core.TemplateField{Type: "text", Name: "Field 4"},
				core.TemplateField{Type: "file", Name: "Field 5"},
			},
		},
	}

	templates[0] = createTemplate(
		templates[0],
		"000000000001",
		"2007-07-07T10:00:00.000Z",
		"2007-07-07T10:00:00.000Z",
	)
	templates[1] = createTemplate(
		templates[1],
		"000000000002",
		"2007-07-07T11:00:00.000Z",
		"2007-07-07T11:00:00.000Z",
	)

	return templates
}
