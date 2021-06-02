package main

import (
	"github.com/cdugga/bookmark/controller"
	"github.com/cdugga/bookmark/env"
	"github.com/cdugga/bookmark/http"
)

var (
	en         env.Provider      = env.NewEnv()
	mainRouter http.Router       = http.NewMuxRouter()
)

func main(){
	initApp()
	initRoutes()

	mainRouter.Serve()
}

func initApp() {
	// Init environment provider
	en.Init()


}

func initRoutes() {
	mainRouter.Get("/{id}", controller.GetOrgById)

	//userRouter := mainRouter.RegisterSubRoute("/user")
	//userRouter.Post("/signup", controller.Signup)
	//userRouter.Post("/login", controller.Login)
	//userRouter.Get("/{id}", controller.GetUserById)
	//
	//orgRouter := mainRouter.RegisterSubRoute("/org")
	//orgRouter.Get("/{id}", controller.GetOrgById)
}