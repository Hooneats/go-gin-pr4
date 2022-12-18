package person

import "github.com/gin-gonic/gin"

type PersonController interface {
	GetByName(c *gin.Context)
	GetByPnum(c *gin.Context)
	GetAll(c *gin.Context)
	PostOne(c *gin.Context)
	DeleteByPnum(c *gin.Context)
	PutAgeByPnum(c *gin.Context)
}
