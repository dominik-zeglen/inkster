package core

import "fmt"

// ErrNoEmpty informs about missing property
func ErrNoEmpty(key string) error {
	return fmt.Errorf("Property cannot be omitted: %s", key)
}

// ErrNoField informs about non-existing field
func ErrNoField(field string) error {
	return fmt.Errorf("Field %s does not exist", field)
}

// ErrNoFieldType informs about non-existing field type
func ErrNoFieldType(fieldType string) error {
	return fmt.Errorf("Field type %s does not exist", fieldType)
}

// ErrFieldExists informs about existing field
func ErrFieldExists(field string) error {
	return fmt.Errorf("Field %s already exists", field)
}
