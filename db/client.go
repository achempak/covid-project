package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Pool struct {
	*pgxpool.Pool
}

// Client is a pgxpool client
func Client(ctx context.Context) (*Pool, error) {
	conn, err := pgxpool.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, fmt.Errorf("Error connecting to database:\n%s", err)
	}
	return &Pool{conn}, nil
}


// PGClient is a database/sql client
func PGClient() (*sql.DB, error) {
	return sql.Open("postgres", os.Getenv("DATABASE_NAME"))
}