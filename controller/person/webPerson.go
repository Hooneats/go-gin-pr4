package person

import "github.com/Hooneats/go-gin-pr4/model/person"

type WebPerson struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Pnum string `json:"pnum"`
}

func NewWebPerson(p person.Person) *WebPerson {
	return &WebPerson{
		Name: p.Name,
		Age:  p.Age,
		Pnum: p.Pnum,
	}
}
