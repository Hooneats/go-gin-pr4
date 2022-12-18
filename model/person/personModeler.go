package person

type PersonModeler interface {
	FindByName(name string) (*Person, error)
	FindByPnum(pnum string) (*Person, error)
	FindAll() ([]*Person, error)
	InsertOne(person *Person) (*Person, error)
	DeleteByPnum(pnum string) error
	UpdateAgeByPnum(age int, pnum string) error
}
