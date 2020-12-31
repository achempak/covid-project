package main

import (
	"context"
	"covidProject/dataloaders"
	"covidProject/db"
	"covidProject/graphql"
	"covidProject/logger"
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func main() {
	logger.SetLogLevel(logger.ERROR)

	err := godotenv.Load("../.env")
	if err != nil {
		logger.Error("can't load .env file")
		os.Exit(1)
	}
	ctx := context.Background()
	conn, err := db.Client(ctx)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	defer conn.Close()

	dataDir := os.Getenv("DATA_DIR")
	err = db.PullData(dataDir)
	if err != nil && err.Error() != "already up-to-date" {
		logger.Error(fmt.Sprintf("can't fetch latest covid data from GitHub" +
			"\n%s", err))
	} else {
		err = conn.UpsertAll(ctx)
		if err != nil {
			logger.Error(err)
		}
	}

	// initialize the db
	pg, err := db.PGClient()
	if err != nil {
		panic(err)
	}
	defer pg.Close()
	if err := pg.Ping(); err != nil {
		panic(err)
	}

	// initialize the repository
	repo := db.NewRepository(pg)

	// initialize the dataloaders
	dl := dataloaders.NewRetriever()

	// configure the server
	mux := http.NewServeMux()
	mux.Handle("/", graphql.NewPlaygroundHandler("/query"))
	dlMiddleware := dataloaders.Middleware(repo)
	queryHandler := graphql.NewHandler(repo, dl)
	mux.Handle("/query", dlMiddleware(queryHandler))

	// run the server
	port := ":8080"
	fmt.Fprintf(os.Stdout, "🚀 Server ready at http://localhost%s\n", port)
	fmt.Fprintln(os.Stderr, http.ListenAndServe(port, mux))
}
