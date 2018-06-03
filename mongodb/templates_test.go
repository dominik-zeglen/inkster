package mongodb

import (
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/dominik-zeglen/ecoknow/core"
)

func init() {
	resetDatabase()
}

func TestTemplates(t *testing.T) {
	t.Run("Test setters", func(t *testing.T) {
		t.Run("Add template", func(t *testing.T) {
			defer resetDatabase()
			template := core.Template{
				Name: "New Template",
				Fields: []core.TemplateField{
					core.TemplateField{Type: "unique", Name: "Field 6"},
					core.TemplateField{Type: "text", Name: "Field 7"},
				},
			}
			result, err := dataSource.AddTemplate(template)
			result.ID = ""
			if err != nil {
				t.Fatal(err)
			}
			data, err := ToJSON(result)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotT(t, data)
		})
		t.Run("Add template without name", func(t *testing.T) {
			defer resetDatabase()
			template := core.Template{}
			_, err := dataSource.AddTemplate(template)
			if err == nil {
				t.Error(ErrNoError)
			}
		})
		t.Run("Add template with duplicated field names", func(t *testing.T) {
			defer resetDatabase()
			template := core.Template{
				Name: "New Template",
				Fields: []core.TemplateField{
					core.TemplateField{Type: "unique", Name: "Field"},
					core.TemplateField{Type: "text", Name: "Field"},
				},
			}
			_, err := dataSource.AddTemplate(template)
			if err == nil {
				t.Error(ErrNoError)
			}
		})
		t.Run("Add template without fields", func(t *testing.T) {
			defer resetDatabase()
			template := core.Template{
				Name: "New Template",
			}
			result, err := dataSource.AddTemplate(template)
			result.ID = ""
			if err != nil {
				t.Fatal(err)
			}
			data, err := ToJSON(result)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotT(t, data)
		})
		t.Run("Add template with name of existing template", func(t *testing.T) {
			defer resetDatabase()
			template := core.Template{
				Name: templates[0].Name,
			}
			_, err := dataSource.AddTemplate(template)
			if err == nil {
				t.Error(ErrNoError)
			}
		})
		t.Run("Add field to template", func(t *testing.T) {
			defer resetDatabase()
			field := core.TemplateField{Type: "text", Name: "New Field"}
			err := dataSource.AddTemplateField(templates[0].ID, field)
			if err != nil {
				t.Error(err)
			}
			template, err := dataSource.GetTemplate(templates[0].ID)
			data, err := ToJSON(template)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotT(t, data)
		})
		t.Run("Add field to template with name of existing field", func(t *testing.T) {
			defer resetDatabase()
			field := core.TemplateField{Type: "text", Name: "Field 1"}
			err := dataSource.AddTemplateField(templates[0].ID, field)
			if err == nil {
				t.Error(ErrNoError)
			}
		})
		t.Run("Update template", func(t *testing.T) {
			err := dataSource.UpdateTemplate(templates[0].ID, core.TemplateInput{
				Name: "Updated template name",
			})
			if err != nil {
				t.Error(err)
			}
			template, err := dataSource.GetTemplate(templates[0].ID)
			if err != nil {
				t.Error(err)
			}
			data, err := ToJSON(template)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotT(t, data)
		})
		t.Run("Update template with name of existing template", func(t *testing.T) {
			defer resetDatabase()
			err := dataSource.UpdateTemplate(
				templates[0].ID,
				core.TemplateInput{
					Name: templates[1].Name,
				},
			)
			if err == nil {
				t.Error(ErrNoError)
			}
		})
		t.Run("Remove template", func(t *testing.T) {
			defer resetDatabase()
			err := dataSource.RemoveTemplate(templates[0].ID)
			if err != nil {
				t.Error(err)
			}
			templates, err := dataSource.GetTemplateList()
			if err != nil {
				t.Error(err)
			}
			data, err := ToJSON(templates)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotT(t, data)
		})
		t.Run("Remove template that does not exist", func(t *testing.T) {
			err := dataSource.RemoveTemplate("000000000099")
			if err == nil {
				t.Error(ErrNoError)
			}
		})
		t.Run("Remove field from template", func(t *testing.T) {
			defer resetDatabase()
			err := dataSource.RemoveTemplateField(templates[0].ID, "Field 1")
			if err != nil {
				t.Error(err)
			}
			template, err := dataSource.GetTemplate(templates[0].ID)
			data, err := ToJSON(template)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotT(t, data)
		})
		t.Run("Remove template field that does not exist", func(t *testing.T) {
			defer resetDatabase()
			err := dataSource.RemoveTemplateField(templates[0].ID, "Field 3")
			if err == nil {
				t.Error(ErrNoError)
			}
		})
	})
	t.Run("Test getters", func(t *testing.T) {
		t.Run("Get template", func(t *testing.T) {
			template, err := dataSource.GetTemplate(templates[0].ID)
			if err != nil {
				t.Error(err)
			}
			data, err := ToJSON(template)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotT(t, data)
		})
		t.Run("Get template list", func(t *testing.T) {
			templates, err := dataSource.GetTemplateList()
			if err != nil {
				t.Error(err)
			}
			data, err := ToJSON(templates)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotT(t, data)
		})
	})
}
