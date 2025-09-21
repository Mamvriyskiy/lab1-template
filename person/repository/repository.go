package repository

import (
	"github.com/jmoiron/sqlx"
)

type RepoPersons interface {
	
}

type Repository struct {
	RepoPersons
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		RepoPersons: NewPersonsPostgres(db),
	}
}