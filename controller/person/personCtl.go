package person

import (
	api "github.com/Hooneats/go-gin-pr4/common"
	"github.com/Hooneats/go-gin-pr4/model"
	"github.com/Hooneats/go-gin-pr4/model/person"
	"github.com/Hooneats/go-gin-pr4/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var instance *PersonCtl

type PersonCtl struct {
	PersonModel person.PersonModeler
}

const colName = "tPerson"

func GetPersonCtl(m model.Modeler) *PersonCtl {
	if instance != nil {
		return instance
	}
	instance = &PersonCtl{
		PersonModel: person.GetPersonModel(m.GetCollection("tPerson")),
	}
	return instance
}

func (pCtl *PersonCtl) GetPersonByName(c *gin.Context) {
	_, cancel := util.GetContext(util.ControllerTimeOut)
	defer cancel()
	name := c.Param("name")

	var resData api.ApiResponse[any]
	findPerson, err := pCtl.PersonModel.FindPersonByName(name)
	if err != nil {
		log.Fatal(err)
		resData = api.Fail(api.NewError(err, http.StatusNotFound))
	}
	resPerson := NewWebPerson(*findPerson)

	resData = api.Success(resPerson)

	log.Println(resData)
	//jsonRes, _ := json.Marshal(resData)
	c.JSON(http.StatusOK, resData)
}
