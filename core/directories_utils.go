package core

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
