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
		PersonModel: person.GetPersonModel(m.GetCollection("tPerson")),
	}
	return instance
}

func (pCtl *PersonControl) GetByName(c *gin.Context) {
	_, cancel := util.GetContext(util.ControllerTimeOut)
	defer cancel()
	name := c.Param("name")

	var response api.ApiResponse[any]
	findPerson, err := pCtl.PersonModel.FindByName(name)
	if err != nil {
		log.Fatal(err)
		response = api.Fail(api.NewError(err, http.StatusNotFound))
	} else {
		personData := NewWebPerson(*findPerson)
		response = api.SuccessData(personData)
	}
	log.Println(response)
	c.JSON(http.StatusOK, response)
}

func (pCtl *PersonControl) GetByPnum(c *gin.Context) {
	_, cancel := util.GetContext(util.ControllerTimeOut)
	defer cancel()
	pnum := c.Param("pnum")

	var response api.ApiResponse[any]
	findPerson, err := pCtl.PersonModel.FindByPnum(pnum)
	if err != nil {
		log.Fatal(err)
		response = api.Fail(api.NewError(err, http.StatusNotFound))
	} else {
		personData := NewWebPerson(*findPerson)
		response = api.SuccessData(personData)
	}

	log.Println(response)
	//jsonRes, _ := json.Marshal(resData)
	c.JSON(http.StatusOK, response)
}
func (pCtl *PersonControl) GetAll(c *gin.Context) {
	_, cancel := util.GetContext(util.ControllerTimeOut)
	defer cancel()

	var resData api.ApiResponse[any]
	findPersons, err := pCtl.PersonModel.FindAll()
	if err != nil {
		log.Fatal(err)
		resData = api.Fail(api.NewError(err, http.StatusNotFound))
	} else {
		personDatas := make([]*WebPerson, len(findPersons))
		for index, findPerson := range findPersons {
			personDatas[index] = NewWebPerson(*findPerson)
		}
		resData = api.SuccessData(personDatas)
	}
	log.Println(resData)
	//jsonRes, _ := json.Marshal(resData)
	c.JSON(http.StatusOK, resData)

}
func (pCtl *PersonControl) PostOne(c *gin.Context) {
	_, cancel := util.GetContext(util.ControllerTimeOut)
	defer cancel()
	var resData api.ApiResponse[any]
	var webPerson *WebPerson
	err := c.BindJSON(webPerson)
	if err != nil {
		log.Fatal(err)
		resData = api.Fail(api.NewError(err, http.StatusBadRequest))
		c.JSON(http.StatusBadRequest, resData)
		return
	}

	person := webPerson.NewPerson()
	intertedP, err := pCtl.PersonModel.InsertOne(person)
	if err != nil {
		resData = api.Fail(api.NewError(err, http.StatusBadRequest))
		c.JSON(http.StatusBadRequest, resData)
		return
	} else {
		resData = api.SuccessData(intertedP)
		log.Println(resData)
		//jsonRes, _ := json.Marshal(resData)
		c.JSON(http.StatusOK, resData)
	}
}
func (pCtl *PersonControl) DeleteByPnum(c *gin.Context) {
	_, cancel := util.GetContext(util.ControllerTimeOut)
	defer cancel()
	var resData api.ApiResponse[any]
	pnum := c.Param("pnum")

	err := pCtl.PersonModel.DeleteByPnum(pnum)
	if err != nil {
		resData = api.Fail(api.NewError(err, http.StatusBadRequest))
		c.JSON(http.StatusBadRequest, resData)
		return
	} else {
		resData = api.Success()
		c.JSON(http.StatusOK, resData)
	}
}
func (pCtl *PersonControl) PutAgeByPnum(c *gin.Context) {
	_, cancel := util.GetContext(util.ControllerTimeOut)
	defer cancel()

	var resData api.ApiResponse[any]
	pnum := c.Param("pnum")
	ageStr := c.Param("age")
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		resData = api.Fail(api.NewError(err, http.StatusBadRequest))
		c.JSON(http.StatusBadRequest, resData)
		return
	}

	err = pCtl.PersonModel.UpdateAgeByPnum(age, pnum)
	if err != nil {
		resData = api.Fail(api.NewError(err, http.StatusBadRequest))
		c.JSON(http.StatusBadRequest, resData)
		return
	} else {
		resData = api.Success()
		c.JSON(http.StatusOK, resData)
	}
}
