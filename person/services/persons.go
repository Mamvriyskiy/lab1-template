package services

import (
	"github.com/Mamvriyskiy/lab1-template/person/model"
	"github.com/Mamvriyskiy/lab1-template/person/repository"
	"github.com/golang/mock/mockgen/model"
)

type PersonsService struct {
	repo repository.Repository
}

func NewPersonsService(repo repository.Repository) *PersonsService {
	return &PersonsService{repo: repo}
}

func (s *Services) GetInfoPerson() (model.Person, error) {
	return s.repo.GetInfoPerson()
}

func (s *Services) GetInfoPersons() ([]model.Person, error) {
	return s.repo.GetInfoPersons(), nil
}

func (s *Services) CreateNewRecordPerson() (model.Person, error) {
	return s.repo.CreateNewRecordPerson(), nil
}

func (s *Services) UpdateRecordPerson() (model.Person, error) {
	return s.repo.UpdateRecordPerson(), nil
}

func (s *Services) DeleteRecordPerson() error {
	return s.repo.DeleteRecordPerson(), nil
}
