package core

func CreateDirectory(
	directory Directory,
	dataSource AbstractDataContext,
) (*Directory, []ValidationError, error) {
	validationErrors := directory.Validate()

	if directory.ParentID != nil {
		exists, err := checkIfDirectoryExist(*directory.ParentID, dataSource)
		if err != nil {
			return nil, validationErrors, err
		}

		if !exists {
			validationErrors = append(validationErrors, ValidationError{
				Code:  ErrDoesNotExist,
				Field: "ParentID",
			})
		}
	}

	if len(validationErrors) > 0 {
		return nil, validationErrors, nil
	}

	directory.CreatedAt = dataSource.GetCurrentTime()
	directory.UpdatedAt = dataSource.GetCurrentTime()

	_, err := dataSource.
		DB().
		Model(&directory).
		Insert()

	return &directory, validationErrors, err
}

func UpdateDirectory(
	directory Directory,
	dataSource AbstractDataContext,
) (*Directory, []ValidationError, error) {
	validationErrors := directory.Validate()

	if directory.ParentID != nil {
		parent := Directory{}
		parent.ID = *directory.ParentID
		exists, err := dataSource.
			DB().
			Model(&parent).
			WherePK().
			Exists()

		if err != nil {
			return nil, validationErrors, err
		}

		if !exists {
			validationErrors = append(validationErrors, ValidationError{
				Code:  ErrDoesNotExist,
				Field: "ParentID",
			})
		}
	}

	if len(validationErrors) > 0 {
		return nil, validationErrors, nil
	}

	directory.UpdatedAt = dataSource.GetCurrentTime()

	_, err := dataSource.
		DB().
		Model(&directory).
		WherePK().
		Update()

	return &directory, validationErrors, err
}

func RemoveDirectory(id int, dataSource AbstractDataContext) error {
	_, err := dataSource.
		DB().
		Exec("DELETE FROM directories WHERE id = ?", id)

	return err
}
