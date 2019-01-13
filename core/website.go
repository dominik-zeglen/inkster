package core

import (
	"fmt"
)

// Website model stores info about managed website
// Does not compose BaseModel struct because of its singleton nature
type Website struct {
	ID          string `json:"id"`
	Name        string `sql:",notnull" json:"name" validate:"required,min=3"`
	Description string `json:"description"`
	Domain      string `sql:",notnull" json:"domain" validate:"required,url"`
}

func (website Website) String() string {
	return fmt.Sprintf("Website<%s>", website.Name)
}

func (website Website) Validate() []ValidationError {
	return ValidateModel(website)
}

const (
	WEBSITE_DB_ID = "default"
)
