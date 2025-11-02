package data

import (
	"context"

	"borg/pkg/db"
)

type DataStore interface {
	UserRepository
}

type dataStore struct {
	UserRepository
}

func NewDataStore(ctx context.Context, url string) (DataStore, error) {
	q, err := db.GetDB(ctx, url)
	if err != nil {
		return nil, err
	}
	users := newUserRepository(q)
	return &dataStore{users}, nil
}
