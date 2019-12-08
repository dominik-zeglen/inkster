package core

import "testing"

func testValidation(
	expected []ValidationError,
	returned []ValidationError,
	index int,
	t *testing.T,
) {
	if len(expected) != len(returned) {
		t.Fatalf(
			"Test %d failed: expected %d errors, got %d",
			index,
			len(expected),
			len(returned),
		)
	}
	for errIndex, err := range returned {
		if expected[errIndex].Code != err.Code {
			t.Fatalf(
				"Test %d failed: expected error %d, got %d",
				index,
				expected[errIndex].Code,
				err.Code,
			)
		}
		if expected[errIndex].Field != err.Field {
			t.Fatalf(
				"Test %d failed: expected error in field %s, got %s",
				index,
				expected[errIndex].Field,
				err.Field,
			)
		}
	}
}
