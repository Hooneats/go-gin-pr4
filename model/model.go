package model

import (
	"github.com/Hooneats/go-gin-pr4/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
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

func (m *model) CreateIndex(colName string, indexName ...string) {
	ctx, cancel := util.GetContext(util.ModelTimeOut)
	defer cancel()
	col := m.GetCollection(colName)
	var indexModels []mongo.IndexModel
	for _, name := range indexName {
		idxModel := mongo.IndexModel{
			Keys: bson.M{name: 1}, Options: options.Index().SetUnique(true),
		}
		indexModels = append(indexModels, idxModel)
	}
	_, err := col.Indexes().CreateMany(ctx, indexModels)
	if err != nil {
		log.Println(err)
		return
	}
}
