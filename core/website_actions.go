package core

func UpdateWebsite(
	website Website,
	dataSource AbstractDataContext,
) (*Website, []ValidationError, error) {
	website.ID = WEBSITE_DB_ID
	validationErrors := website.Validate()

	if len(validationErrors) > 0 {
		return nil, validationErrors, nil
	}

	_, err := dataSource.
		DB().
		Model(&website).
		WherePK().
		UpdateNotNull()

	return &website, validationErrors, err
}
