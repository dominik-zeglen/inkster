package core

import "gopkg.in/go-playground/validator.v9"

func ValidateModel(model interface{}) []ValidationError {
	validationErrors := []ValidationError{}
	errors := validate.Struct(model)
	if errors != nil {
		for _, err := range errors.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, ToValidationError(err))
		}
	}
	return validationErrors
}
