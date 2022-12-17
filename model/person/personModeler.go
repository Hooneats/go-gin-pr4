package person

type PersonModeler interface {
	FindPersonByName(name string) (*Person, error)
	FindPersonByPnum(pnum string) (*Person, error)
	FindAllPerson() ([]*Person, error)
	InsertPerson(person *Person) (*Person, error)
	DeletePersonByPnum(pnum string) error
	UpdatePersonAgeByPnum(age int, pnum string) error
}
