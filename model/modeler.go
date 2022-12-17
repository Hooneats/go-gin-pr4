package model

import "go.mongodb.org/mongo-driver/mongo"

type Modeler interface {
	GetCollection(collection string) *mongo.Collection
}

//type NewModelFunc func() (Modeler, error)
