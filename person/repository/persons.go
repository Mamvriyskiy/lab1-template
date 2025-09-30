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

func (r *Repository) GetInfoPerson(persinID int) (model.Person, error) {
	var person model.Person
	queryGetPersons := `select * from person where id = $1`
	err := r.db.Select(&person, queryGetPersons, persinID)
	if err != nil {
		return model.Person{}, err
	}

	return person, nil
}

func (r *Repository) GetInfoPersons() ([]model.Person, error) {
	var persons []model.Person
	queryGetPersons := `select * from person where id = $1`
	err := r.db.Select(&persons, queryGetPersons, persinID)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func (r *Repository) CreateNewRecordPerson(person model.Person) (model.Person, error) {
	var newPerson model.Person
	queryCreatePerson := `insert into person (Name, Age, Address, Work) values($1, $2, $3, $4) returning *`
	row := r.db.QueryRow(queryCreatePerson, person.Name, person.Age, person.Address, person.Work)
	err := row.Scan(&newPerson)
	if err != nil {
		return model.Person{}, err
	}

	return newPerson, nil
}

func (r *Repository) UpdateRecordPerson(person model.Person) (model.Person, error) {
	queryUpdatePerson := `update person set name = $1, age = $2, address = $3, work = $4 where id = $5`
	updatePerson, err := r.db.Exec(queryUpdatePerson, person.Name, person.Work, person.Address)
	if err != nil {
		retirn model.Person{}, err
	}

	return updatePerson, nil
}

func (r *Repository) DeleteRecordPerson(personID int) error {

	queryDeletePerson := `delete from person where personid = $1`
	_, err := r.db.Exec(queryDeleteAccessHomeID, personID)

	return err
}
