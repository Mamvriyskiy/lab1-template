package model

type Person struct {
	PersonID int    `db:"personid"`
	Name     string `db:"name" json:"name"`
	Age      int    `db:"age" json:"age"`
	Address  string `db:"address" json:"address"`
	Work     string `db:"work" json:"work"`
}
