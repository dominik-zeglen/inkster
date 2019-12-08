package core

import (
	"testing"
)

func TestPageValidation(t *testing.T) {
	defaultPage := Page{
		AuthorID: 1,
		Fields: []PageField{
			PageField{
				Slug:  "field1",
				Type:  "text",
				Value: "Example value",
			},
			PageField{
				Slug:  "field2",
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
				Name:        "",
				Slug:        defaultPage.Slug,
				ParentID:    defaultPage.ParentID,
				IsPublished: defaultPage.IsPublished,
				Fields:      defaultPage.Fields,
			},
			[]ValidationError{
				ValidationError{
					Code:  ErrRequired,
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
						Slug:  "",
						Type:  "invalid",
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
				Fields:      append(defaultPage.Fields, defaultPage.Fields[0]),
			},
			[]ValidationError{
				ValidationError{
					Code:  ErrNotUnique,
					Field: "Fields",
				},
			},
		},
		{
			Page{
				AuthorID:    99,
				Name:        defaultPage.Name,
				Slug:        defaultPage.Slug,
				ParentID:    defaultPage.ParentID,
				IsPublished: defaultPage.IsPublished,
				Fields:      defaultPage.Fields,
			},
			[]ValidationError{
				ValidationError{
					Code:  ErrDoesNotExist,
					Field: "AuthorId",
				},
			},
		},
		{
			Page{
				AuthorID:    defaultPage.AuthorID,
				Name:        defaultPage.Name,
				Slug:        defaultPage.Slug,
				ParentID:    99,
				IsPublished: defaultPage.IsPublished,
				Fields:      defaultPage.Fields,
			},
			[]ValidationError{
				ValidationError{
					Code:  ErrDoesNotExist,
					Field: "ParentId",
				},
			},
		},
	}

	for index, testData := range testSuites {
		result, err := testData.page.Validate(dataSource)
		if err != nil {
			t.Fatal(err)
		}

		testValidation(testData.expected, result, index, t)
	}
}
