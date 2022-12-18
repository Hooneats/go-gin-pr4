package controller

import (
	"errors"
	ctl "github.com/Hooneats/go-gin-pr4/controller/person"
	"github.com/Hooneats/go-gin-pr4/model"
)

var instance *control

type control struct {
	model     model.Modeler
	personCtl ctl.PersonController
}

func GetControl(mod model.Modeler) (*control, error) {
	if instance != nil {
		return instance, nil
	}
	if mod == nil {
		return nil, errors.New("modeler must is not nil")
	}
	instance = &control{
		model:     mod,
		personCtl: ctl.GetPersonControl(mod),
	}
	return instance, nil
}

func (c *control) PersonControl() ctl.PersonController {
	return c.personCtl
}
