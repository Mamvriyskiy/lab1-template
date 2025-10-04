package repository

import (
	"errors"
	"regexp"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Mamvriyskiy/lab1-template/person/model"
	"github.com/stretchr/testify/assert"
)

func setupMockDB(t *testing.T) (*PersonsPostgres, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := NewPersonsPostgres(sqlxDB)

	// Возвращаем репозиторий, mock и функцию закрытия
	return repo, mock, func() { db.Close() }
}

// ================= GetInfoPerson =================
func TestGetInfoPerson_Success(t *testing.T) {
	repo, mock, close := setupMockDB(t)
	defer close()

	rows := sqlmock.NewRows([]string{"personid", "name", "age", "address", "work"}).
		AddRow(1, "John", 30, "NY", "Engineer")

	mock.ExpectQuery(regexp.QuoteMeta("select * from person where personid = $1")).
		WithArgs(1).
		WillReturnRows(rows)

	person, err := repo.GetInfoPerson(1)
	assert.NoError(t, err)
	assert.Equal(t, "John", person.Name)
}

func TestGetInfoPerson_Error(t *testing.T) {
	repo, mock, close := setupMockDB(t)
	defer close()

	mock.ExpectQuery(regexp.QuoteMeta("select * from person where personid = $1")).
		WithArgs(1).
		WillReturnError(errors.New("db error"))

	_, err := repo.GetInfoPerson(1)
	assert.Error(t, err)
}

// ================= GetInfoPersons =================
func TestGetInfoPersons_Success(t *testing.T) {
	repo, mock, close := setupMockDB(t)
	defer close()

	rows := sqlmock.NewRows([]string{"personid", "name", "age", "address", "work"}).
		AddRow(1, "John", 30, "NY", "Engineer").
		AddRow(2, "Alice", 25, "LA", "Designer")

	mock.ExpectQuery(regexp.QuoteMeta("select * from person")).
		WillReturnRows(rows)

	persons, err := repo.GetInfoPersons()
	assert.NoError(t, err)
	assert.Len(t, persons, 2)
}

func TestGetInfoPersons_Error(t *testing.T) {
	repo, mock, close := setupMockDB(t)
	defer close()

	mock.ExpectQuery(regexp.QuoteMeta("select * from person")).
		WillReturnError(errors.New("db error"))

	_, err := repo.GetInfoPersons()
	assert.Error(t, err)
}

// ================= CreateNewRecordPerson =================
func TestCreateNewRecordPerson_Success(t *testing.T) {
	repo, mock, close := setupMockDB(t)
	defer close()

	newPerson := model.Person{Name: "Bob", Age: 28, Address: "Chicago", Work: "Chef"}

	mock.ExpectQuery(regexp.QuoteMeta(
		"insert into person (Name, Age, Address, Work) values($1, $2, $3, $4) returning *")).
		WithArgs(newPerson.Name, newPerson.Age, newPerson.Address, newPerson.Work).
		WillReturnRows(sqlmock.NewRows([]string{"personid", "name", "age", "address", "work"}).
			AddRow(1, "Bob", 28, "Chicago", "Chef"))

	created, err := repo.CreateNewRecordPerson(newPerson)
	assert.NoError(t, err)
	assert.Equal(t, "Bob", created.Name)
}

func TestCreateNewRecordPerson_Error(t *testing.T) {
	repo, mock, close := setupMockDB(t)
	defer close()

	newPerson := model.Person{Name: "Bob", Age: 28, Address: "Chicago", Work: "Chef"}

	mock.ExpectQuery(regexp.QuoteMeta(
		"insert into person (Name, Age, Address, Work) values($1, $2, $3, $4) returning *")).
		WillReturnError(errors.New("insert error"))

	_, err := repo.CreateNewRecordPerson(newPerson)
	assert.Error(t, err)
}

// ================= UpdateRecordPerson =================
func TestUpdateRecordPerson_Success(t *testing.T) {
	repo, mock, close := setupMockDB(t)
	defer close()

	person := model.Person{PersonID: 1, Name: "Bob", Age: 29, Address: "Chicago", Work: "Chef"}

	mock.ExpectQuery(regexp.QuoteMeta(`
			UPDATE person
			SET 
				name    = COALESCE(NULLIF($1, ''), name),
				age     = COALESCE(NULLIF($2, 0), age),
				address = COALESCE(NULLIF($3, ''), address),
				work    = COALESCE(NULLIF($4, ''), work)
			WHERE personid = $5
			RETURNING personid, name, age, address, work;
		`)).
		WithArgs(person.Name, person.Age, person.Address, person.Work, person.PersonID).
		WillReturnRows(sqlmock.NewRows([]string{"personid", "name", "age", "address", "work"}).
			AddRow(1, "Bob", 29, "Chicago", "Chef"))

	updated, err := repo.UpdateRecordPerson(person)
	assert.NoError(t, err)
	assert.Equal(t, 29, updated.Age)
}

func TestUpdateRecordPerson_Error(t *testing.T) {
	repo, mock, close := setupMockDB(t)
	defer close()

	person := model.Person{PersonID: 1}

	mock.ExpectQuery(regexp.QuoteMeta(`
        UPDATE person 
        SET name = $1, age = $2, address = $3, work = $4 
        WHERE personid= $5
        RETURNING id, name, age, address, work`)).
		WillReturnError(errors.New("update error"))

	_, err := repo.UpdateRecordPerson(person)
	assert.Error(t, err)
}

// ================= DeleteRecordPerson =================
func TestDeleteRecordPerson_Success(t *testing.T) {
	repo, mock, close := setupMockDB(t)
	defer close()

	mock.ExpectExec(regexp.QuoteMeta("delete from person where personid = $1")).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.DeleteRecordPerson(1)
	assert.NoError(t, err)
}

func TestDeleteRecordPerson_Error(t *testing.T) {
	repo, mock, close := setupMockDB(t)
	defer close()

	mock.ExpectExec(regexp.QuoteMeta("delete from person where personid = $1")).
		WithArgs(1).
		WillReturnError(errors.New("delete error"))

	err := repo.DeleteRecordPerson(1)
	assert.Error(t, err)
}
