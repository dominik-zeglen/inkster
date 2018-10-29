package mongodb

import (
	"github.com/dominik-zeglen/inkster/core"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gosimple/slug"
)

// AddPage puts page in the database
func (adapter Adapter) AddPage(page core.Page) (core.Page, error) {
	errs := page.Validate()
	if len(errs) > 0 {
		return core.Page{}, core.ErrNotValidated
	}
	session := adapter.Session.Copy()
	session.SetSafe(&mgo.Safe{})
	defer session.Close()

	collection := session.DB(adapter.DBName).C("pages")
	if page.Slug != "" {
		found, err := collection.
			Find(bson.M{"slug": page.Slug}).
			Count()
		if err != nil {
			return core.Page{}, err
		}
		if found > 0 {
			return core.Page{}, core.ErrPageExists(page.Name)
		}
	}
	if page.ID == "" {
		page.ID = bson.NewObjectId()
	}
	if page.Slug == "" {
		page.Slug = slug.Make(page.Name)
	}
	page.CreatedAt = adapter.GetCurrentTime()
	page.UpdatedAt = adapter.GetCurrentTime()

	return page, collection.Insert(page)
}

// AddPageFromTemplate creates new page based on a chosen template
func (adapter Adapter) AddPageFromTemplate(
	page core.PageInput,
	templateID bson.ObjectId,
) (core.Page, error) {
	template, err := adapter.GetTemplate(templateID)
	if err != nil {
		return core.Page{}, err
	}
	var fields []core.PageField
	for _, field := range template.Fields {
		fields = append(fields, core.PageField{
			Name:  field.Name,
			Type:  field.Type,
			Value: "",
		})
	}
	if page.Name == nil {
		return core.Page{}, core.ErrNoEmpty("name")
	}
	if page.ParentID == nil {
		return core.Page{}, core.ErrNoEmpty("parentID")
	}
	inputPage := core.Page{
		Name:     *page.Name,
		ParentID: *page.ParentID,
		Fields:   fields,
	}
	if page.Slug != nil {
		inputPage.Slug = *page.Slug
	} else {
		slug := slug.Make(*page.Name)
		inputPage.Slug = slug
	}
	return adapter.AddPage(inputPage)
}

// AddPageField adds to page a new field at the end of it's field list
func (adapter Adapter) AddPageField(pageID bson.ObjectId, field core.PageField) error {
	errs := field.Validate()
	if len(errs) > 0 {
		return core.ErrNotValidated
	}

	session := adapter.Session.Copy()
	session.SetSafe(&mgo.Safe{})
	defer session.Close()

	collection := session.DB(adapter.DBName).C("pages")
	found, err := collection.Find(bson.M{
		"_id": pageID,
		"fields": bson.M{
			"$elemMatch": bson.M{
				"name": field.Name,
			},
		},
	}).Count()
	if err != nil {
		return err
	}
	if found != 0 {
		return core.ErrFieldExists(field.Name)
	}
	return collection.UpdateId(pageID, bson.M{
		"$set": bson.M{
			"updatedAt": adapter.GetCurrentTime(),
		},
		"$push": bson.M{
			"fields": field,
		},
	})
}

// GetPage allows user to fetch page by ID from database
func (adapter Adapter) GetPage(id bson.ObjectId) (core.Page, error) {
	session := adapter.Session.Copy()
	session.SetSafe(&mgo.Safe{})
	defer session.Close()

	collection := session.DB(adapter.DBName).C("pages")
	var page core.Page
	err := collection.
		FindId(id).
		One(&page)
	if err != nil {
		return core.Page{}, err
	}
	return page, nil
}

// GetPageBySlug allows user to fetch page by slug from database
func (adapter Adapter) GetPageBySlug(slug string) (core.Page, error) {
	session := adapter.Session.Copy()
	session.SetSafe(&mgo.Safe{})
	defer session.Close()

	collection := session.DB(adapter.DBName).C("pages")
	var page core.Page
	err := collection.
		Find(bson.M{"slug": slug}).
		One(&page)
	if err != nil {
		return core.Page{}, err
	}
	return page, nil
}

// GetPagesFromDirectory allows user to fetch pages by their parentId from database
func (adapter Adapter) GetPagesFromDirectory(id bson.ObjectId) ([]core.Page, error) {
	session := adapter.Session.Copy()
	session.SetSafe(&mgo.Safe{})
	defer session.Close()

	collection := session.DB(adapter.DBName).C("pages")
	var pages []core.Page
	err := collection.
		Find(bson.M{"parentId": id}).
		All(&pages)
	if err != nil {
		return nil, err
	}
	return pages, nil
}

type pageUpdateInput struct {
	Data      core.PageInput `bson:",inline"`
	UpdatedAt string         `bson:"updatedAt"`
}

// UpdatePage allows user to update page properties
func (adapter Adapter) UpdatePage(pageID bson.ObjectId, data core.PageInput) error {
	session := adapter.Session.Copy()
	session.SetSafe(&mgo.Safe{})
	defer session.Close()

	collection := session.DB(adapter.DBName).C("pages")

	if data.Slug != nil {
		count, err := collection.
			Find(bson.M{
				"_id": bson.M{
					"$ne": pageID,
				},
				"slug": *data.Slug,
			}).
			Count()
		if err != nil {
			return err
		}
		if count != 0 {
			return core.ErrPageExists(*data.Slug)
		}
	}
	return collection.UpdateId(pageID, bson.M{
		"$set": pageUpdateInput{
			Data:      data,
			UpdatedAt: adapter.GetCurrentTime(),
		},
	})
}

// UpdatePageField removes field from page
func (adapter Adapter) UpdatePageField(pageID bson.ObjectId, pageFieldName string, data string) error {
	session := adapter.Session.Copy()
	session.SetSafe(&mgo.Safe{})
	defer session.Close()

	collection := session.DB(adapter.DBName).C("pages")
	found, err := collection.Find(bson.M{
		"_id": pageID,
		"fields": bson.M{
			"$elemMatch": bson.M{
				"name": pageFieldName,
			},
		},
	}).Count()
	if err != nil {
		return err
	}
	if found == 0 {
		return core.ErrNoField(pageFieldName)
	}
	return collection.Update(bson.M{
		"_id":         pageID,
		"fields.name": pageFieldName,
	}, bson.M{
		"$set": bson.M{
			"updatedAt":      adapter.GetCurrentTime(),
			"fields.$.value": data,
		},
	})
}

// RemovePage removes page from database
func (adapter Adapter) RemovePage(pageID bson.ObjectId) error {
	session := adapter.Session.Copy()
	session.SetSafe(&mgo.Safe{})
	defer session.Close()

	collection := session.DB(adapter.DBName).C("pages")
	return collection.RemoveId(pageID)
}

// RemovePageField removes field from page
func (adapter Adapter) RemovePageField(pageID bson.ObjectId, pageFieldName string) error {
	session := adapter.Session.Copy()
	session.SetSafe(&mgo.Safe{})
	defer session.Close()

	collection := session.DB(adapter.DBName).C("pages")
	found, err := collection.Find(bson.M{
		"_id": pageID,
		"fields": bson.M{
			"$elemMatch": bson.M{
				"name": pageFieldName,
			},
		},
	}).Count()
	if err != nil {
		return err
	}
	if found == 0 {
		return core.ErrNoField(pageFieldName)
	}
	return collection.UpdateId(pageID, bson.M{
		"$set": bson.M{
			"updatedAt": adapter.GetCurrentTime(),
		},
		"$pull": bson.M{
			"fields": bson.M{
				"name": pageFieldName,
			},
		},
	})
}
