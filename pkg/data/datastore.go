package data

import (
	"context"
	"encoding/json"

	"borg/pkg/db"
)

type DataStore interface {
	UserRepository() UserRepository
	PostRepository() PostRepository
	LikeRepository() LikeRepository
	ShareRepository() ShareRepository
	Opts() json.Options
}

type Repository[T, C, U any] interface {
	Create(context.Context, C) error
	GetByID(context.Context, int32) (T, error)
	Update(context.Context, U) error
	Delete(context.Context, int32) error
}

type SearchableByUserID[T, C, U any] interface {
	Repository[T, C, U]
	IdentifiedByUser[T]
}

type SearchableByPostID[T, C, U any] interface {
	Repository[T, C, U]
	IdentifiedByPost[T]
}

type IdentifiedByUser[T any] interface {
	GetByUserID(context.Context, int32) ([]T, error)
}

type IdentifiedByPost[T any] interface {
	GetByPostID(context.Context, int32) ([]T, error)
}

type dataStore struct {
	users  UserRepository
	posts  PostRepository
	likes  LikeRepository
	shares ShareRepository
	opts   json.Options
}

func NewDataStore(ctx context.Context, url string) (DataStore, error) {
	q, err := db.GetDB(ctx, url)
	if err != nil {
		return nil, err
	}
	ds := &dataStore{opts: getOptions()}
	ds.users = newUserRepository(q)
	ds.posts = newPostRepository(q)
	ds.likes = newLikeRepository(q)
	ds.shares = newShareRepository(q)
	return ds, nil
}

func (ds *dataStore) UserRepository() UserRepository {
	return ds.users
}

func (ds *dataStore) PostRepository() PostRepository {
	return ds.posts
}

func (ds *dataStore) LikeRepository() LikeRepository {
	return ds.likes
}

func (ds *dataStore) ShareRepository() ShareRepository {
	return ds.shares
}

func (ds *dataStore) Opts() json.Options {
	return ds.opts
}
