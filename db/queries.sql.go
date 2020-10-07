// Code generated by sqlc. DO NOT EDIT.
// source: queries.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const getCasesByDate = `-- name: GetCasesByDate :many
SELECT c.Last_Update, c.Confirmed, c.Deaths, c.Recovered, c.Active, c.Incident_Rate, c.People_Tested, c.People_Hospitalized, c.Mortality_Rate, c.uid, c.Testing_Rate, c.Hospitalization_Rate, c.Created_At, l."Province_State" FROM covid_usa.cases_by_date c
LEFT JOIN covid_usa.locations l on l.uid = c.uid
WHERE c."Created_At" = $1
`

type GetCasesByDateRow struct {
	LastUpdate          sql.NullTime
	Confirmed           sql.NullInt32
	Deaths              sql.NullInt32
	Recovered           sql.NullFloat64
	Active              sql.NullFloat64
	IncidentRate        sql.NullString
	PeopleTested        sql.NullFloat64
	PeopleHospitalized  sql.NullFloat64
	MortalityRate       sql.NullString
	Uid                 int64
	TestingRate         sql.NullString
	HospitalizationRate sql.NullString
	CreatedAt           time.Time
	ProvinceState       sql.NullString
}

func (q *Queries) GetCasesByDate(ctx context.Context, createdAt time.Time) ([]GetCasesByDateRow, error) {
	rows, err := q.db.QueryContext(ctx, getCasesByDate, createdAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCasesByDateRow
	for rows.Next() {
		var i GetCasesByDateRow
		if err := rows.Scan(
			&i.LastUpdate,
			&i.Confirmed,
			&i.Deaths,
			&i.Recovered,
			&i.Active,
			&i.IncidentRate,
			&i.PeopleTested,
			&i.PeopleHospitalized,
			&i.MortalityRate,
			&i.Uid,
			&i.TestingRate,
			&i.HospitalizationRate,
			&i.CreatedAt,
			&i.ProvinceState,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCasesByState = `-- name: GetCasesByState :many
SELECT c.Last_Update, c.Confirmed, c.Deaths, c.Recovered, c.Active, c.Incident_Rate, c.People_Tested, c.People_Hospitalized, c.Mortality_Rate, c.uid, c.Testing_Rate, c.Hospitalization_Rate, c.Created_At, l."Province_State" FROM covid_usa.cases_by_date c
LEFT JOIN covid_usa.locations l on l.uid = c.uid
WHERE l."Province_State" = $1
`

type GetCasesByStateRow struct {
	LastUpdate          sql.NullTime
	Confirmed           sql.NullInt32
	Deaths              sql.NullInt32
	Recovered           sql.NullFloat64
	Active              sql.NullFloat64
	IncidentRate        sql.NullString
	PeopleTested        sql.NullFloat64
	PeopleHospitalized  sql.NullFloat64
	MortalityRate       sql.NullString
	Uid                 int64
	TestingRate         sql.NullString
	HospitalizationRate sql.NullString
	CreatedAt           time.Time
	ProvinceState       sql.NullString
}

func (q *Queries) GetCasesByState(ctx context.Context, provinceState sql.NullString) ([]GetCasesByStateRow, error) {
	rows, err := q.db.QueryContext(ctx, getCasesByState, provinceState)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCasesByStateRow
	for rows.Next() {
		var i GetCasesByStateRow
		if err := rows.Scan(
			&i.LastUpdate,
			&i.Confirmed,
			&i.Deaths,
			&i.Recovered,
			&i.Active,
			&i.IncidentRate,
			&i.PeopleTested,
			&i.PeopleHospitalized,
			&i.MortalityRate,
			&i.Uid,
			&i.TestingRate,
			&i.HospitalizationRate,
			&i.CreatedAt,
			&i.ProvinceState,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCasesByStateOnDate = `-- name: GetCasesByStateOnDate :many
