package core

import (
	"testing"
)

func TestPageValidation(t *testing.T) {
	defaultPage := Page{
		Fields: []PageField{
			PageField{
				Name:  "field1",
				Type:  "text",
				Value: "Example value",
			},
			PageField{
				Name:  "field2",
				Type:  "text",
				Value: "Example value",
			},
		},
		IsPublished: true,
		Name:        "Example page",
		Slug:        "example-page",
		ParentID:    "000000000001",
	}

	testSuites := []struct {
		page     Page
		expected []ValidationError
	}{
		{
			defaultPage,
			[]ValidationError{},
		},
		{
			Page{
				Name:        "a",
				Slug:        defaultPage.Slug,
				ParentID:    defaultPage.ParentID,
				IsPublished: defaultPage.IsPublished,
				Fields:      defaultPage.Fields,
			},
			[]ValidationError{
				ValidationError{
					Code:  ErrMinLength,
					Field: "Name",
				},
			},
		},
		{
			Page{
				Name:        defaultPage.Name,
				Slug:        "not valid",
				ParentID:    defaultPage.ParentID,
				IsPublished: defaultPage.IsPublished,
				Fields:      defaultPage.Fields,
			},
			[]ValidationError{
				ValidationError{
					Code:  ErrNotSlug,
					Field: "Slug",
				},
			},
		},
		{
			Page{
				Name:        defaultPage.Name,
				Slug:        defaultPage.Slug,
				ParentID:    defaultPage.ParentID,
				IsPublished: defaultPage.IsPublished,
				Fields: []PageField{
					{
						Type:  "invalid",
						Name:  "",
						Value: "",
					},
				},
			},
			[]ValidationError{
				ValidationError{
					Code:  ErrEqual,
					Field: "Type",
				},
				ValidationError{
					Code:  ErrRequired,
					Field: "Name",
				},
			},
		},
	}

	for index, testData := range testSuites {
		result := testData.page.Validate()
		testValidation(testData.expected, result, index, t)
	}
}
