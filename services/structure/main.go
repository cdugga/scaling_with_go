package main

import (
	"github.com/cdugga/user-service-go/controller"
	"github.com/cdugga/user-service-go/database"
	"github.com/cdugga/user-service-go/env"
	"github.com/cdugga/user-service-go/http"
	"github.com/cdugga/user-service-go/service"
)

var (
	en         env.Provider      = env.NewEnv()
	db         database.Provider = database.NewPG()
	mainRouter http.Router       = http.NewMuxRouter()
)

func main() {
	initApp()
	initRoutes()

	mainRouter.Serve()
}

func initApp() {
	// Init environment provider
	en.Init()
	// Connect database
	db.Connect(en)
	// Init service
	service.NewUserService()
}

func initRoutes() {
	mainRouter.Get("/health", controller.Health)

	userRouter := mainRouter.RegisterSubRoute("/user")
	userRouter.Post("/signup", controller.Signup)
	userRouter.Post("/login", controller.Login)
	userRouter.Get("/{id}", controller.GetUserById)

	orgRouter := mainRouter.RegisterSubRoute("/org")
	orgRouter.Get("/{id}", controller.GetOrgById)
}