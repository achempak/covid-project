package graphql

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"covidProject/dataloaders"
	"covidProject/db"
	"covidProject/graphql/generated"
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

const dateError = "could not properly parse date. Should be in format \"YYYY-MM-DD\""
const caseError = "could not fetch case"
const locationError = "could not fetch location"
func fixSqlType(n interface{}) interface{} {
	switch v := n.(type) {
	case sql.NullInt64:
		if v.Valid {
			return &v.Int64
		}
		return (*int64)(nil)
	case sql.NullInt32:
		if v.Valid {
			n64 := int64(v.Int32)
			return &n64
		}
		return (*int64)(nil)
	case sql.NullFloat64:
		if v.Valid {
			return &v.Float64
		}
		return (*float64)(nil)
	case sql.NullString:
		if v.Valid {
			return &v.String
		}
		return (*string)(nil)
	case sql.NullBool:
		if v.Valid {
			return &v.Bool
		}
		return (*bool)(nil)
	case sql.NullTime:
		if v.Valid {
			timeString := v.Time.String()
			return &timeString
		}
		return (*string)(nil)
	case time.Time:
		return v.String()
	case int64, string, bool, float64:
		return v
	}
	return nil
}

func stringToFloat(s *string) *float64 {
	f, err := strconv.ParseFloat(*s, 64)
	if err != nil {
		return (*float64)(nil)
	}
	return &f
}

type Resolver struct{
	Repository db.Repository
	DataLoaders dataloaders.Retriever
}

func (r *caseResolver) ID(ctx context.Context, obj *db.CovidUsaCasesByDate) (int64, error) {
	return fixSqlType(obj.Uid).(int64), nil
}

func (r *caseResolver) Location(ctx context.Context, obj *db.CovidUsaCasesByDate) (*db.CovidUsaLocation, error) {
	//loc, err := r.Repository.GetLocationByUID(ctx, obj.Uid)
	//if err != nil {
	//	return nil, err
	//}
	//return &loc, nil
	return r.DataLoaders.Retrieve(ctx).LocationByCaseID.Load(obj.Uid)
}

func (r *caseResolver) CreatedAt(ctx context.Context, obj *db.CovidUsaCasesByDate) (string, error) {
	return fixSqlType(obj.CreatedAt).(string), nil
}

func (r *caseResolver) LastUpdate(ctx context.Context, obj *db.CovidUsaCasesByDate) (*string, error) {
	return fixSqlType(obj.LastUpdate).(*string), nil
}

func (r *caseResolver) Confirmed(ctx context.Context, obj *db.CovidUsaCasesByDate) (*int64, error) {
	return fixSqlType(obj.Confirmed).(*int64), nil
}

func (r *caseResolver) Deaths(ctx context.Context, obj *db.CovidUsaCasesByDate) (*int64, error) {
	return fixSqlType(obj.Deaths).(*int64), nil
}

func (r *caseResolver) Recovered(ctx context.Context, obj *db.CovidUsaCasesByDate) (*int64, error) {
	n := int64(*fixSqlType(obj.Recovered).(*float64))
	return &n, nil
}

func (r *caseResolver) Active(ctx context.Context, obj *db.CovidUsaCasesByDate) (*int64, error) {
	n := int64(*fixSqlType(obj.Active).(*float64))
	return &n, nil
}

func (r *caseResolver) IncidentRate(ctx context.Context, obj *db.CovidUsaCasesByDate) (*float64, error) {
	return stringToFloat(fixSqlType(obj.IncidentRate).(*string)), nil
}

func (r *caseResolver) PeopleTested(ctx context.Context, obj *db.CovidUsaCasesByDate) (*float64, error) {
	return fixSqlType(obj.PeopleTested).(*float64), nil
}

func (r *caseResolver) PeopleHospitalized(ctx context.Context, obj *db.CovidUsaCasesByDate) (*float64, error) {
	return fixSqlType(obj.PeopleHospitalized).(*float64), nil
}

func (r *caseResolver) MortalityRate(ctx context.Context, obj *db.CovidUsaCasesByDate) (*float64, error) {
	return stringToFloat(fixSqlType(obj.MortalityRate).(*string)), nil
}

func (r *caseResolver) TestingRate(ctx context.Context, obj *db.CovidUsaCasesByDate) (*float64, error) {
	return stringToFloat(fixSqlType(obj.TestingRate).(*string)), nil
}

func (r *caseResolver) HospitalizationRate(ctx context.Context, obj *db.CovidUsaCasesByDate) (*float64, error) {
	return stringToFloat(fixSqlType(obj.HospitalizationRate).(*string)), nil
}

func (r *locationResolver) ID(ctx context.Context, obj *db.CovidUsaLocation) (int64, error) {
	return fixSqlType(obj.Uid).(int64), nil
}

