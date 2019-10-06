package api

import (
	"context"

	"github.com/dominik-zeglen/inkster/core"
	"github.com/dominik-zeglen/inkster/middleware"
)

func (res *Resolver) Website(ctx context.Context) *websiteResolver {
	website := ctx.Value(middleware.WebsiteContextKey).(core.Website)

	return &websiteResolver{
		data: &website,
	}
}
