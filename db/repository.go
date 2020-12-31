package db

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"time"
)

// Repository is the application's data layer functionality.
type Repository interface {
	GetCases(ctx context.Context) ([]CovidUsaCasesByDate, error)
	GetCasesByUID(ctx context.Context, uid int64) ([]CovidUsaCasesByDate, error)
	GetCasesByState(ctx context.Context, state sql.NullString) ([]CovidUsaCasesByDate, error)
	GetCasesByDate(ctx context.Context, date time.Time) ([]CovidUsaCasesByDate, error)
	GetCasesSinceDate(ctx context.Context, date time.Time) ([]CovidUsaCasesByDate, error)
	GetCaseByUIDOnDate(ctx context.Context, arg GetCaseByUIDOnDateParams) (CovidUsaCasesByDate, error)
	GetCasesByUIDSinceDate(ctx context.Context, arg GetCasesByUIDSinceDateParams) ([]CovidUsaCasesByDate, error)
	GetLocationByUID(ctx context.Context, uid int64) (CovidUsaLocation, error)
	GetLocationByName(ctx context.Context, location sql.NullString) ([]CovidUsaLocation, error)
	GetLocations(ctx context.Context) ([]CovidUsaLocation, error)
	ListLocationsByCaseIDs(ctx context.Context, caseIDs []int64) ([]ListLocationsByCaseIDsRow, error)
}

type repoSvc struct {
	*Queries
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repoSvc{
		Queries: New(db),
		db: db,
	}
}