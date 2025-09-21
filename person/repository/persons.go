package repository

type PersonsPostgres struct {
	db *sqlx.DB
}

func NewPersonsPostgres(db *sqlx.DB) *PersonsPostgres {
	return &PersonsPostgres{db: db}
}

