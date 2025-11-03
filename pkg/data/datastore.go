package data

import (
	"context"
	"encoding/json"

	"borg/pkg/db"
)

type DataStore interface {
	UserRepository
	GetOpts() json.Options
}

type dataStore struct {
	UserRepository
	opts json.Options
}

func NewDataStore(ctx context.Context, url string) (DataStore, error) {
	q, err := db.GetDB(ctx, url)
	if err != nil {
		return nil, err
	}
	ds := &dataStore{opts: getOptions()}
	ds.UserRepository = newUserRepository(q)
	return ds, nil
}

func (ds *dataStore) GetOpts() json.Options {
	return ds.opts
}
