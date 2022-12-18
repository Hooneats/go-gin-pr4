package main

import (
	"context"
	"github.com/Hooneats/go-gin-pr4/config"
	ctl "github.com/Hooneats/go-gin-pr4/controller"
	"github.com/Hooneats/go-gin-pr4/model"
	"github.com/Hooneats/go-gin-pr4/route"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	g errgroup.Group
)

func main() {
	config.EnvLoad()
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

		//우아한 종료
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		log.Println("Shutdown Server ...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := mapi.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown", err)
		}

		select {
		case <-ctx.Done():
			log.Println("timeout of 5 seconds.")
		}
		log.Println("Server exiting")
	}

	err := g.Wait()
	if err != nil {
		log.Println(err)
	}
}