SELECT c.Last_Update, c.Confirmed, c.Deaths, c.Recovered, c.Active, c.Incident_Rate, c.People_Tested, c.People_Hospitalized, c.Mortality_Rate, c.uid, c.Testing_Rate, c.Hospitalization_Rate, c.Created_At, l."Province_State" FROM covid_usa.cases_by_date c
LEFT JOIN covid_usa.locations l on l.uid = c.uid
WHERE l."Province_State" = $1 AND c."Created_At" = $2
`

type GetCasesByStateOnDateParams struct {
	ProvinceState sql.NullString
	CreatedAt     time.Time
}

type GetCasesByStateOnDateRow struct {
	LastUpdate          sql.NullTime
	Confirmed           sql.NullInt32
	Deaths              sql.NullInt32
	Recovered           sql.NullFloat64
	Active              sql.NullFloat64
	IncidentRate        sql.NullString
	PeopleTested        sql.NullFloat64
	PeopleHospitalized  sql.NullFloat64
	MortalityRate       sql.NullString
	Uid                 int64
	TestingRate         sql.NullString
	HospitalizationRate sql.NullString
	CreatedAt           time.Time
	ProvinceState       sql.NullString
}

func (q *Queries) GetCasesByStateOnDate(ctx context.Context, arg GetCasesByStateOnDateParams) ([]GetCasesByStateOnDateRow, error) {
	rows, err := q.db.QueryContext(ctx, getCasesByStateOnDate, arg.ProvinceState, arg.CreatedAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCasesByStateOnDateRow
	for rows.Next() {
		var i GetCasesByStateOnDateRow
		if err := rows.Scan(
			&i.LastUpdate,
			&i.Confirmed,
			&i.Deaths,
			&i.Recovered,
			&i.Active,
			&i.IncidentRate,
			&i.PeopleTested,
			&i.PeopleHospitalized,
			&i.MortalityRate,
			&i.Uid,
			&i.TestingRate,
			&i.HospitalizationRate,
			&i.CreatedAt,
			&i.ProvinceState,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCasesByStateSinceDate = `-- name: GetCasesByStateSinceDate :many
SELECT c.Last_Update, c.Confirmed, c.Deaths, c.Recovered, c.Active, c.Incident_Rate, c.People_Tested, c.People_Hospitalized, c.Mortality_Rate, c.uid, c.Testing_Rate, c.Hospitalization_Rate, c.Created_At, l."Province_State" FROM covid_usa.cases_by_date c
LEFT JOIN covid_usa.locations l on l.uid = c.uid
WHERE l."Province_State" = $1 AND c."Created_At" >= $2
`

type GetCasesByStateSinceDateParams struct {
	ProvinceState sql.NullString
	CreatedAt     time.Time
}

type GetCasesByStateSinceDateRow struct {
	LastUpdate          sql.NullTime
	Confirmed           sql.NullInt32
	Deaths              sql.NullInt32
	Recovered           sql.NullFloat64
	Active              sql.NullFloat64
	IncidentRate        sql.NullString
	PeopleTested        sql.NullFloat64
	PeopleHospitalized  sql.NullFloat64
	MortalityRate       sql.NullString
	Uid                 int64
	TestingRate         sql.NullString
	HospitalizationRate sql.NullString
	CreatedAt           time.Time
	ProvinceState       sql.NullString
}

func (q *Queries) GetCasesByStateSinceDate(ctx context.Context, arg GetCasesByStateSinceDateParams) ([]GetCasesByStateSinceDateRow, error) {
	rows, err := q.db.QueryContext(ctx, getCasesByStateSinceDate, arg.ProvinceState, arg.CreatedAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCasesByStateSinceDateRow
	for rows.Next() {
		var i GetCasesByStateSinceDateRow
		if err := rows.Scan(
			&i.LastUpdate,
			&i.Confirmed,
			&i.Deaths,
			&i.Recovered,
			&i.Active,
			&i.IncidentRate,
			&i.PeopleTested,
			&i.PeopleHospitalized,
			&i.MortalityRate,
			&i.Uid,
			&i.TestingRate,
			&i.HospitalizationRate,
			&i.CreatedAt,
			&i.ProvinceState,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCasesByUID = `-- name: GetCasesByUID :many
