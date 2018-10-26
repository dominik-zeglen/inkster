package core

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
)

// Directory is used to create tree-like structures
type Directory struct {
	BaseModel   `bson:",inline"`
	Name        string        `json:"name" validate:"required,min=3"`
	ParentID    bson.ObjectId `bson:"parentId,omitempty" json:"parentId"`
	IsPublished bool          `bson:"isPublished" json:"isPublished"`
}

// DirectoryInput is transactional model of an update properties
type DirectoryInput struct {
	Name        *string        `bson:"name,omitempty" validate:"min=3"`
	ParentID    *bson.ObjectId `bson:"parentId,omitempty"`
	IsPublished *bool          `bson:"isPublished,omitempty"`
}

func (directory Directory) String() string {
	return fmt.Sprintf("Directory<%s>", directory.Name)
}

func (directory Directory) Validate() []ValidationError {
	return ValidateModel(directory)
}
