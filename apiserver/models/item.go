package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name,omitempty"`
	Price int                `bson:"price,omitempty"`
}
