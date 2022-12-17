package person

import "go.mongodb.org/mongo-driver/bson/primitive"

type Person struct {
	id   primitive.ObjectID `bson:"_id"`
	Name string             `validate:"required" bson:"name"`
	Age  int                `validate:"required" bson:"age"`
	Pnum string             `bson:"pnum,omitempty"`
}

func NewPerson(name, pnum string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
		Pnum: pnum,
	}
}

func (p *Person) EditPerson(name, pnum string, age int) {
	p.Name = name
	p.Age = age
	p.Pnum = pnum
}

func (p *Person) GetId() primitive.ObjectID {
	return p.id
}
