package person

import (
	"github.com/Hooneats/go-gin-pr4/controller/person"
	"github.com/gin-gonic/gin"
)

func AddPersonRoute(rg *gin.RouterGroup, pct person.PersonCtler) {
	persons := rg.Group("/persons")
	{
		persons.GET("/:name", pct.GetPersonByName)
	}
}
