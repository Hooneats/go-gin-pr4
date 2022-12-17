package person

import "github.com/gin-gonic/gin"

type PersonCtler interface {
	GetPersonByName(c *gin.Context)
}
