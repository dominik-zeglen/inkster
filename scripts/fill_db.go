package util

import "github.com/dominik-zeglen/foxxy/core"

func FillDB() {
	containerNames := []string{"Lorem", "Ipsum", "Dolor", "Sit", "Amet"}
	for _, name := range containerNames {
		container := core.Container{Name: name}
		container, err := core.AddContainer(container)
		if err != nil {
			panic(err)
		}
	}

	containerChildrenNames := []string{"Lorem Lorem", "Lorem Ipsum", "Lorem Dolor", "Lorem Sit", "Lorem Amet"}
	for _, name := range containerChildrenNames {
		container := core.Container{Name: name, ParentId: 1}
		container, err := core.AddContainer(container)
		if err != nil {
			panic(err)
		}
	}
}
