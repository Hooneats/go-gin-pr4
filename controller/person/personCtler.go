package person

import "github.com/gin-gonic/gin"

type PersonController interface {
	GetByName(c *gin.Context)
}
