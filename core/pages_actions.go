package core

func CreatePage(
	page Page,
	dataSource AbstractDataContext,
) (*Page, []ValidationError, error) {
	validationErrors, err := page.Validate(dataSource)

	if len(validationErrors) > 0 {
		return nil, validationErrors, nil
	}

	page.CreatedAt = dataSource.GetCurrentTime()
	page.UpdatedAt = dataSource.GetCurrentTime()

	_, err = dataSource.
		DB().
		Model(&page).
		Insert()

	return &page, validationErrors, err
}

func UpdatePage(
	page Page,
	dataSource AbstractDataContext,
) (*Page, []ValidationError, error) {
	validationErrors, err := page.Validate(dataSource)

	if err != nil || len(validationErrors) > 0 {
		return nil, validationErrors, err
	}

	page.UpdatedAt = dataSource.GetCurrentTime()

	_, err = dataSource.
		DB().
		Model(&page).
		WherePK().
		Update()

	return &page, validationErrors, err
}

func RemovePage(id int, dataSource AbstractDataContext) error {
	_, err := dataSource.
		DB().
		Exec("DELETE FROM pages WHERE id = ?", id)

	return err
}
