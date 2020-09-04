package service

import (
	"github.com/7dc/gin-structure/app/entity"
)

type PersonService interface {
	Save(entity.Person) entity.Person
	FindAll() []entity.Person
}

type personService struct {
	people []entity.Person
}

func New() PersonService {
	return &personService{}
}

func (service *personService) Save(person entity.Person) entity.Person {
	service.people = append(service.people, person)
	return person
}

func (service *personService) FindAll() []entity.Person {
	return service.people
}
