package core

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
)

// Container is used to create tree-like structures
type Container struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name     string        `json:"name"`
	ParentID bson.ObjectId `bson:"parentId,omitempty" json:"parentId"`
}

// ContainerInput is transactional model of an update properties
type ContainerInput struct {
	Name     *string        `bson:",omitempty"`
	ParentID *bson.ObjectId `bson:"parentId,omitempty"`
}

func (container Container) String() string {
	return fmt.Sprintf("Container<%s>", container.Name)
}
