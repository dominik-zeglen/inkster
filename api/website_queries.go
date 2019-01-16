package api

import (
	"context"

	"github.com/dominik-zeglen/inkster/core"
)

func (res *Resolver) Website(ctx context.Context) (*websiteResolver, error) {
	website := core.Website{}
	website.ID = core.WEBSITE_DB_ID

	err := res.
		dataSource.
		DB().
		Model(&website).
		WherePK().
		Select()

	if err != nil {
		return nil, err
	}

	return &websiteResolver{
		data: &website,
	}, nil
}
