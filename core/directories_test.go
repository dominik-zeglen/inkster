package core

import (
	"testing"
)

const (
	DIR_NAME   = "directory name"
	DIR_PARENT = 2
)

func TestDirectoryValidation(t *testing.T) {
	testSuites := []struct {
		directory Directory
		expected  []ValidationError
	}{
		{
			Directory{
				Name:        DIR_NAME,
				ParentID:    DIR_PARENT,
				IsPublished: true,
			},
			[]ValidationError{},
		},
		{
			Directory{
				Name:        "a",
				ParentID:    DIR_PARENT,
				IsPublished: true,
			},
			[]ValidationError{
				ValidationError{
					Code:  ErrMinLength,
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

func TestDirectoryInputValidation(t *testing.T) {
	dirName := DIR_NAME
	newDirName := "a"

	testSuites := []struct {
		directoryInput DirectoryInput
		expected       []ValidationError
	}{
		{
			DirectoryInput{
				Name: &dirName,
			},
			[]ValidationError{},
		},
		{
			DirectoryInput{
				Name: &newDirName,
			},
			[]ValidationError{
				ValidationError{
					Code:  ErrMinLength,
					Field: "Name",
				},
			},
		},
	}

	for index, testData := range testSuites {
		result := ValidateModel(testData.directoryInput)
		testValidation(testData.expected, result, index, t)
	}
}
