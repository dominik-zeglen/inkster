package mongodb

import (
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/dominik-zeglen/ecoknow/core"
)

func init() {
	resetDatabase()
}

func TestAddPage(t *testing.T) {
	defer resetDatabase()
	page := core.Page{
		Name:     "New Page",
		Slug:     "new-page",
		ParentID: containers[0].ID,
		Fields: []core.PageField{
			core.PageField{
				Type:  "unique",
				Name:  "Field 1",
				Value: "1",
			},
			core.PageField{
				Type:  "text",
				Name:  "Field 2",
				Value: "Some text",
			},
		},
	}
	result, err := dataSource.AddPage(page)
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

func TestAddPageWithoutParentID(t *testing.T) {
	defer resetDatabase()
	page := core.Page{
		Name: "New Page",
		Slug: "new-page",
		Fields: []core.PageField{
			core.PageField{
				Type:  "unique",
				Name:  "Field 1",
				Value: "1",
			},
			core.PageField{
				Type:  "text",
				Name:  "Field 2",
				Value: "Some text",
			},
		},
	}
	_, err := dataSource.AddPage(page)
	if err == nil {
		t.Error(ErrNoError)
	}
}

func TestAddPageWithExistingSlug(t *testing.T) {
	defer resetDatabase()
	page := core.Page{
		Name: "New Page",
		Slug: pages[0].Slug,
		Fields: []core.PageField{
			core.PageField{
				Type:  "unique",
				Name:  "Field 1",
				Value: "1",
			},
			core.PageField{
				Type:  "text",
				Name:  "Field 2",
				Value: "Some text",
			},
		},
	}
	_, err := dataSource.AddPage(page)
	if err == nil {
		t.Error(ErrNoError)
	}
}

func TestAddFieldToPage(t *testing.T) {
	defer resetDatabase()
	field := core.PageField{
		Type:  "text",
		Name:  "New Field",
		Value: "New Value",
	}
	err := dataSource.AddPageField(pages[0].ID, field)
	if err != nil {
		t.Error(err)
	}
	page, err := dataSource.GetPage(pages[0].ID)
	data, err := ToJSON(page)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, data)
}

func TestAddFieldToPageWithNonExistingType(t *testing.T) {
	defer resetDatabase()
	field := core.PageField{
		Type:  "idontexist",
		Name:  "New Field",
		Value: "New Value",
	}
	err := dataSource.AddPageField(pages[0].ID, field)
	if err == nil {
		t.Error(ErrNoError)
	}
}

func TestAddExistingFieldToPage(t *testing.T) {
	defer resetDatabase()
	field := core.PageField{Type: "text", Name: "Field 1"}
	err := dataSource.AddPageField(pages[0].ID, field)
	if err == nil {
		t.Error(ErrNoError)
	}
}

func TestAddPageFromTemplate(t *testing.T) {
	defer resetDatabase()
	result, err := dataSource.AddPageFromTemplate(
		"New page",
		containers[0].ID,
		templates[0],
	)
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

func TestGetPage(t *testing.T) {
	result, err := dataSource.GetPage(pages[0].ID)
	if err != nil {
		t.Fatal(err)
	}
	data, err := ToJSON(result)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, data)
}

func TestGetPageBySlug(t *testing.T) {
	result, err := dataSource.GetPageBySlug(pages[0].Slug)
	if err != nil {
		t.Fatal(err)
	}
	data, err := ToJSON(result)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, data)
}

func TestGetPagesFromContainer(t *testing.T) {
	result, err := dataSource.GetPagesFromContainer(containers[0].ID)
	if err != nil {
		t.Fatal(err)
	}
	data, err := ToJSON(result)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, data)
}

func TestUpdatePage(t *testing.T) {
	defer resetDatabase()
	err := dataSource.UpdatePage(pages[0].ID, core.UpdatePageArguments{
		Name: "Updated page name",
	})
	if err != nil {
		t.Error(err)
	}
	page, err := dataSource.GetPage(pages[0].ID)
	if err != nil {
		t.Error(err)
	}
	data, err := ToJSON(page)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, data)
}

func TestUpdatePageWithExistingSlug(t *testing.T) {
	defer resetDatabase()
	err := dataSource.UpdatePage(pages[0].ID, core.UpdatePageArguments{
		Slug: pages[1].Slug,
	})
	if err == nil {
		t.Error(ErrNoError)
	}
}

func TestUpdatePageField(t *testing.T) {
	err := dataSource.UpdatePageField(pages[0].ID, "Field 1", "99")
	if err != nil {
		t.Error(err)
	}
	page, err := dataSource.GetPage(pages[0].ID)
	if err != nil {
		t.Error(err)
	}
	data, err := ToJSON(page)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, data)
}

func TestUpdateNonExistingPageField(t *testing.T) {
	err := dataSource.UpdatePageField(pages[0].ID, "Field 3", "99")
	if err == nil {
		t.Error(ErrNoError)
	}
}

func TestRemovePage(t *testing.T) {
	defer resetDatabase()
	err := dataSource.RemovePage(pages[0].ID)
	if err != nil {
		t.Error(err)
	}
	pages, err := dataSource.GetPagesFromContainer(containers[0].ID)
	if err != nil {
		t.Error(err)
	}
	data, err := ToJSON(pages)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, data)
}

func TestRemoveFieldFromPage(t *testing.T) {
	defer resetDatabase()
	err := dataSource.RemovePageField(pages[0].ID, "Field 1")
	if err != nil {
		t.Error(err)
	}
	page, err := dataSource.GetPage(pages[0].ID)
	data, err := ToJSON(page)
	if err != nil {
		t.Error(err)
	}
	cupaloy.SnapshotT(t, data)
}

func TestRemoveNonExistentFieldFromPage(t *testing.T) {
	defer resetDatabase()
	err := dataSource.RemovePageField(pages[0].ID, "Field 3")
	if err == nil {
		t.Error(ErrNoError)
	}
}
