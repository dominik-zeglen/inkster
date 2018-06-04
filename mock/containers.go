package mock

import (
	"fmt"

	"github.com/dominik-zeglen/ecoknow/core"
	"github.com/globalsign/mgo/bson"
)

func (adapter Adapter) findContainer(id bson.ObjectId) (int, error) {
	for index := range adapter.containers {
		if adapter.containers[index].ID == id {
			return index, nil
		}
	}
	return 0, fmt.Errorf("Container %s does not exist", id)
}

// AddContainer puts container in a in-memory array
func (adapter Adapter) AddContainer(container core.Container) (core.Container, error) {
	if container.ID == "" {
		container.ID = bson.NewObjectId()
	} else {
		_, err := adapter.findContainer(container.ID)
		if err == nil {
			return core.Container{}, fmt.Errorf("Could not add container with ID: %s; container already exists", container.ID)
		}
	}
	adapter.containers = append(adapter.containers, container)
	return container, nil
}

// GetContainer gets container from a in-memory array
func (adapter Adapter) GetContainer(id bson.ObjectId) (core.Container, error) {
	index, err := adapter.findContainer(id)
	if err != nil {
		return core.Container{}, err
	}
	return adapter.containers[index], nil
}

// GetContainerList gets all containers from a in-memory array
func (adapter Adapter) GetContainerList() ([]core.Container, error) {
	return adapter.containers, nil
}

// GetRootContainerList gets only containers from a in-memory array that don't have parent
func (adapter Adapter) GetRootContainerList() ([]core.Container, error) {
	result := []core.Container{}
	for k := range adapter.containers {
		if adapter.containers[k].ParentID == "" {
			result = append(result, adapter.containers[k])
		}
	}
	return result, nil
}

// GetContainerChildrenList gets containers from a in-memory array which
// ParentID equals to function id parameter
func (adapter Adapter) GetContainerChildrenList(id bson.ObjectId) ([]core.Container, error) {
	result := []core.Container{}
	for k := range adapter.containers {
		if adapter.containers[k].ParentID == id {
			result = append(result, adapter.containers[k])
		}
	}
	return result, nil
}

// UpdateContainer updates container with given properties
func (adapter Adapter) UpdateContainer(id bson.ObjectId, data core.ContainerInput) error {
	index, err := adapter.findContainer(id)
	if err != nil {
		return err
	}
	if data.Name != nil {
		adapter.containers[index].Name = *data.Name
	}
	if data.ParentID != nil {
		adapter.containers[index].ParentID = *data.ParentID
	}
	return nil
}

// RemoveContainer removes container from a in-memory array
func (adapter Adapter) RemoveContainer(id bson.ObjectId) error {
	for index := range adapter.containers {
		if adapter.containers[index].ID == id {
			adapter.containers = append(adapter.containers[:index], adapter.containers[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not remove container with ID: %s. Container does not exist", id)
}
