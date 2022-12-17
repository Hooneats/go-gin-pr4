package person

import (
	"context"
	"github.com/Hooneats/go-gin-pr4/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var instance *personModel

type personModel struct {
	collection *mongo.Collection
}

func GetPersonModel(col *mongo.Collection) *personModel {
	if instance != nil {
		return instance
	}
	createIndex(col)
	instance = &personModel{
		collection: col,
	}
	return instance
}

func createIndex(col *mongo.Collection) {
	ctx, cancel := util.GetContext(util.ModelTimeOut)
	defer cancel()
	_, err := col.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{Keys: bson.M{"name": 1}, Options: options.Index().SetUnique(true)},
		{Keys: bson.M{"pnum": 1}, Options: options.Index().SetUnique(true)},
	})
	if err != nil {
		log.Fatal(err)
		return
	}
}

func (p *personModel) FindPersonByName(name string) (*Person, error) {
	ctx, cancel := util.GetContext(util.ModelTimeOut)
	defer cancel()

	filterM := bson.M{"name": name}
	person, err := p.findOne(ctx, filterM)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return person, nil
}

func (p *personModel) FindPersonByPnum(pnum string) (*Person, error) {
	ctx, cancel := util.GetContext(util.ModelTimeOut)
	defer cancel()

	filterM := bson.M{"pnum": pnum}
	person, err := p.findOne(ctx, filterM)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return person, nil
}
func (p *personModel) findOne(ctx context.Context, filterM bson.M) (*Person, error) {
	var person *Person

	if err := p.collection.FindOne(ctx, filterM).Decode(&person); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return person, nil
}
func (p *personModel) FindAllPerson() ([]*Person, error) {
	ctx, cancel := util.GetContext(util.ModelTimeOut)
	defer cancel()

	cursor, err := p.collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var persons []*Person
	if err = cursor.All(ctx, &persons); err != nil {
		log.Fatal(err)
		return nil, err
	}

	findPersons := make([]*Person, len(persons))
	for index, person := range persons {
		if err := cursor.Decode(&person); err != nil {
			log.Fatal(err)
			return nil, err
		}
		findPersons[index] = person
	}
	return findPersons, err
}
func (p *personModel) InsertPerson(person *Person) (*Person, error) {
	ctx, cancel := util.GetContext(util.ModelTimeOut)
	defer cancel()

	person.id = primitive.NewObjectID()
	result, err := p.collection.InsertOne(ctx, person)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Printf("Inserted person +count %d, id is %s\n", 1, result.InsertedID)
	return person, nil
}
func (p *personModel) DeletePersonByPnum(pnum string) error {
	ctx, cancel := util.GetContext(util.ModelTimeOut)
	defer cancel()

	filter := bson.M{"pnum": pnum}
	result, err := p.collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("Deleted person -count %d by pnum %s\n", result.DeletedCount, pnum)
	return nil
}
func (p *personModel) UpdatePersonAgeByPnum(age int, pnum string) error {
	ctx, cancel := util.GetContext(util.ModelTimeOut)
	defer cancel()

	filter := bson.D{{"pnum", pnum}}
	update := bson.D{{"$set", bson.M{"age": age}}}
	_, err := p.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("Updated %s person By pnum %s, After age is %d\n", pnum, age)
	return nil
}