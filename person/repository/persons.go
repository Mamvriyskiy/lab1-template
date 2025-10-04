package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/Mamvriyskiy/lab1-template/person/model"
)

type PersonsPostgres struct {
	db *sqlx.DB
}

func NewPersonsPostgres(db *sqlx.DB) *PersonsPostgres {
	return &PersonsPostgres{db: db}
}

func (r *PersonsPostgres) GetInfoPerson(persinID int) (model.Person, error) {
	var person model.Person
	queryGetPersons := `select * from person where personid = $1`
	err := r.db.Get(&person, queryGetPersons, persinID)
	if err != nil {
		return model.Person{}, err
	}

	return person, nil
}

func (r *PersonsPostgres) GetInfoPersons() ([]model.Person, error) {
	var persons []model.Person
	queryGetPersons := `select * from person`
	err := r.db.Select(&persons, queryGetPersons)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func (r *PersonsPostgres) CreateNewRecordPerson(person model.Person) (model.Person, error) {
	var newPerson model.Person
	queryCreatePerson := `insert into person (Name, Age, Address, Work) values($1, $2, $3, $4) returning *`
	err := r.db.Get(&newPerson, queryCreatePerson, person.Name, person.Age, person.Address, person.Work)
	if err != nil {
		return model.Person{}, err
	}

	return newPerson, nil
}

func (r *PersonsPostgres) UpdateRecordPerson(person model.Person) (model.Person, error) {
	queryUpdatePerson := `
        UPDATE person 
        SET name = $1, age = $2, address = $3, work = $4 
        WHERE personid = $5
        RETURNING personid, name, age, address, work`
    
    var updatedPerson model.Person
    err := r.db.QueryRow(
        queryUpdatePerson, 
        person.Name, 
        person.Age, 
        person.Address, 
        person.Work, 
        person.PersonID,
    ).Scan(
        &updatedPerson.PersonID,
        &updatedPerson.Name, 
        &updatedPerson.Age,
        &updatedPerson.Address,
        &updatedPerson.Work,
    )
    
    if err != nil {
        return model.Person{}, err
    }

    return updatedPerson, nil
}

func (r *PersonsPostgres) DeleteRecordPerson(personID int) error {

	queryDeletePerson := `delete from person where personid = $1`
	_, err := r.db.Exec(queryDeletePerson, personID)

	return err
}
