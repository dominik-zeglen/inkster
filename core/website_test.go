package core

import (
	"testing"
)

const (
	WEBSITE_NAME   = "Test website"
	WEBSITE_DOMAIN = "http://testwebsite.example"
)

func TestWebsiteValidation(t *testing.T) {
	testSuites := []struct {
		website  Website
		expected []ValidationError
	}{
		{
			Website{
				Domain: WEBSITE_DOMAIN,
				Name:   WEBSITE_NAME,
			},
			[]ValidationError{},
		},
		{
			Website{
				Domain: "a",
				Name:   WEBSITE_NAME,
			},
			[]ValidationError{
				ValidationError{
					Code:  ErrNotUrl,
					Field: "Domain",
				},
			},
		},
		{
			Website{
				Domain: WEBSITE_DOMAIN,
				Name:   "a",
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
		result := testData.website.Validate()
		testValidation(testData.expected, result, index, t)
	}
}
