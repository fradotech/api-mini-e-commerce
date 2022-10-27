package server

import (
	"mini_e_commerce/src/server/controller"

	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	router     *gin.Engine
	auth       *controller.AuthHandler
	user       *controller.UserHandler
	category   *controller.CategoryHandler
	product    *controller.ProductHandler
	middleware *Middleware
}

func NewRouterGin(
	router *gin.Engine,
	auth *controller.AuthHandler,
	user *controller.UserHandler,
	category *controller.CategoryHandler,
	product *controller.ProductHandler,
	middleware *Middleware,
) *GinRouter {
	return &GinRouter{
		router:     router,
		auth:       auth,
		user:       user,
		category:   category,
		product:    product,
		middleware: middleware,
	}
}

func (r *GinRouter) Start(port string) {
	r.router.Use(r.middleware.Trace)

	auth := r.router.Group("/auth")
	auth.GET("/", r.middleware.Auth, r.middleware.CheckRole(r.auth.Profile, []string{"admin", "cashier", "member"}))
	auth.POST("/register", r.auth.Register)
	auth.POST("/login", r.auth.Login)

	user := r.router.Group("/users")
	user.GET("/", r.middleware.Auth, r.middleware.CheckRole(r.user.FindAll, []string{"admin"}))
	user.POST("/", r.middleware.Auth, r.middleware.CheckRole(r.user.Create, []string{"admin"}))

	category := r.router.Group("/categories")
	category.GET("/", r.middleware.Auth, r.middleware.CheckRole(r.category.GetCategories, []string{"admin", "cashier", "member"}))
	category.GET("/:categoryId", r.middleware.Auth, r.middleware.CheckRole(r.category.GetCategoryById, []string{"admin", "cashier", "member"}))
	category.POST("/", r.middleware.Auth, r.middleware.CheckRole(r.category.CreateCategory, []string{"admin"}))
	category.PUT("/:categoryId", r.middleware.Auth, r.middleware.CheckRole(r.category.UpdateCategory, []string{"admin"}))
	category.PATCH("/:categoryId", r.middleware.Auth, r.middleware.CheckRole(r.category.DeleteCategory, []string{"admin"}))

	product := r.router.Group("/products")
	product.GET("/", r.middleware.Auth, r.middleware.CheckRole(r.product.GetProducts, []string{"admin", "cashier", "member"}))
	product.GET("/:productId", r.middleware.Auth, r.middleware.CheckRole(r.product.GetProductById, []string{"admin", "cashier", "member"}))
	product.POST("/", r.middleware.Auth, r.middleware.CheckRole(r.product.CreateProduct, []string{"admin"}))
	product.PUT("/:productId", r.middleware.Auth, r.middleware.CheckRole(r.product.UpdateProduct, []string{"admin"}))
	product.PATCH("/:productId", r.middleware.Auth, r.middleware.CheckRole(r.product.DeleteProduct, []string{"admin"}))

	r.router.Run(port)
}
