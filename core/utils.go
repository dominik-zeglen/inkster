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

func checkIfDirectoryExist(
	id int,
	dataSource AbstractDataContext,
) (bool, error) {
	directory := Directory{}
	directory.ID = id
	exists, err := dataSource.
		DB().
		Model(&directory).
		WherePK().
		Exists()

	if err != nil {
		return false, err
	}

	return exists, nil
}

func checkIfPageExist(
	id int,
	dataSource AbstractDataContext,
) (bool, error) {
	page := Page{}
	page.ID = id
	exists, err := dataSource.
		DB().
		Model(&page).
		WherePK().
		Exists()

	if err != nil {
		return false, err
	}

	return exists, nil
}

func checkIfUserExist(
	id int,
	dataSource AbstractDataContext,
) (bool, error) {
	user := User{}
	user.ID = id
	exists, err := dataSource.
		DB().
		Model(&user).
		WherePK().
		Exists()

	if err != nil {
		return false, err
	}

	return exists, nil
}
