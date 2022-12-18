package controller

import (
	"github.com/Hooneats/go-gin-pr4/controller/person"
)

type Controller interface {
	PersonControl() person.PersonController
}
