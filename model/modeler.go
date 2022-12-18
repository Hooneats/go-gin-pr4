package model

import "go.mongodb.org/mongo-driver/mongo"

type Modeler interface {
	GetCollection(collection string) *mongo.Collection
	CreateIndex(colName string, indexName ...string)
}

//type NewModelFunc func() (Modeler, error)
