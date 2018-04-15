package test

import (
	"testing"

	"github.com/dominik-zeglen/foxxy/core"
)

// TODO: Write a lot of tests
func TestGetContainer(t *testing.T) {
	container, _ := core.GetContainer(1)

	if container.Id != 1 {
		t.Error("IDs not matching")
	}
	if container.Name != "Lorem" {
		t.Error("Names not matching")
	}
}

func TestGetContainerList(t *testing.T) {
	containers, _ := core.GetContainerList()
	if len(containers) != 5 {
		t.Error("List count not matching")
	}
}

func TestAddContainer(t *testing.T) {
	container := core.Container{Name: "Lorem"}
	container, err := core.AddContainer(container)
	if err != nil {
		t.Error(err)
	}
	if container.Name != "Lorem" {
		t.Error("Names not matching")
	}
}

func TestRemoveContainer(t *testing.T) {
	err := core.RemoveContainer(1)
	if err != nil {
		t.Error(err)
	}
	_, err = core.GetContainer(1)
	if err == nil {
		t.Error("Threw no error")
	}
}
