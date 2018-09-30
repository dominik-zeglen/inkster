package api

import (
	"errors"
	"fmt"
)

func errNoEmpty(field string) error {
	return fmt.Errorf("Field `%s` cannot be empty", field)
}

var errNoPermissions = errors.New("No permissions")
