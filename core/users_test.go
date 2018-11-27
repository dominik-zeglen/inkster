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
		err := user.CreatePassword(PASSWD)
		if err != nil {
			t.Fatal(err)
		}

		result := user.AuthPassword(testData.password)
		if result != testData.expected {
			t.Fatalf("Expected: %t, got: %t", testData.expected, result)
		}
	}
}

func TestUserValidation(t *testing.T) {
	testSuites := []struct {
		user     User
		expected []ValidationError
	}{
		{
			User{
				Active:   true,
				Email:    EMAIL,
				Password: []byte(PASSWD_HASH),
				Salt:     []byte(PASSWD_SALT),
			},
			[]ValidationError{},
		},
		{
			User{
				Active:   true,
				Email:    "notanemail",
				Password: []byte(PASSWD_HASH),
				Salt:     []byte(PASSWD_SALT),
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
		testValidation(testData.expected, result, index, t)
	}
}
