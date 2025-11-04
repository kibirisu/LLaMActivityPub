package data

import (
	"context"
	"encoding/json"

	"borg/pkg/db"
)

type DataStore interface {
	UserRepository() UserRepository
	PostRepository() PostRepository
	Opts() json.Options
}

type Repository[T, C, U any] interface {
	Create(context.Context, C) error
	GetByID(context.Context, int32) (T, error)
	GetAll(context.Context) ([]T, error)
	Update(context.Context, U) error
	Delete(context.Context, int32) error
}

type dataStore struct {
	users UserRepository
	posts PostRepository
	opts  json.Options
}

func NewDataStore(ctx context.Context, url string) (DataStore, error) {
	q, err := db.GetDB(ctx, url)
	if err != nil {
		return nil, err
	}
	ds := &dataStore{opts: getOptions()}
	ds.users = newUserRepository(q)
	ds.posts = newPostRepository(q)
	return ds, nil
}

func (ds *dataStore) UserRepository() UserRepository {
	return ds.users
}

func (ds *dataStore) PostRepository() PostRepository {
	return ds.posts
}

func (ds *dataStore) Opts() json.Options {
	return ds.opts
}
