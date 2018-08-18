package mongodb

import (
	"github.com/dominik-zeglen/inkster/core"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gosimple/slug"
)

// AddPage puts page in the database
func (adapter Adapter) AddPage(page core.Page) (core.Page, error) {
	err := page.Validate()
	if err != nil {
		return core.Page{}, err
	}
	db, err := mgo.Dial(adapter.ConnectionURI)
	db.SetSafe(&mgo.Safe{})
	defer db.Close()
	if err != nil {
		return core.Page{}, err
	}
	collection := db.DB(adapter.DBName).C("pages")
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
	err = collection.Insert(page)
	return page, err
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
	err := field.Validate()
	if err != nil {
		return err
	}
	db, err := mgo.Dial(adapter.ConnectionURI)
	db.SetSafe(&mgo.Safe{})
	defer db.Close()
	if err != nil {
		return err
	}
	collection := db.DB(adapter.DBName).C("pages")
	found, err := collection.Find(bson.M{
		"_id": pageID,
		"fields": bson.M{
			"$elemMatch": bson.M{
				"name": field.Name,
			},
		},
	}).Count()
	if found != 0 {
		return core.ErrFieldExists(field.Name)
	}
	return collection.UpdateId(pageID, bson.M{
		"$push": bson.M{
			"fields": field,
		},
	})
}

// GetPage allows user to fetch page by ID from database
func (adapter Adapter) GetPage(id bson.ObjectId) (core.Page, error) {
	db, err := mgo.Dial(adapter.ConnectionURI)
	db.SetSafe(&mgo.Safe{})
	defer db.Close()
	if err != nil {
		return core.Page{}, err
	}
	collection := db.DB(adapter.DBName).C("pages")
	var page core.Page
	err = collection.
		FindId(id).
		One(&page)
	if err != nil {
		return core.Page{}, err
	}
	return page, nil
}

// GetPageBySlug allows user to fetch page by slug from database
func (adapter Adapter) GetPageBySlug(slug string) (core.Page, error) {
	db, err := mgo.Dial(adapter.ConnectionURI)
	db.SetSafe(&mgo.Safe{})
	defer db.Close()
	if err != nil {
		return core.Page{}, err
	}
	collection := db.DB(adapter.DBName).C("pages")
	var page core.Page
	err = collection.
		Find(bson.M{"slug": slug}).
		One(&page)
	if err != nil {
		return core.Page{}, err
	}
	return page, nil
}

// GetPagesFromDirectory allows user to fetch pages by their parentId from database
func (adapter Adapter) GetPagesFromDirectory(id bson.ObjectId) ([]core.Page, error) {
	db, err := mgo.Dial(adapter.ConnectionURI)
	db.SetSafe(&mgo.Safe{})
	defer db.Close()
	if err != nil {
		return nil, err
	}
	collection := db.DB(adapter.DBName).C("pages")
	var pages []core.Page
	err = collection.
		Find(bson.M{"parentId": id}).
		All(&pages)
	if err != nil {
		return nil, err
	}
	return pages, nil
}

// UpdatePage allows user to update page properties
func (adapter Adapter) UpdatePage(pageID bson.ObjectId, data core.PageInput) error {
	db, err := mgo.Dial(adapter.ConnectionURI)
	db.SetSafe(&mgo.Safe{})
	defer db.Close()
	if err != nil {
		return err
	}
	collection := db.DB(adapter.DBName).C("pages")

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
		"$set": data,
	})
}

// UpdatePageField removes field from page
func (adapter Adapter) UpdatePageField(pageID bson.ObjectId, pageFieldName string, data string) error {
	db, err := mgo.Dial(adapter.ConnectionURI)
	db.SetSafe(&mgo.Safe{})
	defer db.Close()
	if err != nil {
		return err
	}
	collection := db.DB(adapter.DBName).C("pages")
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
			"fields.$.value": data,
		},
	})
}

// RemovePage removes page from database
func (adapter Adapter) RemovePage(pageID bson.ObjectId) error {
	db, err := mgo.Dial(adapter.ConnectionURI)
	db.SetSafe(&mgo.Safe{})
	defer db.Close()
	if err != nil {
		return err
	}
	collection := db.DB(adapter.DBName).C("pages")
	return collection.RemoveId(pageID)
}

// RemovePageField removes field from page
func (adapter Adapter) RemovePageField(pageID bson.ObjectId, pageFieldName string) error {
	db, err := mgo.Dial(adapter.ConnectionURI)
	db.SetSafe(&mgo.Safe{})
	defer db.Close()
	if err != nil {
		return err
	}
	collection := db.DB(adapter.DBName).C("pages")
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
		"$pull": bson.M{
			"fields": bson.M{
				"name": pageFieldName,
			},
		},
	})
}

func (adapter Adapter) ValidatePage(page core.Page) error {
	return page.Validate()
}
