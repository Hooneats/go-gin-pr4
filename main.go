package main

import (
	"github.com/Hooneats/go-gin-pr4/config"
	ctl "github.com/Hooneats/go-gin-pr4/controller"
	"github.com/Hooneats/go-gin-pr4/model"
	"github.com/Hooneats/go-gin-pr4/route"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	g errgroup.Group
)

func main() {
	config.EnvLoad()

	////model 모듈 선언
	//app := util.App()
	//app.RegisterModel(model.GetModel)
	//app.RegisterController(ctl.GetController)
	//app.RegisterRouter(route.GetRouter)
	//app.RegisterOptions(func() *http.Server {
	//	port := os.Getenv("PORT")
	//	return &http.Server{
	//		Addr:           port,
	//		Handler:        app.Router.Handle(),
	//		ReadTimeout:    2 * time.Second,
	//		WriteTimeout:   5 * time.Second,
	//		MaxHeaderBytes: 1 << 20,
	//	}
	//}())
	//g.Go(func() error {
	//	return app.Run()
	//})

	//model 모듈 선언
	if mod, err := model.GetModel(); err != nil {
		panic(err)
	} else if controller, err := ctl.GetControl(mod); err != nil { //controller 모듈 설정
		panic(err)
	} else if rt, err := route.GetRouter(controller); err != nil { //router 모듈 설정
		panic(err)
	} else {
		port := os.Getenv("PORT")
		mapi := &http.Server{
			Addr:           port,
			Handler:        rt.Handle(),
			ReadTimeout:    2 * time.Second,
			WriteTimeout:   5 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

		g.Go(func() error {
			return mapi.ListenAndServe()
		})
	}

	err := g.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
