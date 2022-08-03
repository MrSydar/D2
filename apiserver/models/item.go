package models

type Item struct {
	Protected

	Name  string `bson:"name,omitempty"`
	Price int    `bson:"price,omitempty"`
}
