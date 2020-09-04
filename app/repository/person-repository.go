package repository

import (
	"context"
	"fmt"

	"github.com/7dc/gin-structure/app/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PersonRepository interface {
	Save(person entity.Person) error
	FindAll() []entity.Person
}

type personRepository struct {
	collection *mongo.Collection
}

func New(c *mongo.Database) PersonRepository {
	return &personRepository{
		collection: c.Collection("person"),
	}
}

func (repository *personRepository) Save(person entity.Person) error {
	_, err := repository.collection.InsertOne(context.TODO(), person)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (repository *personRepository) FindAll() []entity.Person {
	cursor, err := repository.collection.Find(context.TODO(), bson.M{})
	var people []entity.Person
	if err != nil {
		fmt.Println(err.Error())
	}
	if err = cursor.All(context.TODO(), &people); err != nil {
		panic(err)
	}
	return people
}
