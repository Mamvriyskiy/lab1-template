package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/Mamvriyskiy/lab1-template/person/model"
)

type RepoPersons interface {
	GetInfoPerson(persinID int) (model.Person, error)
	GetInfoPersons() ([]model.Person, error)
	CreateNewRecordPerson(person model.Person) (model.Person, error)
	UpdateRecordPerson(person model.Person) (model.Person, error)
	DeleteRecordPerson(personID int) error
}

type Repository struct {
	RepoPersons
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		RepoPersons: NewPersonsPostgres(db),
	}
}
