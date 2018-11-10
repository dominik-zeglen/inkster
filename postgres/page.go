package postgres

import (
	"github.com/dominik-zeglen/inkster/core"
	"github.com/gosimple/slug"
)

func cleanAddPageInput(page *core.Page) {
	if page.Slug == "" {
		page.Slug = slug.Make(page.Name)
	}
}

// AddPage puts page in the database
func (adapter Adapter) AddPage(page core.Page) (core.Page, error) {
	cleanAddPageInput(&page)
	errs := page.Validate()
	if len(errs) > 0 {
		return core.Page{}, core.ErrNotValidated
	}

	page.CreatedAt = adapter.GetCurrentTime()
	page.UpdatedAt = adapter.GetCurrentTime()

	_, err := adapter.
		Session.
		Model(&page).
		Insert()

	if err != nil {
		return core.Page{}, err
	}

	for fieldIndex, _ := range page.Fields {
		page.Fields[fieldIndex].PageID = page.ID
	}

	_, err = adapter.
		Session.
		Model(&page.Fields).
		Insert()

	return page, err
}

// AddPageFromTemplate creates new page based on a chosen template
func (adapter Adapter) AddPageFromTemplate(
	page core.PageInput,
	templateID int,
) (core.Page, error) {
	template, err := adapter.GetTemplate(templateID)
	if err != nil {
		return core.Page{}, err
	}

	var fields []core.PageField
	for _, field := range template.Fields {
		fields = append(fields, core.PageField{
			Name:  field.Name,
			Type:  field.Type,
			Value: "",
		})
	}

	inputPage := core.Page{
		Name:     *page.Name,
		ParentID: *page.ParentID,
		Fields:   fields,
	}
	if page.Slug != nil {
		inputPage.Slug = *page.Slug
	}

	return adapter.AddPage(inputPage)
}

// AddPageField adds to page a new field at the end of it's field list
func (adapter Adapter) AddPageField(pageID int, field core.PageField) error {
	errs := field.Validate()
	if len(errs) > 0 {
		return core.ErrNotValidated
	}

	field.PageID = pageID
	field.CreatedAt = adapter.GetCurrentTime()
	field.UpdatedAt = adapter.GetCurrentTime()

	_, err := adapter.
		Session.
		Model(&field).
		Insert()

	if err != nil {
		return err
	}

	_, err = adapter.
		Session.
		Exec(
			"UPDATE pages SET updated_at = ? WHERE id = ?",
			adapter.GetCurrentTime(),
			pageID,
		)

	return err
}

// GetPage allows user to fetch page by ID from database
func (adapter Adapter) GetPage(id int) (core.Page, error) {
	page := core.Page{}

	err := adapter.
		Session.
		Model(&page).
		Where("id = ?", id).
		Relation("Fields").
		Select()

	return page, err
}

// GetPageBySlug allows user to fetch page by slug from database
func (adapter Adapter) GetPageBySlug(slug string) (core.Page, error) {
	page := core.Page{}

	err := adapter.
		Session.
		Model(&page).
		Where("slug = ?", slug).
		Relation("Fields").
		Select()

	return page, err
}

// GetPagesFromDirectory allows user to fetch pages by their parentId from database
func (adapter Adapter) GetPagesFromDirectory(id int) ([]core.Page, error) {
	pages := []core.Page{}

	err := adapter.
		Session.
		Model(&pages).
		Where("parent_id = ?", id).
		Relation("Fields").
		Select()

	return pages, err
}

// UpdatePage allows user to update page properties
func (adapter Adapter) UpdatePage(id int, data core.PageInput) error {
	if len(data.Validate()) > 0 {
		return core.ErrNotValidated
	}

	page := core.Page{}
	page.ID = id
	page.UpdatedAt = adapter.GetCurrentTime()

	query := adapter.
		Session.
		Model(&page).
		Column("updated_at")

	if data.IsPublished != nil {
		page.IsPublished = *data.IsPublished
		query = query.Column("is_published")
	}
	if data.Name != nil {
		page.Name = *data.Name
		query = query.Column("name")
	}
	if data.ParentID != nil {
		page.ParentID = *data.ParentID
		query = query.Column("parent_id")
	}
	if data.Slug != nil {
		page.Slug = *data.Slug
		query = query.Column("slug")
	}

	_, err := query.
		WherePK().
		Update()

	return err
}

// UpdatePageField allows user to update page's field's properties
func (adapter Adapter) UpdatePageField(id int, data core.PageFieldInput) error {
	if len(data.Validate()) > 0 {
		return core.ErrNotValidated
	}

	pageField := core.PageField{}
	pageField.ID = id
	pageField.UpdatedAt = adapter.GetCurrentTime()

	query := adapter.
		Session.
		Model(&pageField).
		Column("updated_at")

	if data.Name != nil {
		pageField.Name = *data.Name
		query = query.Column("name")
	}
	if data.Value != nil {
		pageField.Value = *data.Value
		query = query.Column("value")
	}

	_, err := query.
		WherePK().
		Update()

	return err
}

// RemovePage removes page from database
func (adapter Adapter) RemovePage(id int) error {
	_, err := adapter.
		Session.
		Exec("DELETE FROM pages WHERE id = ?", id)
	return err
}

// RemovePageField removes field from page
func (adapter Adapter) RemovePageField(id int) error {
	pageField := core.PageField{}
	pageField.ID = id
	err := adapter.
		Session.
		Model(&pageField).
		WherePK().
		Select()

	if err != nil {
		return err
	}

	_, err = adapter.
		Session.
		Model(&pageField).
		WherePK().
		Delete()

	if err != nil {
		return err
	}

	_, err = adapter.
		Session.
		Exec(
			"UPDATE pages SET updated_at = ? WHERE id = ?",
			adapter.GetCurrentTime(),
			pageField.PageID,
		)

	return err
}
