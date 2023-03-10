package person

import "github.com/Hooneats/go-gin-pr4/model/person"

type WebPerson struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Pnum string `json:"pnum"`
}

const (
	Name = "name"
	Age  = "age"
	Pnum = "pnum"
)

func NewWebPerson(p person.Person) *WebPerson {
	return &WebPerson{
		Name: p.Name,
		Age:  p.Age,
		Pnum: p.Pnum,
	}
}

func (p *WebPerson) NewPerson() *person.Person {
	return &person.Person{
		Name: p.Name,
		Age:  p.Age,
		Pnum: p.Pnum,
	}
}
