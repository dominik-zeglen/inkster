package graphql

import "fmt"

func errNoEmpty(field string) error {
	return fmt.Errorf("Field `%s` cannot be empty", field)
}
