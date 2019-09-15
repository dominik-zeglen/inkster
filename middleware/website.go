package middleware

import (
	"context"
	"net/http"

	"github.com/dominik-zeglen/inkster/core"
)

// WebsiteContextKey defines key holding website data in request context
const WebsiteContextKey = ContextKey("website")

// WithWebsite provides website data to request
func WithWebsite(next http.Handler, dataSource core.DataContext) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			website := core.Website{}
			website.ID = core.WEBSITE_DB_ID

			err := dataSource.
				DB().
				Model(&website).
				WherePK().
				Select()

			if err != nil {
				panic(err)
			}

			ctx := context.WithValue(r.Context(), WebsiteContextKey, website)

			next.ServeHTTP(w, r.WithContext(ctx))
			return
		},
	)
}
