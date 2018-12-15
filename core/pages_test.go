package core

import (
	"testing"
)

func TestPageValidation(t *testing.T) {
	defaultPage := Page{
		AuthorID: 1,
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
		ParentID:    1,
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
				AuthorID:    defaultPage.AuthorID,
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
				AuthorID:    defaultPage.AuthorID,
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
				AuthorID:    defaultPage.AuthorID,
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
