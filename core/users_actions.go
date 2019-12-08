package core

type UserCreateInput struct {
	Email string
}

func CreateUser(
	input UserCreateInput,
	dataSource AbstractDataContext,
) (*User, []ValidationError, error) {
	user := User{
		Active: false,
		Email:  input.Email,
	}
	user.CreateRandomPassword()

	validationErrors := user.Validate()

	if len(validationErrors) > 0 {
		return nil, validationErrors, nil
	}

	user.CreatedAt = dataSource.GetCurrentTime()
	user.UpdatedAt = dataSource.GetCurrentTime()

	_, err := dataSource.
		DB().
		Model(&user).
		Insert()

	return &user, validationErrors, err
}

func UpdateUser(
	user User,
	dataSource AbstractDataContext,
) (*User, []ValidationError, error) {
	validationErrors := user.Validate()

	if user.Email != "" {
		existingID, err := getUserByEmail(user.Email, dataSource)
		exists := existingID != nil

		if err != nil {
			return nil, validationErrors, err
		} else if exists && *existingID != user.ID {
			validationErrors = append(validationErrors, ValidationError{
				Code:  ErrNotUnique,
				Field: "email",
				Param: &user.Email,
			})
		}
	}

	if len(validationErrors) > 0 {
		return nil, validationErrors, nil
	}

	user.UpdatedAt = dataSource.GetCurrentTime()

	_, err := dataSource.
		DB().
		Model(&user).
		WherePK().
		Update()

	return &user, validationErrors, err
}

func RemoveUser(id int, dataSource AbstractDataContext) error {
	_, err := dataSource.
		DB().
		Exec("DELETE FROM users WHERE id = ?", id)

	return err
}
