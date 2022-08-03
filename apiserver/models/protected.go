package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Protected struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	OwnerAccountID primitive.ObjectID `bson:"ownerAccountID,omitempty"`
}
