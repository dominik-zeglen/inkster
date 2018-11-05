package core

import (
	"fmt"
)

// Page is a object representing site content
type Page struct {
	BaseModel   `bson:",inline"`
	Name        string      `json:"name" validate:"required,min=3"`
	Slug        string      `json:"slug" validate:"omitempty,slug,min=3"`
	ParentID    string      `bson:"parentId" json:"parentId" validate:"required"`
	IsPublished bool        `bson:"isPublished" json:"isPublished"`
	Fields      []PageField `json:"fields" validate:"dive"`
}

func (page Page) String() string {
	return fmt.Sprintf("Page<%s: %s>", page.Name, page.Slug)
}

// Validate checks if page can be put into database
func (page Page) Validate() []ValidationError {
	return ValidateModel(page)
}

func NewPage() Page {
	page := Page{}
	return page
}

// UpdatePageArguments is transactional model of an update properties
type UpdatePageArguments struct {
	Name        *string `bson:"name",omitempty`
	Slug        *string `bson:"slug",omitempty`
	IsPublished *bool   `bson:"isPublished",omitempty`
}

// PageField represents a single field in template
type PageField struct {
	Type  string `json:"type" validate:"required,oneof=text longText image file"`
	Name  string `json:"name" validate:"required,min=3"`
	Value string `json:"value"`
}

// Validate checks if field can be put into database
func (field PageField) Validate() []ValidationError {
	return ValidateModel(field)
}

func (field PageField) String() string {
	return fmt.Sprintf("PageField<%s: %s (%s)>",
		field.Name,
		field.Type,
		field.Value[:20],
	)
}

// PageInput is transactional model of an creation properties
type PageInput struct {
	Name        *string      `bson:"name,omitempty" validate:"min=3"`
	Slug        *string      `bson:"slug,omitempty" validate:"min=3"`
	ParentID    *string      `bson:"parentId,omitempty"`
	IsPublished *bool        `bson:"isPublished",omitempty`
	Fields      *[]PageField `bson:"fields,omitempty" validate:"dive"`
}

func (pageInput PageInput) Validate() []ValidationError {
	return ValidateModel(pageInput)
}

// UpdatePageFieldArguments is transactional model of an update properties
type UpdatePageFieldArguments struct {
	Name  string `bson:",omitempty" validate:"min=3"`
	Value string `bson:",omitempty"`
}
