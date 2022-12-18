package route

import (
	"errors"
	ctl "github.com/Hooneats/go-gin-pr4/controller"
	rt "github.com/Hooneats/go-gin-pr4/route/person"
	"github.com/gin-gonic/gin"
)

var instance *Router

type Router struct {
	ct ctl.Controller
}

//type NewRouterFunc func(ctl *ctl.Controller) (*Router, error)

func GetRouter(ctl ctl.Controller) (*Router, error) {
	if ctl == nil {
		return nil, errors.New("controller must is not nil")
	}
	if instance != nil {
		return instance, nil
	}
	instance = &Router{ct: ctl}
	return instance, nil
}

// 실제 라우팅
func (p *Router) Handle() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(CORS())

	v1 := r.Group("/v1")
	rt.AddPersonRoute(v1, p.ct.PersonControl())

	return r
}
