package core

import "testing"

const EMAIL = "user@example.com"
const PASSWD = "password"

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
