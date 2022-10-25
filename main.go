package main

import (
	"mini_e_commerce/src/adaptor"
	"mini_e_commerce/src/db"
	"mini_e_commerce/src/server"
	"mini_e_commerce/src/server/controller"
	"mini_e_commerce/src/server/repository/gorm_mysql"
	"mini_e_commerce/src/server/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.ConnectGormDB()
	if err != nil {
		panic(err)
	}

	typicodeAdaptor := adaptor.NewTypicodeAdaptor("https://jsonplaceholder.typicode.com/posts")

	userRepo := gorm_mysql.NewUserRepo(db)
	authSvc := service.NewAuthServices(userRepo, typicodeAdaptor)
	authHandler := controller.NewAuthHandler(authSvc)

	userSvc := service.NewUserServices(userRepo, typicodeAdaptor)
	userHandler := controller.NewUserHandler(userSvc)

	router := gin.Default()
	router.Use(gin.Logger())

	middleware := server.NewMiddleware(userSvc)

	app := server.NewRouterGin(
		router,
		middleware,
		authHandler,
		userHandler,
	)

	app.Start(":4444")
}
