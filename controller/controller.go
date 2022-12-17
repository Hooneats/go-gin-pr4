package controller

import (
	"errors"
	ctl "github.com/Hooneats/go-gin-pr4/controller/person"
	"github.com/Hooneats/go-gin-pr4/model"
)

var instance *Controller

type Controller struct {
	model     model.Modeler
	PersonCtl ctl.PersonCtler
}

//type NewControllerFunc func(model.Modeler) (*Controller, error)

func GetController(mod model.Modeler) (*Controller, error) {
	if instance != nil {
		return instance, nil
	}
	if mod == nil {
		return nil, errors.New("modeler must is not nil")
	}
	instance = &Controller{
		model:     mod,
		PersonCtl: ctl.GetPersonCtl(mod),
	}
	return instance, nil
}
