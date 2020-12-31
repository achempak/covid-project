package graphql

import (
	"covidProject/dataloaders"
	"covidProject/db"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"net/http"
	"covidProject/graphql/generated"
)

// NewHandler returns a new graphql endpoint handler.
func NewHandler(repo db.Repository, dl dataloaders.Retriever) http.Handler {
	return handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{
			Repository: repo,
			DataLoaders: dl,
		},
	}))
}

// NewPlaygroundHandler returns a new GraphQL Playground handler.
func NewPlaygroundHandler(endpoint string) http.Handler {
	return playground.Handler("GraphQL Playground", endpoint)
}
