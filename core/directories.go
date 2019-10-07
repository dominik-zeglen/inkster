package core

import (
	"fmt"
)

// Directory is used to create tree-like structures
type Directory struct {
	BaseModel   `bson:",inline"`
	Name        string     `sql:",notnull" json:"name" validate:"required"`
	ParentID    *int       `sql:",on_delete:CASCADE" bson:"parentId,omitempty" json:"parentId"`
	Parent      *Directory `json:"-"`
	IsPublished bool       `sql:",notnull" bson:"isPublished" json:"isPublished"`
}

func (directory Directory) String() string {
	return fmt.Sprintf("Directory<%s>", directory.Name)
}

func (directory Directory) Validate() []ValidationError {
	validationErrors := ValidateModel(directory)
	if directory.ParentID != nil {
		if directory.ID == *directory.ParentID {
			message := "its ID"
			validationErrors = append(validationErrors, ValidationError{
				Code:  ErrNotEqual,
				Field: "ParentID",
				Param: &message,
			})
		}
	}
	return validationErrors
}
