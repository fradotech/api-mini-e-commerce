package controller

import (
	"fmt"
	"mini_e_commerce/src/server/params"
	"mini_e_commerce/src/server/service"
	"mini_e_commerce/src/server/view"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc *service.UserServices
}

func NewUserHandler(svc *service.UserServices) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (u *UserHandler) Create(c *gin.Context) {
	var req params.UserCreate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp := view.ErrBadRequest(err.Error())
		WriteJsonResponseGin(c, resp)
		return
	}

	err = params.Validate(req)
	if err != nil {
		resp := view.ErrBadRequest(err.Error())
		WriteJsonResponseGin(c, resp)
		return
	}

	resp := u.svc.Create(&req)
	fmt.Println(resp)
	WriteJsonResponseGin(c, resp)
}

func (u *UserHandler) FindAll(c *gin.Context) {
	fmt.Println("Log from ", c.GetString("USER_EMAIL"))
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

	resp := u.svc.FindAll()
	fmt.Println("zzzzzzzzzzzzzzzzz")

	WriteJsonResponseGin(c, resp)
}
