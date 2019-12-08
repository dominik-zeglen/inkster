package core

import (
	"testing"
)

func TestDirectoryValidation(t *testing.T) {
	parentID := 2
	fakeParentID := 99
	defaultDirectory := Directory{
		Name:        "directory name",
		ParentID:    &parentID,
		IsPublished: true,
	}
	defaultDirectory.ID = 5

	testSuites := []struct {
		directory Directory
		expected  []ValidationError
	}{
		{
			defaultDirectory,
			[]ValidationError{},
		},
		{
			Directory{
				Name:        "",
				ParentID:    defaultDirectory.ParentID,
				IsPublished: defaultDirectory.IsPublished,
			},
			[]ValidationError{
				ValidationError{
					Code:  ErrRequired,
					Field: "Name",
				},
			},
		},
		{
			Directory{
				Name:        defaultDirectory.Name,
				ParentID:    &fakeParentID,
				IsPublished: defaultDirectory.IsPublished,
			},
			[]ValidationError{
				ValidationError{
					Code:  ErrDoesNotExist,
					Field: "ParentID",
				},
			},
		},
	}

	for index, testData := range testSuites {
		result, err := testData.directory.Validate(dataSource)
		if err != nil {
			t.Fatal(err)
		}

		testValidation(testData.expected, result, index, t)
	}
}
