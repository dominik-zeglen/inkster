package mock

import (
	"fmt"

	"github.com/dominik-zeglen/ecoknow/core"
	"github.com/globalsign/mgo/bson"
)

var containers []core.Container

// AddContainer puts container in a in-memory array
func (adapter *Adapter) AddContainer(container core.Container) (core.Container, error) {
	if container.ID == "" {
		container.ID = bson.NewObjectId()
	}
	found := false
	for id := range containers {
		if containers[id].ID == container.ID {
			found = true
		}
	}
	if found != false {
		return core.Container{}, fmt.Errorf("Could not add container with ID: %s; container already exists", container.ID)
	}
	containers = append(containers, container)
	return container, nil
}

// GetContainer gets container from a in-memory array
func (adapter *Adapter) GetContainer(id bson.ObjectId) (core.Container, error) {
	for containerID := range containers {
		if containers[containerID].ID == bson.ObjectId(id) {
			return containers[containerID], nil
		}
	}
	return core.Container{}, fmt.Errorf("Could not find container with ID: %s", id)
}

// GetContainerList gets all containers from a in-memory array
func (adapter *Adapter) GetContainerList() ([]core.Container, error) {
	var result = []core.Container{}
	for k := range containers {
		result = append(result, containers[k])
	}
	return result, nil
}

// GetRootContainerList gets only containers from a in-memory array that don't have parent
func (adapter *Adapter) GetRootContainerList() ([]core.Container, error) {
	result := []core.Container{}
	for k := range containers {
		if containers[k].ParentID == "" {
			result = append(result, containers[k])
		}
	}
	return result, nil
}

// GetContainerChildrenList gets containers from a in-memory array which
// ParentID equals to function id parameter
func (adapter *Adapter) GetContainerChildrenList(id bson.ObjectId) ([]core.Container, error) {
	result := []core.Container{}
	for k := range containers {
		if containers[k].ParentID == id {
			result = append(result, containers[k])
		}
	}
	return result, nil
}

// RemoveContainer removes container from a in-memory array
func (adapter *Adapter) RemoveContainer(id bson.ObjectId) error {
	for index := range containers {
		if containers[index].ID == id {
			containers = append(containers[:index], containers[:index+1]...)
			return nil
		}
	}
	return fmt.Errorf("Could not remove container with ID: %s. Container does not exist", id)
}
