package testing

import (
	"github.com/dominik-zeglen/inkster/core"
)

func createDirectory(directory core.Directory, id int, createdAt string, updatedAt string) core.Directory {
	directory.ID = id
	directory.CreatedAt = createdAt
	directory.UpdatedAt = updatedAt

	return directory
}

func CreateDirectories() []core.Directory {
	directories := []core.Directory{
		core.Directory{Name: "Directory 1"},
		core.Directory{Name: "Directory 2"},
		core.Directory{Name: "Directory 3"},
		core.Directory{Name: "Directory 1.1", ParentID: 1},
	}

	directories[0] = createDirectory(
		directories[0],
		1,
		"2007-07-07T10:00:00.000Z",
		"2007-07-07T10:00:00.000Z",
	)
	directories[1] = createDirectory(
		directories[1],
		2,
		"2007-07-07T11:00:00.000Z",
		"2007-07-07T11:00:00.000Z",
	)
	directories[2] = createDirectory(
		directories[2],
		3,
		"2007-07-07T12:00:00.000Z",
		"2007-07-07T12:00:00.000Z",
	)
	directories[3] = createDirectory(
		directories[3],
		4,
		"2007-07-07T13:00:00.000Z",
		"2007-07-07T13:00:00.000Z",
	)

	return directories
}