func (r *locationResolver) Iso2(ctx context.Context, obj *db.CovidUsaLocation) (*string, error) {
	return fixSqlType(obj.Iso2).(*string), nil
}

func (r *locationResolver) Iso3(ctx context.Context, obj *db.CovidUsaLocation) (*string, error) {
	return fixSqlType(obj.Iso3).(*string), nil
}

func (r *locationResolver) Code3(ctx context.Context, obj *db.CovidUsaLocation) (*int64, error) {
	return fixSqlType(obj.Code3).(*int64), nil
}

func (r *locationResolver) Fips(ctx context.Context, obj *db.CovidUsaLocation) (*int64, error) {
	//n := int64(*(fixSqlType(obj.Fips).(*float64)))
	if obj.Fips.Valid {
		num := int64(obj.Fips.Float64)
		return &num, nil
	}
	return nil, nil
}

func (r *locationResolver) Admin2(ctx context.Context, obj *db.CovidUsaLocation) (*string, error) {
	return fixSqlType(obj.Admin2).(*string), nil
}

func (r *locationResolver) ProvinceState(ctx context.Context, obj *db.CovidUsaLocation) (*string, error) {
	return fixSqlType(obj.ProvinceState).(*string), nil
}

func (r *locationResolver) CountryRegion(ctx context.Context, obj *db.CovidUsaLocation) (*string, error) {
	return fixSqlType(obj.CountryRegion).(*string), nil
}

func (r *locationResolver) Lat(ctx context.Context, obj *db.CovidUsaLocation) (*float64, error) {
	return fixSqlType(obj.Lat).(*float64), nil
}

func (r *locationResolver) Long(ctx context.Context, obj *db.CovidUsaLocation) (*float64, error) {
	return fixSqlType(obj.Long).(*float64), nil
}

func (r *locationResolver) CombinedKey(ctx context.Context, obj *db.CovidUsaLocation) (*string, error) {
	return fixSqlType(obj.CombinedKey).(*string), nil
}

func (r *locationResolver) Population(ctx context.Context, obj *db.CovidUsaLocation) (*int64, error) {
	return fixSqlType(obj.Population).(*int64), nil
}

func (r *queryResolver) Case(ctx context.Context, id int64, createdAt string) (*db.CovidUsaCasesByDate, error) {
	t, err := time.Parse("2006-01-02", createdAt)
	if err != nil {
		return nil, fmt.Errorf(dateError)
	}
	c, err := r.Repository.GetCaseByUIDOnDate(ctx, db.GetCaseByUIDOnDateParams{
		Uid:       id,
		CreatedAt: t,
	})
	return &c, err
}

func (r *queryResolver) Cases(ctx context.Context, id *int64, createdAt *string) ([]db.CovidUsaCasesByDate, error) {
	if id == nil && createdAt == nil {
		return r.Repository.GetCases(ctx)
	}
	if id != nil && createdAt != nil {
		c, err := r.Case(ctx, *id, *createdAt)
		if err != nil {
			return nil, err
		}
		return []db.CovidUsaCasesByDate{*c}, nil
	}
	if createdAt == nil {
		return r.Repository.GetCasesByUID(ctx, *id)
	}

	// if id == nil
	t, err := time.Parse("2006-01-02", *createdAt)
	if err != nil {
		return nil, fmt.Errorf(dateError)
	}
	return r.Repository.GetCasesByDate(ctx, t)
}

func (r *queryResolver) CasesSince(ctx context.Context, id *int64, createdSince *string) ([]db.CovidUsaCasesByDate, error) {
	if id == nil && createdSince == nil {
		return nil, nil
	}
	if createdSince == nil {
		return r.Cases(ctx, id, nil)
	}

	t, err := time.Parse("2006-01-02", *createdSince)
	if err != nil {
		return nil, fmt.Errorf(dateError)
	}

	if id == nil {
		return r.Repository.GetCasesSinceDate(ctx, t)
	}
	return r.Repository.GetCasesByUIDSinceDate(ctx, db.GetCasesByUIDSinceDateParams{
		Uid:       *id,
		CreatedAt: t,
	})
}

func (r *queryResolver) Location(ctx context.Context, id int64) (*db.CovidUsaLocation, error) {
	l, err := r.Repository.GetLocationByUID(ctx, id)
	return &l, err
}

func (r *queryResolver) Locations(ctx context.Context) ([]db.CovidUsaLocation, error) {
	return r.Repository.GetLocations(ctx)
}

// Case returns generated.CaseResolver implementation.
func (r *Resolver) Case() generated.CaseResolver { return &caseResolver{r} }

// Location returns generated.LocationResolver implementation.
func (r *Resolver) Location() generated.LocationResolver { return &locationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type caseResolver struct{ *Resolver }
type locationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
