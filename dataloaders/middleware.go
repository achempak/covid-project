package dataloaders

import (
	"context"
	"covidProject/db"
	"net/http"
)

// Middleware stores Loaders as a request-scoped context value.
func Middleware(repo db.Repository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			loaders := newLoaders(ctx, repo)
			augmentedCtx := context.WithValue(ctx, key, loaders)
			r = r.WithContext(augmentedCtx)
			next.ServeHTTP(w, r)
		})
	}
}