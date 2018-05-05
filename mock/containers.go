package mock

import (
	"fmt"

	"github.com/dominik-zeglen/ecoknow/core"
)

var containers = []core.Container{
	core.Container{ID: 1, Name: "Container 1"},
	core.Container{ID: 2, Name: "Container 2"},
	core.Container{ID: 3, Name: "Container 3"},
	core.Container{ID: 4, Name: "Container 1.1", ParentID: 1},
}
var containerIDCounter = int32(5)

// AddContainer puts container in a in-memory array
func (adapter *Adapter) AddContainer(container core.Container) (core.Container, error) {
	if container.ID != 0 {
		if container.ID < containerIDCounter {
			return container, fmt.Errorf("ID cannot be less than current counter: %d < %d", containerIDCounter, container.ID)
		}
		containers = append(containers, container)
		containerIDCounter = container.ID + 1
		return container, nil
	}
	container.ID = containerIDCounter
	containers = append(containers, container)
	containerIDCounter++
	return container, nil
}

// GetContainer gets container from a in-memory array
func (adapter *Adapter) GetContainer(id int32) (core.Container, error) {
	for containerID := range containers {
		if containers[containerID].ID == id {
			return containers[containerID], nil
		}
	}
	return core.Container{}, fmt.Errorf("Could not find container with ID: %d", id)
}

// GetContainerList gets all containers from a in-memory array
func (adapter *Adapter) GetContainerList() ([]core.Container, error) {
	return containers, nil
}

// GetRootContainerList gets only containers from a in-memory array that don't have parent
func (adapter *Adapter) GetRootContainerList() ([]core.Container, error) {
	result := []core.Container{}
	for containerID := range containers {
		if containers[containerID].ParentID == 0 {
			result = append(result, containers[containerID])
		}
	}
	return result, nil
}

// GetContainerChildrenList gets containers from a in-memory array which
// ParentID equals to function id parameter
func (adapter *Adapter) GetContainerChildrenList(id int32) ([]core.Container, error) {
	result := []core.Container{}
	for containerID := range containers {
		if containers[containerID].ParentID == id {
			result = append(result, containers[containerID])
		}
	}
	return result, nil
}

// RemoveContainer removes container from a in-memory array
func (adapter *Adapter) RemoveContainer(id int32) error {
	for containerID := range containers {
		if containers[containerID].ID == id {
			containers = append(containers[:containerID], containers[:containerID+1]...)
			return nil
		}
	}
	return fmt.Errorf("Could not remove container with ID: %d. Container does not exist", id)
}
