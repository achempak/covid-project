package main

import (
	"context"
	"covidProject/logger"
	"fmt"
	"github.com/joho/godotenv"
	"os"

	"covidProject/db"
)

func main() {
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
}
