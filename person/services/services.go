package services

import (
	"github.com/Mamvriyskiy/lab1-template/person/repository"
)

type Persons interface {

}

type Services struct {
	Persons
}

func NewServices(repo * repository.Repository) *Services {
	return &Services{
		Persons: NewPersonsService(repo.PersonsRepo),
	}
}

