package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Place struct {
	Protected

	Menu      []primitive.ObjectID `bson:"menu,omitempty"`
	Address   string               `bson:"address,omitempty"`
	City      string               `bson:"city,omitempty"`
	Country   string               `bson:"country,omitempty"`
	Latitude  float32              `bson:"latitude,omitempty"`
	Longitude float32              `bson:"longitude,omitempty"`
}
