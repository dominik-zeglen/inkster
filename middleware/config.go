package middleware

import (
	"context"
	"net/http"

	appConfig "github.com/dominik-zeglen/inkster/config"
)

// ConfigContextKey defines key holding website config data in request context
const ConfigContextKey = ContextKey("config")

// WithConfig provides website config data to request
func WithConfig(next http.Handler, config appConfig.Config) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), ConfigContextKey, config)

			next.ServeHTTP(w, r.WithContext(ctx))
			return
		},
	)
}
