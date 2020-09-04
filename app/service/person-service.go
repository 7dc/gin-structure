package service

import (
	"github.com/7dc/gin-structure/app/entity"
	"github.com/7dc/gin-structure/app/repository"
)

type PersonService interface {
	Save(entity.Person) error
	FindAll() []entity.Person
}

type personService struct {
	people     []entity.Person
	repository repository.PersonRepository
}

func New(repository repository.PersonRepository) PersonService {
	return &personService{
		repository: repository,
	}
}

func (service *personService) Save(person entity.Person) error {
	err := service.repository.Save(person)
	return err
}

func (service *personService) FindAll() []entity.Person {
	return service.repository.FindAll()
}
