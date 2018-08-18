package testing

import (
	"github.com/dominik-zeglen/inkster/core"
	"github.com/globalsign/mgo/bson"
)

func createPage(page core.Page, id string, createdAt string, updatedAt string) core.Page {
	page.ID = bson.ObjectId(id)
	page.CreatedAt = &createdAt
	page.UpdatedAt = &updatedAt

	return page
}

func CreatePages() []core.Page {
	pages := []core.Page{
		core.Page{
			Name:     "Page 1",
			Slug:     "page-1",
			ParentID: Directories[0].ID,
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
		},
		core.Page{
			Name:     "Page 2",
			Slug:     "page-2",
			ParentID: Directories[0].ID,
			Fields: []core.PageField{
				core.PageField{
					Type:  "unique",
					Name:  "Field 3",
					Value: "2",
				},
				core.PageField{
					Type:  "file",
					Name:  "Field 4",
					Value: "example.com/file",
				},
			},
		},
		core.Page{
			Name:     "Page 3",
			Slug:     "page-3",
			ParentID: Directories[1].ID,
			Fields: []core.PageField{
				core.PageField{
					Type:  "text",
					Name:  "Field 5",
					Value: "Some textual text",
				},
			},
		},
	}

	pages[0] = createPage(
		pages[0],
		"000000000001",
		"2007-07-07T10:00:00.000Z",
		"2007-07-07T10:00:00.000Z",
	)
	pages[1] = createPage(
		pages[1],
		"000000000002",
		"2007-07-07T11:00:00.000Z",
		"2007-07-07T11:00:00.000Z",
	)
	pages[2] = createPage(
		pages[2],
		"000000000003",
		"2007-07-07T12:00:00.000Z",
		"2007-07-07T12:00:00.000Z",
	)

	return pages
}
