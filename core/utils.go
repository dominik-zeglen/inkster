package core

import (
	"github.com/go-pg/pg"
	"gopkg.in/go-playground/validator.v9"
)

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

func getUserByEmail(
	email string,
	dataSource AbstractDataContext,
) (*int, error) {
	user := User{
		Email: email,
	}
	err := dataSource.
		DB().
		Model(&user).
		Where("email = ?", user.Email).
		First()

	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &user.ID, nil
}
