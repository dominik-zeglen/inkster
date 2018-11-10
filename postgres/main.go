package postgres

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
	directories []core.Directory,
	templates []core.Template,
	pages []core.Page,
	users []core.User,
) error {
	var directoriesQuery []core.Directory
	var templatesQuery []core.Template
	var pagesQuery []core.Page
	var pageFieldsQuery []core.PageField
	var usersQuery []core.User

	_, err := adapter.
		Session.
		Model(&pageFieldsQuery).
		Where("1=1").
		Delete()
	if err != nil {
		return err
	}

	_, err = adapter.
		Session.
		Model(&pagesQuery).
		Where("1=1").
		Delete()
	if err != nil {
		return err
	}

	_, err = adapter.
		Session.
		Model(&directoriesQuery).
		Where("1=1").
		Delete()
	if err != nil {
		return err
	}

	_, err = adapter.
		Session.
		Model(&templatesQuery).
		Where("1=1").
		Delete()
	if err != nil {
		return err
	}

	_, err = adapter.
		Session.
		Model(&usersQuery).
		Where("1=1").
		Delete()
	if err != nil {
		return err
	}

	err = adapter.Session.Insert(&directories)
	if err != nil {
		return err
	}

	err = adapter.Session.Insert(&templates)
	if err != nil {
		return err
	}

	err = adapter.Session.Insert(&pages)
	if err != nil {
		return err
	}

	err = adapter.Session.Insert(&users)
	if err != nil {
		return err
	}

	pageFields := []core.PageField{}
	for _, page := range pages {
		for fieldIndex, _ := range page.Fields {
			page.Fields[fieldIndex].PageID = page.ID
		}
		pageFields = append(pageFields, page.Fields...)
	}

	err = adapter.Session.Insert(&pageFields)
	return err
}
