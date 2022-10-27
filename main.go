package main

import (
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

	// rajaongkirAdaptor := adaptor.NewRajaOngkirAdaptor("https://api.rajaongkir.com")

	userRepo := gorm_mysql.NewUserRepo(db)
	authSvc := service.NewAuthServices(userRepo)
	authHandler := controller.NewAuthHandler(authSvc)

	userSvc := service.NewUserServices(userRepo)
	userHandler := controller.NewUserHandler(userSvc)

	categoryRepo := gorm_mysql.NewCategoryRepo(db)
	categorySvc := service.NewCategoryServices(categoryRepo)
	categoryHandler := controller.NewCategoryHandler(categorySvc)

	productRepo := gorm_mysql.NewProductRepo(db)
	productSvc := service.NewProductServices(productRepo, categoryRepo)
	productHandler := controller.NewProductHandler(productSvc)

	cartRepo := gorm_mysql.NewCartRepo(db)
	cartSvc := service.NewCartServices(cartRepo, productRepo, userRepo)
	cartHandler := controller.NewCartHandler(cartSvc)

	router := gin.Default()
	router.Use(gin.Logger())

	middleware := server.NewMiddleware(userSvc)

	app := server.NewRouterGin(
		router,
		authHandler,
		userHandler,
		categoryHandler,
		productHandler,
		cartHandler,
		middleware,
	)

	app.Start(":4444")
}
