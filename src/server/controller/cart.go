package controller

import (
	"mini_e_commerce/src/server/params"
	"mini_e_commerce/src/server/service"
	"mini_e_commerce/src/server/view"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	svc *service.CartService
}

func NewCartHandler(svc *service.CartService) *CartHandler {
	return &CartHandler{
		svc: svc,
	}
}

func (p *CartHandler) GetCarts(c *gin.Context) {
	resp := p.svc.GetCarts()

	WriteJsonResponseGin(c, resp)
}

func (p *CartHandler) GetUserCarts(c *gin.Context) {
	userId := c.GetString("USER_ID")
	resp := p.svc.GetUserCarts(userId)

	WriteJsonResponseGin(c, resp)
}

func (m *CartHandler) CreateCart(c *gin.Context) {
	var req params.CartCreate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		WriteJsonResponseGin(c, view.ErrBadRequest(err.Error()))
		return
	}

	err = params.Validate(req)
	if err != nil {
		resp := view.ErrBadRequest(err.Error())
		WriteJsonResponseGin(c, resp)
		return
	}

	userId := c.GetString("USER_ID")
	req.UserId = userId
	resp := m.svc.CreateCart(userId, &req)

	WriteJsonResponseGin(c, resp)
}

func (m *CartHandler) UpdateCart(c *gin.Context) {
	cartId, isExist := c.Params.Get("cartId")
	if !isExist {
		WriteJsonResponseGin(c, view.ErrBadRequest("cart not found"))
		return
	}

	var req params.CartUpdate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		WriteJsonResponseGin(c, view.ErrBadRequest(err.Error()))
		return
	}

	err = params.Validate(req)
	if err != nil {
		resp := view.ErrBadRequest(err.Error())
		WriteJsonResponseGin(c, resp)
		return
	}

	userId := c.GetString("USER_ID")
	req.UserId = userId
	resp := m.svc.UpdateCart(userId, cartId, &req)

	WriteJsonResponseGin(c, resp)
}

func (m *CartHandler) DeleteCart(c *gin.Context) {
	cartId, isExist := c.Params.Get("cartId")
	if !isExist {
		WriteJsonResponseGin(c, view.ErrBadRequest("cart not found"))
		return
	}

	resp := m.svc.DeleteCart(cartId)

	WriteJsonResponseGin(c, resp)
}

func (p *CartHandler) GetCartById(c *gin.Context) {
	cartId, isExist := c.Params.Get("cartId")
	if !isExist {
		WriteJsonResponseGin(c, view.ErrBadRequest("cart id not found in params"))
		return
	}

	resp := p.svc.GetCartById(cartId)
	WriteJsonResponseGin(c, resp)
}
