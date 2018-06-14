package testing

import (
	"github.com/dominik-zeglen/ecoknow/core"
)

// Containers are part of testing data
var Containers = []core.Container{
	core.Container{ID: "000000000001", Name: "Container 1"},
	core.Container{ID: "000000000002", Name: "Container 2"},
	core.Container{ID: "000000000003", Name: "Container 3"},
	core.Container{ID: "000000000004", Name: "Container 1.1", ParentID: "000000000001"},
}

// Templates are part of testing data
var Templates = []core.Template{
	core.Template{
		ID:   "000000000001",
		Name: "Template 1",
		Fields: []core.TemplateField{
			core.TemplateField{Type: "text", Name: "Field 1"},
			core.TemplateField{Type: "image", Name: "Field 2"},
		},
	},
	core.Template{
		ID:   "000000000002",
		Name: "Template 2",
		Fields: []core.TemplateField{
			core.TemplateField{Type: "unique", Name: "Field 3"},
			core.TemplateField{Type: "text", Name: "Field 4"},
			core.TemplateField{Type: "file", Name: "Field 5"},
		},
	},
}

// Pages are part of testing data
var Pages = []core.Page{
	core.Page{
		ID:       "000000000001",
		Name:     "Page 1",
		ParentID: Containers[0].ID,
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
		ID:       "000000000002",
		Name:     "Page 2",
		ParentID: Containers[0].ID,
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
		ID:       "000000000003",
		Name:     "Page 3",
		ParentID: Containers[1].ID,
		Fields: []core.PageField{
			core.PageField{
				Type:  "text",
				Name:  "Field 5",
				Value: "Some textual text",
			},
		},
	},
}
