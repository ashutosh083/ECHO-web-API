package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Student struct {
	Id      primitive.ObjectID `json:"id,omitempty"`
	Name    string             `json:"name,omitempty" validate:"required"`
	Rollno  string             `json:"rollno,omitempty" validate:"required"`
	Address string             `json:"address,omitempty" validate:"required"`
}
