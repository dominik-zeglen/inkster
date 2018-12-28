package api

import (
	"errors"
	"fmt"
)

func errNoEmpty(field string) error {
	return fmt.Errorf("Field `%s` cannot be empty", field)
}

func ErrNoPaginationLimits() error {
	return fmt.Errorf("Resolver needs \"first\" or \"last\" argument provided")
}

var errNoPermissions = errors.New("No permissions")
