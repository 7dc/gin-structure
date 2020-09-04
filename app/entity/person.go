package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Person struct {
	_id       primitive.ObjectID `bson:"_id,omitempty"`
	Firstname string             `bson:"firstname,omitempty"`
	Lastname  string             `bson:"lastname,omitempty"`
	Age       int                `bson:"age,omitempty"`
}
