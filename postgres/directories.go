package postgres

import (
	"github.com/dominik-zeglen/inkster/core"
)

// AddDirectory puts directory in the database
func (adapter Adapter) AddDirectory(directory core.Directory) (core.Directory, error) {
	errors := directory.Validate()
	if len(errors) > 0 {
		return core.Directory{}, core.ErrNotValidated
	}

	directory.CreatedAt = adapter.GetCurrentTime()
	directory.UpdatedAt = adapter.GetCurrentTime()

	_, err := adapter.
		Session.
		Model(&directory).
		Insert()

	return directory, err
}

// GetDirectory gets directory from the database
func (adapter Adapter) GetDirectory(id int) (core.Directory, error) {
	directory := core.Directory{}
	directory.ID = id

	err := adapter.
		Session.
		Model(&directory).
		WherePK().
		Select()

	return directory, err
}

// GetDirectoryList gets all directories from the database
func (adapter Adapter) GetDirectoryList() ([]core.Directory, error) {
	directories := []core.Directory{}

	err := adapter.
		Session.
		Model(&directories).
		Select()

	return directories, err
}

// GetRootDirectoryList gets only directories from a pg database that don't have parent
func (adapter Adapter) GetRootDirectoryList() ([]core.Directory, error) {
	directories := []core.Directory{}

	err := adapter.
		Session.
		Model(&directories).
		Where("parent_id IS NULL OR parent_id = 0").
		Select()

	return directories, err
}

// GetDirectoryChildrenList gets directories from a pg database which have same parent
func (adapter Adapter) GetDirectoryChildrenList(id int) ([]core.Directory, error) {
	directories := []core.Directory{}

	err := adapter.
		Session.
		Model(&directories).
		Where("parent_id = ?", id).
		Select()

	return directories, err
}

type directoryUpdateInput struct {
	Data      core.DirectoryInput `bson:",inline"`
	UpdatedAt string              `bson:"updatedAt"`
}

// UpdateDirectory allows directory properties updaing
func (adapter Adapter) UpdateDirectory(
	id int,
	data core.DirectoryInput,
) error {
	errors := core.ValidateModel(data)
	if len(errors) > 0 {
		return core.ErrNotValidated
	}

	directory := core.Directory{}
	directory.ID = id
	directory.UpdatedAt = adapter.GetCurrentTime()

	query := adapter.
		Session.
		Model(&directory).
		Column("updated_at")

	if data.IsPublished != nil {
		directory.IsPublished = *data.IsPublished
		query = query.Column("is_published")
	}
	if data.Name != nil {
		directory.Name = *data.Name
		query = query.Column("name")
	}
	if data.ParentID != nil {
		directory.ParentID = *data.ParentID
		query = query.Column("parent_id")
	}

	_, err := query.
		WherePK().
		Update()

	return err
}

// RemoveDirectory removes directory from a pg database
func (adapter Adapter) RemoveDirectory(id int) error {
	_, err := adapter.
		Session.
		Exec("DELETE FROM directories WHERE id = ?", id)

	return err
}
