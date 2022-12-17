package util

//
//import (
//	"github.com/Hooneats/go-gin-pr4/controller"
//	"github.com/Hooneats/go-gin-pr4/model"
//	"github.com/Hooneats/go-gin-pr4/route"
//	"net/http"
//)
//
//type app struct {
//	model.Modeler
//	*controller.Controller
//	*route.Router
//	*http.Server
//}
//
//func App() *app {
//	return &app{}
//}
//
//func (ap *app) RegisterModel(modelFunc func() model.NewModelFunc) *app {
//	mod, err := modelFunc()()
//	if err != nil {
//		panic(err)
//	}
//	ap.Modeler = mod
//	return ap
//}
//
//func (ap *app) RegisterController(controllerFunc func() controller.NewControllerFunc) *app {
//	if ap.Modeler == nil {
//		panic("controller register should be after modeler register")
//	}
//	ctl, err := controllerFunc()(ap.Modeler)
//	if err != nil {
//		panic(err)
//	}
//	ap.Controller = ctl
//	return ap
//}
//
//func (ap *app) RegisterRouter(routerFunc func() route.NewRouterFunc) *app {
//	if ap.Controller == nil {
//		panic("router register should be after controller register")
//	}
//	router, err := routerFunc()(ap.Controller)
//	if err != nil {
//		panic(err)
//	}
//	ap.Router = router
//	return ap
//}
//
//func (ap *app) RegisterOptions(server *http.Server) *app {
//	if ap.Router == nil {
//		panic("if you want running app, should be after router register")
//	}
//	ap.Server = server
//	return ap
//}
//
//func (ap *app) Run() error {
//	if ap.Server == nil {
//		panic("if you want running app, should be after server options register")
//	}
//	return ap.Server.ListenAndServe()
//}
//
//func (ap *app) Finish(func()) {
//
//}
