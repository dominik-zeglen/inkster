package core

import (
	"errors"
	"fmt"
	"github.com/gosimple/slug"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func validateSlug(fl validator.FieldLevel) bool {
	return fl.Field().String() == slug.Make(fl.Field().String())
}

func init() {
	validate = validator.New()
	validate.RegisterValidation("slug", validateSlug)
}

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

// ErrNotValidated is thrown if validation during
// model operations was unsuccessful
var ErrNotValidated = errors.New("Invalid model")

// ErrBadCredentials informs about existing user
var ErrBadCredentials = fmt.Errorf("Bad credentials")

type ValidationErrorCode int

const (
	// Data format errors
	ErrRequired     = 100
	ErrMinLength    = 101
	ErrMaxLength    = 102
	ErrLength       = 103
	ErrTypeMismatch = 104
	ErrNotEqual     = 105
	ErrNotSlug      = 106

	// Model errors
	ErrNotUnique    = 200
	ErrDoesNotExist = 201
)

// ValidationError is thrown if object or input weren't valid
type ValidationError struct {
	Code  ValidationErrorCode
	Field string

	Param *string
	Range *struct {
		min int
		max int
	}
	Type *string
}

func (err ValidationError) Error() string {
	switch err.Code {
	case ErrRequired:
		return fmt.Sprintf("Property cannot be omitted: %s", err.Field)

	case ErrLength:
		return fmt.Sprintf(
			"Property %s must be exactly %s characters long",
			err.Field,
			*err.Param,
		)
	case ErrMinLength:
		return fmt.Sprintf(
			"Property %s must be at least %s characters long",
			err.Field,
			*err.Param,
		)
	case ErrMaxLength:
		return fmt.Sprintf(
			"Property %s exceeds %s character limit",
			err.Field,
			*err.Param,
		)

	case ErrTypeMismatch:
		return fmt.Sprintf(
			"Property %s is not of type %s",
			err.Field,
			*err.Type,
		)

	case ErrNotUnique:
		return fmt.Sprintf(
			"Object with %s %s already exists",
			err.Field,
			*err.Param,
		)

	case ErrDoesNotExist:
		return fmt.Sprintf(
			"Object with %s %s does not exist",
			err.Field,
			*err.Param,
		)

	case ErrNotEqual:
		return fmt.Sprintf(
			"Property %s cannot be equal %s",
			err.Field,
			*err.Param,
		)

	case ErrNotSlug:
		return fmt.Sprintf(
			"Property %s should contain only small characters, numbers and _, - characters",
			err.Field,
		)
	}

	return "Unknown error"
}

func ToValidationError(err validator.FieldError) ValidationError {
	validationError := ValidationError{
		Field: err.Field(),
	}

	switch err.Tag() {
	case "email":
		validationError.Code = ErrTypeMismatch
		fieldType := "email"
		validationError.Type = &fieldType
	case "len":
		validationError.Code = ErrLength
		param := err.Param()
		validationError.Param = &param
	case "min":
		validationError.Code = ErrMinLength
		param := err.Param()
		validationError.Param = &param
	case "max":
		validationError.Code = ErrMaxLength
		param := err.Param()
		validationError.Param = &param
	case "required":
		validationError.Code = ErrRequired
	case "slug":
		validationError.Code = ErrNotSlug
	case "oneof":
		validationError.Code = ErrNotEqual
	}

	return validationError
}
