package dataloaders

import (
	"context"
	"covidProject/db"
	"time"
)

type contextKey string

const key = contextKey("dataloaders")

// Loaders holds references to the individual dataloaders.
type Loaders struct {
	LocationByCaseID *LocationLoader
}

func newLoaders(ctx context.Context, repo db.Repository) *Loaders {
	return &Loaders{
		// individual loaders will be initialized here
		LocationByCaseID: newLocationByCaseID(ctx, repo),
	}
}

// Retriever retrieves dataloaders from the request context.
type Retriever interface {
	Retrieve(context.Context) *Loaders
}

type retriever struct {
	key contextKey
}

func (r *retriever) Retrieve(ctx context.Context) *Loaders {
	return ctx.Value(r.key).(*Loaders)
}

// NewRetriever instantiates a new implementation of Retriever.
func NewRetriever() Retriever {
	return &retriever{key: key}
}

func newLocationByCaseID(ctx context.Context, repo db.Repository) *LocationLoader {
	return NewLocationLoader(LocationLoaderConfig{
		MaxBatch: 100,
		Wait: 5 * time.Millisecond,
		Fetch: func(CaseIDs []int64) ([]*db.CovidUsaLocation, []error) {
			// db query
			res, err := repo.ListLocationsByCaseIDs(ctx, CaseIDs)
			if err != nil {
				return nil, []error{err}
			}
			// map
			groupByCaseID := make(map[int64]*db.CovidUsaLocation, len(CaseIDs))
			for _, r := range res {
				groupByCaseID[r.Uid] = &db.CovidUsaLocation{
					Uid: r.Uid,
					Iso2: r.Iso2,
					Iso3: r.Iso3,
					Code3: r.Code3,
					Fips: r.Fips,
					Admin2: r.Admin2,
					ProvinceState: r.ProvinceState,
					CountryRegion: r.CountryRegion,
					Lat: r.Lat,
					Long: r.Long,
					CombinedKey: r.CombinedKey,
					Population: r.Population,
				}
			}
			// order
			result := make([]*db.CovidUsaLocation, len(CaseIDs))
			for i, caseID := range CaseIDs {
				result[i] = groupByCaseID[caseID]
			}
			return result, nil
		},
	})
}

//go:generate go run github.com/vektah/dataloaden LocationLoader int64 *covidProject/db.CovidUsaLocation
