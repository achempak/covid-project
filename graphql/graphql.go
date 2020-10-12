package graphql

import (
	"covidProject/db"
	"github.com/99designs/gqlgen/handler"
	"net/http"
	"covidProject/graphql/generated"
)

// NewHandler returns a new graphql endpoint handler.
func NewHandler(repo db.Repository) http.Handler {
	return handler.GraphQL(generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{
			Repository: repo,
		},
	}))
}

// NewPlaygroundHandler returns a new GraphQL Playground handler.
func NewPlaygroundHandler(endpoint string) http.Handler {
	return handler.Playground("GraphQL Playground", endpoint)
}
