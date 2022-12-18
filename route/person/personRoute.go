package person

import (
	"errors"
	api "github.com/Hooneats/go-gin-pr4/common"
	"github.com/Hooneats/go-gin-pr4/controller/person"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddPersonRoute(rg *gin.RouterGroup, pct person.PersonController) {
	rg.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
	persons := rg.Group("/persons")
	{
		persons.GET("/", pct.GetAll)
		persons.GET("/name/:name", pct.GetByName)
		persons.GET("/pnum/:pnum", pct.GetByPnum)

		person := persons.Group("/person")
		{
			person.POST("", pct.PostOne)
			person.DELETE("", pct.DeleteByPnum)
			person.PUT("", pct.PutAgeByPnum)
		}
	}
}

func checkJsonType(c *gin.Context) {
	{
		contentType := c.Request.Header.Get("Content-Type")
		if contentType != "application/json" {
			api.Fail(api.NewError(errors.New("content type is must json or html"), http.StatusBadRequest)).Response(c)
			return
		}
		c.Next()
	}
}
