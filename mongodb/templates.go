package mongodb

import (
	"github.com/dominik-zeglen/inkster/core"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// AddTemplate puts template in the database
func (adapter Adapter) AddTemplate(template core.Template) (core.Template, error) {
	err := template.Validate()
	if err != nil {
		return core.Template{}, err
	}
	db, err := mgo.Dial(adapter.ConnectionURI)
	db.SetSafe(&mgo.Safe{})
	defer db.Close()
	if err != nil {
		return core.Template{}, err
	}
	collection := db.DB(adapter.DBName).C("templates")
	found, err := collection.
		Find(bson.M{"name": template.Name}).
		Count()
	if found > 0 {
		return core.Template{}, core.ErrTemplateExists(template.Name)
	}
	if template.ID == "" {
		template.ID = bson.NewObjectId()
	}
	template.CreatedAt = adapter.GetCurrentTime()
	template.UpdatedAt = adapter.GetCurrentTime()

	err = collection.Insert(template)
	return template, err
}

// AddTemplateField adds to template a new field at the end of it's field list
func (adapter Adapter) AddTemplateField(templateID bson.ObjectId, field core.TemplateField) error {
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
	collection := db.DB(adapter.DBName).C("templates")
	found, err := collection.Find(bson.M{
		"_id": templateID,
		"fields": bson.M{
			"$elemMatch": bson.M{
				"name": field.Name,
			},
		},
	}).Count()
	if found > 0 {
		return core.ErrFieldExists(field.Name)
	}
	return collection.UpdateId(templateID, bson.M{
		"$set": bson.M{
			"updatedAt": adapter.GetCurrentTime(),
		},
		"$push": bson.M{
			"fields": field,
		},
	})
}

// GetTemplate allows user to fetch template from database
func (adapter Adapter) GetTemplate(templateID bson.ObjectId) (core.Template, error) {
	db, err := mgo.Dial(adapter.ConnectionURI)
	db.SetSafe(&mgo.Safe{})
	defer db.Close()
	if err != nil {
		return core.Template{}, err
	}
	collection := db.DB(adapter.DBName).C("templates")
	var template core.Template
	err = collection.
		FindId(templateID).
		One(&template)
	return template, err
}

// GetTemplateList allows user to fetch all templates from database
func (adapter Adapter) GetTemplateList() ([]core.Template, error) {
	db, err := mgo.Dial(adapter.ConnectionURI)
	db.SetSafe(&mgo.Safe{})
	defer db.Close()
	if err != nil {
		return []core.Template{}, err
	}
	collection := db.DB(adapter.DBName).C("templates")
	var templates []core.Template
	err = collection.
		Find(bson.M{}).
		All(&templates)
	return templates, err
}

type templateUpdateInput struct {
	Data      core.TemplateInput `bson:",inline"`
	UpdatedAt string             `bson:"updatedAt"`
}

// UpdateTemplate allows user to update template properties
func (adapter Adapter) UpdateTemplate(templateID bson.ObjectId, data core.TemplateInput) error {
	db, err := mgo.Dial(adapter.ConnectionURI)
	db.SetSafe(&mgo.Safe{})
	defer db.Close()
	if err != nil {
		return err
	}
	collection := db.DB(adapter.DBName).C("templates")
	found, err := collection.
		Find(bson.M{"name": data.Name}).
		Count()
	if found > 0 {
		return core.ErrTemplateExists(data.Name)
	}
	return collection.UpdateId(templateID, bson.M{
		"$set": templateUpdateInput{
			Data:      data,
			UpdatedAt: adapter.GetCurrentTime(),
		},
	})
}

// RemoveTemplate removes template from database
func (adapter Adapter) RemoveTemplate(templateID bson.ObjectId) error {
	db, err := mgo.Dial(adapter.ConnectionURI)
	db.SetSafe(&mgo.Safe{})
	defer db.Close()
	if err != nil {
		return err
	}
	collection := db.DB(adapter.DBName).C("templates")
	return collection.RemoveId(templateID)
}

// RemoveTemplateField removes field from template
func (adapter Adapter) RemoveTemplateField(templateID bson.ObjectId, templateFieldName string) error {
	db, err := mgo.Dial(adapter.ConnectionURI)
	db.SetSafe(&mgo.Safe{})
	defer db.Close()
	if err != nil {
		return err
	}
	collection := db.DB(adapter.DBName).C("templates")
	found, err := collection.Find(bson.M{
		"_id": templateID,
		"fields": bson.M{
			"$elemMatch": bson.M{
				"name": templateFieldName,
			},
		},
	}).Count()
	if err != nil {
		return err
	}
	if found == 0 {
		return core.ErrNoField(templateFieldName)
	}
	return collection.UpdateId(templateID, bson.M{
		"$set": bson.M{
			"updatedAt": adapter.GetCurrentTime(),
		},
		"$pull": bson.M{
			"fields": bson.M{
				"name": templateFieldName,
			},
		},
	})
}
