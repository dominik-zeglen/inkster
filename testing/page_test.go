package testing

import (
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/dominik-zeglen/inkster/core"
)

func stripIds(fields []core.PageField) []core.PageField {
	dummyID := 0
	cleanedFields := make([]core.PageField, len(fields))
	copy(cleanedFields, fields)

	for fieldIndex, _ := range cleanedFields {
		cleanedFields[fieldIndex].ID = dummyID
	}

	return cleanedFields
}

func testPages(t *testing.T) {
	dummyID := 0
	t.Run("Test setters", func(t *testing.T) {
		t.Run("Add page", func(t *testing.T) {
			defer resetDatabase()
			page := core.Page{
				Name:     "New Page",
				Slug:     "new-page",
				ParentID: Directories[0].ID,
				Fields: []core.PageField{
					core.PageField{
						Type:  "text",
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
			result.ID = dummyID
			result.Fields = stripIds(result.Fields)
			if err != nil {
				t.Fatal(err)
			}
			data, err := ToJSON(result)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotT(t, data)
		})
		t.Run("Add page without parent", func(t *testing.T) {
			defer resetDatabase()
			page := core.Page{
				Name: "New Page",
				Slug: "new-page",
				Fields: []core.PageField{
					core.PageField{
						Type:  "text",
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
		})
		t.Run("Add page with another's page slug", func(t *testing.T) {
			defer resetDatabase()
			page := core.Page{
				Name: "New Page",
				Slug: Pages[0].Slug,
				Fields: []core.PageField{
					core.PageField{
						Type:  "text",
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
		})
		t.Run("Add page with duplicated fields", func(t *testing.T) {
			defer resetDatabase()
			page := core.Page{
				Name: "New Page",
				Slug: Pages[0].Slug,
				Fields: []core.PageField{
					core.PageField{
						Type:  "text",
						Name:  "Field 1",
						Value: "1",
					},
					core.PageField{
						Type:  "text",
						Name:  "Field 1",
						Value: "Some text",
					},
				},
			}
			_, err := dataSource.AddPage(page)
			if err == nil {
				t.Error(ErrNoError)
			}
		})
		// t.Run("Add page from a template", func(t *testing.T) {
		// 	defer resetDatabase()
		// 	pageName := "New page"
		// 	result, err := dataSource.AddPageFromTemplate(
		// 		core.PageInput{
		// 			Name:     &pageName,
		// 			ParentID: &Directories[0].ID,
		// 		},
		// 		Templates[0].ID,
		// 	)
		// 	result.ID = ""
		// 	result.Fields = stripIds(result.Fields)
		// 	if err != nil {
		// 		t.Fatal(err)
		// 	}
		// 	data, err := ToJSON(result)
		// 	if err != nil {
		// 		t.Error(err)
		// 	}
		// 	cupaloy.SnapshotT(t, data)
		// })
		t.Run("Add field to page", func(t *testing.T) {
			defer resetDatabase()
			field := core.PageField{
				Type:  "text",
				Name:  "New Field",
				Value: "New Value",
			}
			err := dataSource.AddPageField(Pages[0].ID, field)
			if err != nil {
				t.Error(err)
			}
			page, err := dataSource.GetPage(Pages[0].ID)
			if err != nil {
				t.Error(err)
			}
			page.Fields[2].ID = 0
			data, err := ToJSON(page)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotT(t, data)
		})
		t.Run("Add field of unknown type to page", func(t *testing.T) {
			defer resetDatabase()
			field := core.PageField{
				Type:  "idontexist",
				Name:  "New Field",
				Value: "New Value",
			}
			err := dataSource.AddPageField(Pages[0].ID, field)
			if err == nil {
				t.Error(ErrNoError)
			}
		})
		t.Run("Update page", func(t *testing.T) {
			defer resetDatabase()
			pageName := "Updated page name"
			isPublished := true
			err := dataSource.UpdatePage(Pages[0].ID, core.PageInput{
				Name:        &pageName,
				IsPublished: &isPublished,
			})
			if err != nil {
				t.Error(err)
			}
			page, err := dataSource.GetPage(Pages[0].ID)
			if err != nil {
				t.Error(err)
			}
			data, err := ToJSON(page)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotT(t, data)
		})
		t.Run("Update page with another's page slug", func(t *testing.T) {
			defer resetDatabase()
			err := dataSource.UpdatePage(Pages[0].ID, core.PageInput{
				Slug: &Pages[1].Slug,
			})
			if err == nil {
				t.Error(ErrNoError)
			}
		})
		t.Run("Remove page", func(t *testing.T) {
			defer resetDatabase()
			err := dataSource.RemovePage(Pages[0].ID)
			if err != nil {
				t.Error(err)
			}
			pages, err := dataSource.GetPagesFromDirectory(Directories[0].ID)
			if err != nil {
				t.Error(err)
			}
			data, err := ToJSON(pages)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotT(t, data)
		})
		t.Run("Remove page's field", func(t *testing.T) {
			defer resetDatabase()
			err := dataSource.RemovePageField(Pages[0].Fields[0].ID)
			if err != nil {
				t.Error(err)
			}
			page, err := dataSource.GetPage(Pages[0].ID)
			data, err := ToJSON(page)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotT(t, data)
		})
		t.Run("Remove page's field that does not exist", func(t *testing.T) {
			defer resetDatabase()
			err := dataSource.RemovePageField(0)
			if err == nil {
				t.Error(ErrNoError)
			}
		})
	})
	t.Run("Test getters", func(t *testing.T) {
		t.Run("Get page", func(t *testing.T) {
			result, err := dataSource.GetPage(Pages[0].ID)
			if err != nil {
				t.Fatal(err)
			}
			data, err := ToJSON(result)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotT(t, data)
		})
		t.Run("Get page by slug", func(t *testing.T) {
			result, err := dataSource.GetPageBySlug(Pages[0].Slug)
			if err != nil {
				t.Fatal(err)
			}
			data, err := ToJSON(result)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotT(t, data)
		})
		t.Run("Get pages from directory", func(t *testing.T) {
			result, err := dataSource.GetPagesFromDirectory(Directories[0].ID)
			if err != nil {
				t.Fatal(err)
			}
			data, err := ToJSON(result)
			if err != nil {
				t.Error(err)
			}
			cupaloy.SnapshotT(t, data)
		})
	})
}
