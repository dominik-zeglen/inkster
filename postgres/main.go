package mongodb

import (
	"time"

	"github.com/dominik-zeglen/inkster/core"
	"github.com/go-pg/pg"
)

// Adapter is an abstraction over mongodb session
type Adapter struct {
	core.Adapter

	GetTime func() string
	Session *pg.DB
}

func (adapter Adapter) GetCurrentTime() string {
	if adapter.GetTime == nil {
		return time.Now().UTC().Format(time.RFC3339)
	}
	return adapter.GetTime()
}

func (adapter Adapter) String() string {
	return "PostgreSQL"
}

func (adapter Adapter) ResetMockDatabase(
	templates []core.Directory,
	templates []core.Template,
	pages []core.Page,
	users []core.User,
) {
	var directoriesQuery []core.Directory
	err := adapter.db.Model(&directoriesQuery).Select()
	if err != nil {
		panic(err)
	}
	_, err = db.Model(&directoriesQuery).Delete()
	if err != nil {
		panic(err)
	}

	err = db.Insert(&directories)
	if err != nil {
		panic(err)
	}

	var templatesQuery []core.Template
	err := adapter.db.Model(&templatesQuery).Select()
	if err != nil {
		panic(err)
	}
	_, err = db.Model(&templatesQuery).Delete()
	if err != nil {
		panic(err)
	}

	err = db.Insert(&templates)
	if err != nil {
		panic(err)
	}

	var pagesQuery []core.Page
	err := adapter.db.Model(&pagesQuery).Select()
	if err != nil {
		panic(err)
	}
	_, err = db.Model(&pagesQuery).Delete()
	if err != nil {
		panic(err)
	}

	err = db.Insert(&pages)
	if err != nil {
		panic(err)
	}

	var usersQuery []core.Users
	err := adapter.db.Model(&usersQuery).Select()
	if err != nil {
		panic(err)
	}
	_, err = db.Model(&usersQuery).Delete()
	if err != nil {
		panic(err)
	}

	err = db.Insert(&users)
	if err != nil {
		panic(err)
	}
}
