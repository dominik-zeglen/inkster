package core

import (
	"testing"
)

func TestDirectoryValidation(t *testing.T) {
	dirName := "directory name"
	dirParent := 2

	testSuites := []struct {
		directory Directory
		expected  []ValidationError
	}{
		{
			Directory{
				Name:        dirName,
				ParentID:    &dirParent,
				IsPublished: true,
			},
			[]ValidationError{},
		},
		{
			Directory{
				Name:        "",
				ParentID:    &dirParent,
				IsPublished: true,
			},
			[]ValidationError{
				ValidationError{
					Code:  ErrRequired,
					Field: "Name",
				},
			},
		},
	}

	for index, testData := range testSuites {
		result := testData.directory.Validate()
		testValidation(testData.expected, result, index, t)
	}
}
