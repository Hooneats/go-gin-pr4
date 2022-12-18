package controller

import (
	"errors"
	ctl "github.com/Hooneats/go-gin-pr4/controller/person"
	"github.com/Hooneats/go-gin-pr4/model"
	"github.com/Hooneats/go-gin-pr4/model/person"
)

var instance *control

type control struct {
	personCtl ctl.PersonController
}

func GetControl(mod model.Modeler) (*control, error) {
	if instance != nil {
		return instance, nil
	}
	if mod == nil {
		return nil, errors.New("modeler must is not nil")
	}
	personModel := person.GetPersonModel(mod)
	personModel.CreateIndex(mod, ctl.Name, ctl.Pnum)
	instance = &control{
		personCtl: ctl.GetPersonControl(personModel),
	}
	return instance, nil
}

func (c *control) PersonControl() ctl.PersonController {
	return c.personCtl
}
