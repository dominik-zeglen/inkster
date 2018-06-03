package testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/dominik-zeglen/ecoknow/core"
	"github.com/dominik-zeglen/ecoknow/mongodb"
)

var dataSources = []core.Adapter{
	mongodb.Adapter{
		ConnectionURI: os.Getenv("FOXXY_DB_URI"),
		DBName:        os.Getenv("FOXXY_DB_NAME") + "_test",
	},
	// mock.Adapter{},
}
var dataSource = dataSources[0]

var ErrNoError = fmt.Errorf("Did not return error")

// ToJSON is handy snippet for pretty-formatting json snapshots
func ToJSON(object interface{}) (string, error) {
	data, err := json.Marshal(&object)
	if err != nil {
		return "", err
	}
	var out bytes.Buffer
	json.Indent(&out, data, "", "    ")
	return out.String(), nil
}

var containers = []core.Container{
	core.Container{ID: "000000000001", Name: "Container 1"},
	core.Container{ID: "000000000002", Name: "Container 2"},
	core.Container{ID: "000000000003", Name: "Container 3"},
	core.Container{ID: "000000000004", Name: "Container 1.1", ParentID: "000000000001"},
}
var templates = []core.Template{
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
var pages = []core.Page{
	core.Page{
		ID:       "000000000001",
		Name:     "Page 1",
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
	},
	core.Page{
		ID:       "000000000002",
		Name:     "Page 2",
		ParentID: containers[0].ID,
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
		ParentID: containers[1].ID,
		Fields: []core.PageField{
			core.PageField{
				Type:  "text",
				Name:  "Field 5",
				Value: "Some textual text",
			},
		},
	},
}

func resetDatabase() {
	dataSource.ResetMockDatabase(containers, templates, pages)
}

func TestMain(t *testing.T) {
	for index := range dataSources {
		dataSource = dataSources[index]
		t.Run("Test containers", testContainers)
		t.Run("Test templates", testTemplates)
		t.Run("Test pages", testPages)
	}
}
