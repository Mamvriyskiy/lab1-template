package services

import "github.com/!mamvriyskiy/database_course/main/pkg/repository"

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

