package controller

import (
	"github.com/7dc/gin-proper-structure/app/entity"
	"github.com/7dc/gin-proper-structure/app/service"
	"github.com/gin-gonic/gin"
)

type PersonController interface {
	FindAll() []entity.Person
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.PersonService
}

func New(service service.PersonService) PersonController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Person {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var team entity.Person
	err := ctx.ShouldBindJSON(&team)
	if err != nil {
		return err
	}
	c.service.Save(team)
	return nil
}
