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

func (a *UserHandler) Profile(c *gin.Context) {
	fmt.Println("Log from ", c.GetString("USER_EMAIL"))
	resp := a.svc.FindByEmail(c.GetString("USER_EMAIL"))

	WriteJsonResponseGin(c, resp)
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

	resp := u.svc.FindAll()

	WriteJsonResponseGin(c, resp)
}

func (u *UserHandler) FindByEmail(c *gin.Context) {
	fmt.Println("Log from ", c.GetString("USER_EMAIL"))
	email, isExist := c.Params.Get("email")
	if !isExist {
		WriteJsonResponseGin(c, view.ErrBadRequest("email not found in params"))
		return
	}

	resp := u.svc.FindByEmail(email)

	WriteJsonResponseGin(c, resp)
}

func (u *UserHandler) Update(c *gin.Context) {
	email := c.GetString("USER_EMAIL")

	var req params.UserUpdate

	fmt.Println((req))
	err := c.ShouldBindJSON(&req)
	if err != nil {
		WriteJsonResponseGin(c, view.ErrBadRequest(err.Error()))
		return
	}

	err = params.Validate(req)
	if err != nil {
		user := view.ErrBadRequest(err.Error())
		WriteJsonResponseGin(c, user)
		return
	}

	resp := u.svc.Update(email, &req)

	WriteJsonResponseGin(c, resp)
}
