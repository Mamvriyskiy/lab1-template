package services

import (

)

type PersonsService struct {
	repo repository.Repository
}

func NewPersonsService(repo repository.Repository) *PersonsService {
	return &PersonsService{repo: repo}
}

