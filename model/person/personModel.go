package person

import (
	"context"
	"errors"
	"github.com/Hooneats/go-gin-pr4/model"
	"github.com/Hooneats/go-gin-pr4/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var instance *personModel

type personModel struct {
	collection *mongo.Collection
}

const TPerson = "tPerson"

func GetPersonModel(mod model.Modeler) *personModel {
	if instance != nil {
		return instance
	}
	col := mod.GetCollection(TPerson)
	instance = &personModel{
		collection: col,
	}
	return instance
}

func (p *personModel) CreateIndex(mod model.Modeler, indexName ...string) {
	mod.CreateIndex(TPerson, indexName...)
}

func (p *personModel) FindByName(name string) (*Person, error) {
	ctx, cancel := util.GetContext(util.ModelTimeOut)
	defer cancel()

	filterM := bson.M{"name": name}
	person, err := p.findOne(ctx, filterM)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return person, nil
}

func (p *personModel) FindByPnum(pnum string) (*Person, error) {
	ctx, cancel := util.GetContext(util.ModelTimeOut)
	defer cancel()

	filterM := bson.M{"pnum": pnum}
	person, err := p.findOne(ctx, filterM)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return person, nil
}
func (p *personModel) findOne(ctx context.Context, filterM bson.M) (*Person, error) {
	var person *Person

	if err := p.collection.FindOne(ctx, filterM).Decode(&person); err != nil {
		log.Println(err)
		return nil, err
	}

	return person, nil
}
func (p *personModel) FindAll() ([]*Person, error) {
	ctx, cancel := util.GetContext(util.ModelTimeOut)
	defer cancel()

	cursor, err := p.collection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var persons []*Person
	if err = cursor.All(ctx, &persons); err != nil {
		log.Println(err)
		return nil, err
	}
	return persons, err
}
func (p *personModel) InsertOne(person *Person) (*Person, error) {
	ctx, cancel := util.GetContext(util.ModelTimeOut)
	defer cancel()

	person.id = primitive.NewObjectID()
	result, err := p.collection.InsertOne(ctx, person)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Printf("Inserted person +count %d, id is %s\n", 1, result.InsertedID)
	return person, nil
}
func (p *personModel) DeleteByPnum(pnum string) error {
	ctx, cancel := util.GetContext(util.ModelTimeOut)
	defer cancel()

	filter := bson.M{"pnum": pnum}
	result, err := p.collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Println(err)
		return err
	}
	if result.DeletedCount <= 0 {
		err := errors.New("does not deleted person")
		log.Println(err)
		return err
	}
	log.Printf("Deleted person -count %d by pnum %s\n", result.DeletedCount, pnum)
	return nil
}
func (p *personModel) UpdateAgeByPnum(age int, pnum string) error {
	ctx, cancel := util.GetContext(util.ModelTimeOut)
	defer cancel()

	filter := bson.D{{"pnum", pnum}}
	update := bson.D{{"$set", bson.D{{"age", age}}}}
	result, err := p.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println(err)
		return err
	}
	if result.MatchedCount <= 0 {
		err := errors.New("does not updated person")
		log.Println(err)
		return err
	}
	log.Printf("Updated person By pnum %s, After age is %d\n", pnum, age)
	return nil
}
