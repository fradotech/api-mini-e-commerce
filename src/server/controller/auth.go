package controller

import (
	"fmt"
	"mini_e_commerce/src/server/params"
	"mini_e_commerce/src/server/service"
	"mini_e_commerce/src/server/view"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	svc     *service.AuthServices
	userSvc *service.UserServices
}

func NewAuthHandler(svc *service.AuthServices) *AuthHandler {
	return &AuthHandler{
		svc: svc,
	}
}

func (a *AuthHandler) Profile(c *gin.Context) {
	fmt.Println("Log from ", c.GetString("USER_EMAIL"))
	resp := a.userSvc.FindByEmail(c.GetString("USER_EMAIL"))

	WriteJsonResponseGin(c, resp)
}

func (a *AuthHandler) Register(c *gin.Context) {
	var req params.AuthRegister
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

	resp := a.svc.Register(&req)
	fmt.Println(resp)
	WriteJsonResponseGin(c, resp)
}

func (a *AuthHandler) Login(c *gin.Context) {
	var req params.UserLogin
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

	resp := a.svc.Login(&req)
	WriteJsonResponseGin(c, resp)
}
