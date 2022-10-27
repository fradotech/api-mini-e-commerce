package controller

import (
	"fmt"
	"mini_e_commerce/src/server/params"
	"mini_e_commerce/src/server/service"
	"mini_e_commerce/src/server/view"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	svc *service.ProductService
}

func NewProductHandler(svc *service.ProductService) *ProductHandler {
	return &ProductHandler{
		svc: svc,
	}
}

func (p *ProductHandler) GetProducts(c *gin.Context) {
	fmt.Println("heoooo")
	resp := p.svc.GetProducts()

	fmt.Println("heaaaa")
	WriteJsonResponseGin(c, resp)
}

func (m *ProductHandler) CreateProduct(c *gin.Context) {
	var req params.ProductCreate
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
	resp := m.svc.CreateProduct(&req)

	WriteJsonResponseGin(c, resp)
}

func (m *ProductHandler) UpdateProduct(c *gin.Context) {
	productId, isExist := c.Params.Get("productId")
	if !isExist {
		WriteJsonResponseGin(c, view.ErrBadRequest("product not found"))
		return
	}

	var req params.ProductUpdate
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

	resp := m.svc.UpdateProduct(productId, &req)

	WriteJsonResponseGin(c, resp)
}

func (m *ProductHandler) DeleteProduct(c *gin.Context) {
	productId, isExist := c.Params.Get("productId")
	if !isExist {
		WriteJsonResponseGin(c, view.ErrBadRequest("product not found"))
		return
	}

	resp := m.svc.DeleteProduct(productId)

	WriteJsonResponseGin(c, resp)
}

func (p *ProductHandler) GetProductById(c *gin.Context) {
	productId, isExist := c.Params.Get("productId")
	if !isExist {
		WriteJsonResponseGin(c, view.ErrBadRequest("productId not found in params"))
		return
	}

	resp := p.svc.GetProductById(productId)
	WriteJsonResponseGin(c, resp)
}
