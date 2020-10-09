package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq" // we dont need methods of this mod
)

// Store ...
type Store struct {
	db             *sql.DB
	UserRepository *UserRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User ...
func (s *Store) User() *UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}

	s.UserRepository = &UserRepository{
		store: s,
	}

	return s.UserRepository
}
