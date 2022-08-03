package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	Type string `bson:"type,omitempty" validate:"required"`
	Name string `bson:"name,omitempty" validate:"required"`
}
