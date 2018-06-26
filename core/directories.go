package core

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
)

// Directory is used to create tree-like structures
type Directory struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name     string        `json:"name"`
	ParentID bson.ObjectId `bson:"parentId,omitempty" json:"parentId"`
}

// DirectoryInput is transactional model of an update properties
type DirectoryInput struct {
	Name     *string        `bson:",omitempty"`
	ParentID *bson.ObjectId `bson:"parentId,omitempty"`
}

func (directory Directory) String() string {
	return fmt.Sprintf("Directory<%s>", directory.Name)
}