SELECT c.Last_Update, c.Confirmed, c.Deaths, c.Recovered, c.Active, c.Incident_Rate, c.People_Tested, c.People_Hospitalized, c.Mortality_Rate, c.uid, c.Testing_Rate, c.Hospitalization_Rate, c.Created_At, l."Province_State" FROM covid_usa.cases_by_date c
LEFT JOIN covid_usa.locations l on c.uid = l.uid
WHERE c.uid = $1
`

type GetCasesByUIDRow struct {
	LastUpdate          sql.NullTime
	Confirmed           sql.NullInt32
	Deaths              sql.NullInt32
	Recovered           sql.NullFloat64
	Active              sql.NullFloat64
	IncidentRate        sql.NullString
	PeopleTested        sql.NullFloat64
	PeopleHospitalized  sql.NullFloat64
	MortalityRate       sql.NullString
	Uid                 int64
	TestingRate         sql.NullString
	HospitalizationRate sql.NullString
	CreatedAt           time.Time
	ProvinceState       sql.NullString
}

func (q *Queries) GetCasesByUID(ctx context.Context, uid int64) ([]GetCasesByUIDRow, error) {
	rows, err := q.db.QueryContext(ctx, getCasesByUID, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCasesByUIDRow
	for rows.Next() {
		var i GetCasesByUIDRow
		if err := rows.Scan(
			&i.LastUpdate,
			&i.Confirmed,
			&i.Deaths,
			&i.Recovered,
			&i.Active,
			&i.IncidentRate,
			&i.PeopleTested,
			&i.PeopleHospitalized,
			&i.MortalityRate,
			&i.Uid,
			&i.TestingRate,
			&i.HospitalizationRate,
			&i.CreatedAt,
			&i.ProvinceState,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCasesSinceDate = `-- name: GetCasesSinceDate :many
SELECT c.Last_Update, c.Confirmed, c.Deaths, c.Recovered, c.Active, c.Incident_Rate, c.People_Tested, c.People_Hospitalized, c.Mortality_Rate, c.uid, c.Testing_Rate, c.Hospitalization_Rate, c.Created_At, l."Province_State" FROM covid_usa.cases_by_date c
LEFT JOIN covid_usa.locations l on l.uid = c.uid
WHERE c."Created_At" >= $1
`

type GetCasesSinceDateRow struct {
	LastUpdate          sql.NullTime
	Confirmed           sql.NullInt32
	Deaths              sql.NullInt32
	Recovered           sql.NullFloat64
	Active              sql.NullFloat64
	IncidentRate        sql.NullString
	PeopleTested        sql.NullFloat64
	PeopleHospitalized  sql.NullFloat64
	MortalityRate       sql.NullString
	Uid                 int64
	TestingRate         sql.NullString
	HospitalizationRate sql.NullString
	CreatedAt           time.Time
	ProvinceState       sql.NullString
}

func (q *Queries) GetCasesSinceDate(ctx context.Context, createdAt time.Time) ([]GetCasesSinceDateRow, error) {
	rows, err := q.db.QueryContext(ctx, getCasesSinceDate, createdAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCasesSinceDateRow
	for rows.Next() {
		var i GetCasesSinceDateRow
		if err := rows.Scan(
			&i.LastUpdate,
			&i.Confirmed,
			&i.Deaths,
			&i.Recovered,
			&i.Active,
			&i.IncidentRate,
			&i.PeopleTested,
			&i.PeopleHospitalized,
			&i.MortalityRate,
			&i.Uid,
			&i.TestingRate,
			&i.HospitalizationRate,
			&i.CreatedAt,
			&i.ProvinceState,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
