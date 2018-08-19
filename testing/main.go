package testing

import (
	"github.com/dominik-zeglen/inkster/core"
)

// Directories are part of testing data
var Directories = CreateDirectories()

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
var Pages = CreatePages()
