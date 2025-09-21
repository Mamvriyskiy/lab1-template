package repository

import (
	"github.com/Mamvriyskiy/lab1-template/person/model"
)

type PersonsPostgres struct {
	db *sqlx.DB
}

func NewPersonsPostgres(db *sqlx.DB) *PersonsPostgres {
	return &PersonsPostgres{db: db}
}

func (r *Repository) GetInfoPerson() (model.Person, error) {
	return model.Person{}, nil
}

func (r *Repository) GetInfoPersons() ([]model.Person, error) {
	return []model.Person{}, nil
}

func (r *Repository) CreateNewRecordPerson() (model.Person, error) {
	return model.Person{}, nil
}

func (r *Repository) UpdateRecordPerson() (model.Person, error) {
	return model.Person{}, nil
}

func (r *Repository) DeleteRecordPerson() error {
	return model.Person{}, nil
}
