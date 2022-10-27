package server

import (
	"log"
	"mini_e_commerce/src/helper"
	"mini_e_commerce/src/server/controller"
	"mini_e_commerce/src/server/model"
	"mini_e_commerce/src/server/service"
	"mini_e_commerce/src/server/view"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	userSvc *service.UserServices
}

func NewMiddleware(userSvc *service.UserServices) *Middleware {
	return &Middleware{
		userSvc: userSvc,
	}
}

func (m *Middleware) Auth(c *gin.Context) {
	// get bearer token
	bearerToken := c.GetHeader("Authorization") // Bearer <token>

	// get token
	tokenArr := strings.Split(bearerToken, "Bearer ")

	// validate
	if len(tokenArr) != 2 {
		c.Set("ERROR", "no token")
		controller.WriteErrorJsonResponseGin(c, view.ErrUnauthorized())
		return
	}

	// verify token
	myTok, err := helper.VerifyToken(tokenArr[1])
	if err != nil {
		c.Set("ERROR", err.Error())
		controller.WriteErrorJsonResponseGin(c, view.ErrUnauthorized())
		return
	}

	// send to next handler
	user := m.userSvc.FindByEmail(myTok.Email)
	userDetail := user.Payload.(*model.User)
	c.Set("USER_ID", strconv.FormatUint(uint64(userDetail.ID), 10))
	c.Set("USER_EMAIL", userDetail.Email)
	c.Set("USER_NAME", userDetail.Fullname)

	// process to another handler
	c.Next()

}

func (m *Middleware) CheckRole(next gin.HandlerFunc, roles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := ctx.GetString("USER_EMAIL")
		user := m.userSvc.FindByEmail(email)
		userDetail := user.Payload.(*model.User)

		isExist := false

		for _, role := range roles {
			if role == userDetail.Role {
				isExist = true
				break
			}
		}

		if !isExist {
			controller.WriteErrorJsonResponseGin(ctx, view.ErrUnauthorized())
			return
		}

		next(ctx)
	}
}

func (m *Middleware) Trace(c *gin.Context) {
	now := time.Now()
	log.Printf("Get request with method :%v Path :%v\n", c.Request.Method, c.Request.URL)
	c.Next()
	isError := c.GetString("ERROR")
	if isError != "" {
		log.Printf("get error when try to get all typicode :%v\n", isError)
	}
	log.Printf("Finised request with method :%v Path :%v\n", c.Request.Method, c.Request.URL)

	end := time.Since(now).Milliseconds()
	log.Println("response time:", end)
}
