package core

func CreateDirectory(
	directory Directory,
	dataSource AbstractDataContext,
) (*Directory, []ValidationError, error) {
	validationErrors, err := directory.Validate(dataSource)

	if err != nil || len(validationErrors) > 0 {
		return nil, validationErrors, err
	}

	directory.CreatedAt = dataSource.GetCurrentTime()
	directory.UpdatedAt = dataSource.GetCurrentTime()

	_, err = dataSource.
		DB().
		Model(&directory).
		Insert()

	return &directory, validationErrors, err
}

func UpdateDirectory(
	directory Directory,
	dataSource AbstractDataContext,
) (*Directory, []ValidationError, error) {
	validationErrors, err := directory.Validate(dataSource)

	if err != nil || len(validationErrors) > 0 {
		return nil, validationErrors, err
	}

	directory.UpdatedAt = dataSource.GetCurrentTime()

	_, err = dataSource.
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
