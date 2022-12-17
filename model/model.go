package model

import (
	"github.com/Hooneats/go-gin-pr4/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var instance *model

type model struct {
	client     *mongo.Client
	collection map[string]*mongo.Collection
}

func GetModel() (Modeler, error) {
	if instance != nil {
		return instance, nil
	}

	ctx, cancel := util.GetContext(util.ModelTimeOut)
	defer cancel()

	m := &model{}
	var err error
	mongoUri := os.Getenv("MONGOURI")
	opt := options.Client().SetMaxPoolSize(100).SetTimeout(util.DatabaseTimeOut)

	if m.client, err = mongo.Connect(ctx, opt.ApplyURI(mongoUri)); err != nil {
		return nil, err
	}

	if err = m.client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	instance = m
	return instance, nil
}

func (m *model) GetCollection(collection string) *mongo.Collection {
	if m.collection == nil {
		m.collection = make(map[string]*mongo.Collection)
	}
	if col, exists := m.collection[collection]; exists {
		return col
	}

	dbName := os.Getenv("DATABASE")
	db := m.client.Database(dbName)
	col := db.Collection(collection)
	m.collection[collection] = col
	return col
}
