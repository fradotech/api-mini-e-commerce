package server

import (
	"mini_e_commerce/src/server/controller"

	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	router     *gin.Engine
	auth       *controller.AuthHandler
	user       *controller.UserHandler
	middleware *Middleware
}

func NewRouterGin(
	router *gin.Engine,
	middleware *Middleware,
	auth *controller.AuthHandler,
	user *controller.UserHandler,
) *GinRouter {
	return &GinRouter{
		router:     router,
		auth:       auth,
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
	user.GET("/", r.user.FindAll, r.middleware.Auth, r.middleware.CheckRole(r.auth.Profile, []string{"admin"}))
	user.POST("/", r.user.Create, r.middleware.Auth, r.middleware.CheckRole(r.auth.Profile, []string{"admin"}))

	r.router.Run(port)
}
