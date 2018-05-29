package mongodb

import (
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/dominik-zeglen/ecoknow/core"
)

func init() {
	resetDatabase()
}

func TestAddTemplate(t *testing.T) {
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
}

func TestAddTemplateWithoutName(t *testing.T) {
	defer resetDatabase()
	template := core.Template{}
	_, err := dataSource.AddTemplate(template)
	if err == nil {
		t.Error(ErrNoError)
	}
}

func TestAddTemplateWithDuplicatedFields(t *testing.T) {
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
}

func TestAddTemplateWithoutFields(t *testing.T) {
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
}

func TestAddTemplateWithExistingName(t *testing.T) {
	defer resetDatabase()
	template := core.Template{
		Name: templates[0].Name,
	}
	_, err := dataSource.AddTemplate(template)
	if err == nil {
		t.Error(ErrNoError)
	}
}

func TestAddFieldToTemplate(t *testing.T) {
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
}

func TestAddExistingFieldToTemplate(t *testing.T) {
	defer resetDatabase()
	field := core.TemplateField{Type: "text", Name: "Field 1"}
	err := dataSource.AddTemplateField(templates[0].ID, field)
	if err == nil {
		t.Error(ErrNoError)
	}
}

func TestGetTemplate(t *testing.T) {
	template, err := dataSource.GetTemplate(templates[0].ID)
	if err != nil {
		t.Error(err)
	}
	data, err := ToJSON(template)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, data)
}

func TestGetTemplateList(t *testing.T) {
	templates, err := dataSource.GetTemplateList()
	if err != nil {
		t.Error(err)
	}
	data, err := ToJSON(templates)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, data)
}

func TestUpdateTemplate(t *testing.T) {
	err := dataSource.UpdateTemplate(templates[0].ID, core.UpdateTemplateArguments{
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
}

func TestUpdateTemplateWithExistingName(t *testing.T) {
	defer resetDatabase()
	err := dataSource.UpdateTemplate(
		templates[0].ID,
		core.UpdateTemplateArguments{
			Name: templates[1].Name,
		},
	)
	if err == nil {
		t.Error(ErrNoError)
	}
}

func TestRemoveTemplate(t *testing.T) {
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
}

func TestRemoveFieldFromTemplate(t *testing.T) {
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
}

func TestRemoveNonExistentFieldFromTemplate(t *testing.T) {
	defer resetDatabase()
	err := dataSource.RemoveTemplateField(templates[0].ID, "Field 3")
	if err == nil {
		t.Error(ErrNoError)
	}
}
