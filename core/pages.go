package core

import (
	"fmt"
)

// Page is a object representing site content
type Page struct {
	BaseModel
	AuthorID    int         `sql:",notnull" json:"authorId" validate:"required"`
	Author      *User       `json:"-"`
	Name        string      `sql:",notnull" json:"name" validate:"required,min=3"`
	Slug        string      `sql:",notnull" json:"slug" validate:"omitempty,slug,min=3"`
	ParentID    int         `sql:",notnull" json:"parentId" validate:"required"`
	Parent      *Directory  `json:"-"`
	IsPublished bool        `sql:",notnull" json:"isPublished"`
	Fields      []PageField `json:"fields" sql:"type:jsonb" validate:"dive"`
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

// PageField represents a single field in page
type PageField struct {
	Type  string `json:"type" validate:"required,oneof=text longText image file"`
	Slug  string `json:"slug" validate:"required,slug"`
	Value string `json:"value"`
}

// Validate checks if field can be put into database
func (field PageField) Validate() []ValidationError {
	return ValidateModel(field)
}

func (field PageField) String() string {
	return fmt.Sprintf("PageField<%s (%s)>",
		field.Slug,
		field.Type,
	)
}

// PageInput is transactional model of an creation properties
type PageInput struct {
	Name        *string `validate:"omitempty,min=3"`
	Slug        *string `validate:"omitempty,min=3"`
	ParentID    *int
	IsPublished *bool
	Fields      *[]PageField `validate:"omitempty,dive"`
}

func (pageInput PageInput) Validate() []ValidationError {
	return ValidateModel(pageInput)
}

// UpdatePageFieldArguments is transactional model of an update properties
type PageFieldInput struct {
	Name  *string `validate:"min=3"`
	Value *string
}

func (pageFieldInput PageFieldInput) Validate() []ValidationError {
	return ValidateModel(pageFieldInput)
}
