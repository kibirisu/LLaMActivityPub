package datastore

import "borg/pkg/db"

type DataStore interface {
	UserRepository
}

type dataStore struct {
	UserRepository
}

func NewDataStore(q *db.Queries) DataStore {
	users := newUserRepository(q)
	return &dataStore{users}
}
