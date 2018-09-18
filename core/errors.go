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

// ErrTemplateExists informs about existing template
func ErrTemplateExists(template string) error {
	return fmt.Errorf("Template %s already exists", template)
}

// ErrPageExists informs about existing page
func ErrPageExists(page string) error {
	return fmt.Errorf("Page %s already exists", page)
}

// ErrUserExists informs about existing user
func ErrUserExists(user string) error {
	return fmt.Errorf("User %s already exists", user)
}
