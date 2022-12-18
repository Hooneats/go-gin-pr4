package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiResponse[T any] struct {
	Code    int    `json:"code"`
	Data    T      `json:"data"`
	Message string `json:"message"`
	Err     error  `json:"err"`
}

func (a *ApiResponse[T]) Response(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(a.Code, a)
}
func Success() *ApiResponse[any] {
	return SuccessAndCustomMessage("success", "ok")
}

func SuccessData[T any](d T) *ApiResponse[any] {
	return SuccessAndCustomMessage(d, "ok")
}

func SuccessAndCustomMessage[T any](d T, msg string) *ApiResponse[any] {
	return SuccessCustom(http.StatusOK, d, msg)
}

func SuccessCustom[T any](c int, d T, msg string) *ApiResponse[any] {
	return &ApiResponse[any]{
		Code:    c,
		Data:    d,
		Message: msg,
		Err:     nil,
	}
}

func Fail(e Error) *ApiResponse[interface{}] {
	return FailCustomMessage(e, e.Err.Error())
}

func FailCustomMessage(e Error, msg string) *ApiResponse[interface{}] {
	return &ApiResponse[interface{}]{
		Code:    e.Code,
		Data:    nil,
		Message: msg,
		Err:     e.Err,
	}
}
