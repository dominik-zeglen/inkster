package core

import (
	"github.com/go-pg/pg"
)

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

	err = dataSource.DB().RunInTransaction(func(tx *pg.Tx) error {
		_, err = tx.
			DB().
			Model(&page).
			Insert()

		return err
	})

	return &page, validationErrors, err
}
