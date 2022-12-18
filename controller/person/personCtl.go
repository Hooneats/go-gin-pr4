package person

import (
	api "github.com/Hooneats/go-gin-pr4/common"
	"github.com/Hooneats/go-gin-pr4/model"
	"github.com/Hooneats/go-gin-pr4/model/person"
	"github.com/Hooneats/go-gin-pr4/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

var instance *PersonControl

type PersonControl struct {
	PersonModel person.PersonModeler
}

const colName = "tPerson"

func GetPersonControl(m model.Modeler) *PersonControl {
	if instance != nil {
		return instance
	}
	instance = &PersonControl{
		PersonModel: person.GetPersonModel(m.GetCollection(colName)),
	}
	return instance
}

func (pCtl *PersonControl) GetByName(c *gin.Context) {
	_, cancel := util.GetContext(util.ControllerTimeOut)
	defer cancel()
	name := c.Param(Name)

	findPerson, err := pCtl.PersonModel.FindByName(name)
	if err != nil {
		log.Println(err)
		api.Fail(api.NewError(err, http.StatusNotFound)).Response(c)
	} else {
		personData := NewWebPerson(*findPerson)
		api.SuccessData(personData).Response(c)
	}
}

func (pCtl *PersonControl) GetByPnum(c *gin.Context) {
	_, cancel := util.GetContext(util.ControllerTimeOut)
	defer cancel()
	pnum := c.Param(Pnum)

	findPerson, err := pCtl.PersonModel.FindByPnum(pnum)
	if err != nil {
		log.Println(err)
		api.Fail(api.NewError(err, http.StatusNotFound)).Response(c)
	} else {
		personData := NewWebPerson(*findPerson)
		api.SuccessData(personData).Response(c)
	}
}
func (pCtl *PersonControl) GetAll(c *gin.Context) {
	_, cancel := util.GetContext(util.ControllerTimeOut)
	defer cancel()

	findPersons, err := pCtl.PersonModel.FindAll()
	if err != nil {
		log.Println(err)
		api.Fail(api.NewError(err, http.StatusNotFound)).Response(c)
	} else {
		personDatas := make([]*WebPerson, len(findPersons))
		for index, findPerson := range findPersons {
			personDatas[index] = NewWebPerson(*findPerson)
		}
		api.SuccessData(personDatas).Response(c)
	}
}
func (pCtl *PersonControl) PostOne(c *gin.Context) {
	_, cancel := util.GetContext(util.ControllerTimeOut)
	defer cancel()
	var webPerson *WebPerson
	err := c.BindJSON(&webPerson)
	if err != nil {
		log.Println(err)
		api.Fail(api.NewError(err, http.StatusBadRequest)).Response(c)
		return
	}

	person := webPerson.NewPerson()
	intertedP, err := pCtl.PersonModel.InsertOne(person)
	if err != nil {
		api.Fail(api.NewError(err, http.StatusBadRequest)).Response(c)
	} else {
		api.SuccessData(intertedP).Response(c)
	}
}
func (pCtl *PersonControl) DeleteByPnum(c *gin.Context) {
	_, cancel := util.GetContext(util.ControllerTimeOut)
	defer cancel()
	pnum := c.Query(Pnum)

	err := pCtl.PersonModel.DeleteByPnum(pnum)
	if err != nil {
		api.Fail(api.NewError(err, http.StatusBadRequest)).Response(c)
	} else {
		api.Success().Response(c)
	}
}
func (pCtl *PersonControl) PutAgeByPnum(c *gin.Context) {
	_, cancel := util.GetContext(util.ControllerTimeOut)
	defer cancel()

	pnum := c.Query(Pnum)
	ageStr := c.Query(Age)
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		api.Fail(api.NewError(err, http.StatusBadRequest)).Response(c)
		return
	}

	err = pCtl.PersonModel.UpdateAgeByPnum(age, pnum)
	if err != nil {
		api.Fail(api.NewError(err, http.StatusBadRequest)).Response(c)
	} else {
		api.Success().Response(c)
	}
}
