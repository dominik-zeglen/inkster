package core

import (
	"testing"
)

const EMAIL = "user@example.com"
const PASSWD = "password"

const PASSWD_HASH = "randomhash"
const PASSWD_SALT = "randomsalt"

func TestAuth(t *testing.T) {
	testSuites := []struct {
		password string
		active   bool
		expected bool
	}{
		{PASSWD, true, true},
		{PASSWD, false, false},
		{PASSWD + ".", true, false},
	}

	for _, testData := range testSuites {
		user := User{
			Active: testData.active,
			Email:  EMAIL,
		}
		user.CreatePassword(PASSWD)

		result := user.AuthPassword(testData.password)
		if result != testData.expected {
			t.Fatalf("Expected: %t, got: %t", testData.expected, result)
		}
	}
}

func TestValidation(t *testing.T) {
	testSuites := []struct {
		user     User
		expected []ValidationError
	}{
		{
			User{
				Active:   true,
				Email:    EMAIL,
				Password: PASSWD_HASH,
				Salt:     PASSWD_SALT,
			},
			[]ValidationError{},
		},
		{
			User{
				Active:   true,
				Email:    "notanemail",
				Password: PASSWD_HASH,
				Salt:     PASSWD_SALT,
			},
			[]ValidationError{
				ValidationError{
					Code:  ErrTypeMismatch,
					Field: "Email",
				},
			},
		},
	}

	for index, testData := range testSuites {
		result := testData.user.Validate()
		if len(testData.expected) != len(result) {
			t.Fatalf(
				"Test %d failed: expected %d errors, got %d",
				index,
				len(testData.expected),
				len(result),
			)
		}
		for errIndex, err := range result {
			if testData.expected[errIndex].Code != err.Code {
				t.Fatalf(
					"Test %d failed: expected error %d, got %d",
					index,
					testData.expected[errIndex].Code,
					err.Code,
				)
			}
			if testData.expected[errIndex].Field != err.Field {
				t.Fatalf(
					"Test %d failed: expected error in field %s, got %s",
					index,
					testData.expected[errIndex].Field,
					err.Field,
				)
			}
		}
	}
}
