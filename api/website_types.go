package api

import (
	"github.com/dominik-zeglen/inkster/core"
)

type websiteResolver struct {
	data *core.Website
}

func (res *websiteResolver) Name() string {
	return res.data.Name
}

func (res *websiteResolver) Description() string {
	return res.data.Description
}

func (res *websiteResolver) Domain() string {
	return res.data.Domain
}
