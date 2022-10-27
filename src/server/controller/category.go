package controller

import (
	"mini_e_commerce/src/server/params"
	"mini_e_commerce/src/server/service"
	"mini_e_commerce/src/server/view"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	svc *service.CategoryService
}

func NewCategoryHandler(svc *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		svc: svc,
	}
}

func (m *CategoryHandler) GetCategories(c *gin.Context) {
	resp := m.svc.GetCategories()

	WriteJsonResponseGin(c, resp)

}

func (m *CategoryHandler) CreateCategory(c *gin.Context) {
	var req params.CategoryCreate

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
	resp := m.svc.CreateCategory(&req)

	WriteJsonResponseGin(c, resp)
}

func (m *CategoryHandler) UpdateCategory(c *gin.Context) {
	categoryId, isExist := c.Params.Get("categoryId")
	if !isExist {
		WriteJsonResponseGin(c, view.ErrBadRequest("category not found"))
		return
	}

	var req params.CategoryUpdate
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

	resp := m.svc.UpdateCategory(categoryId, &req)

	WriteJsonResponseGin(c, resp)
}

func (m *CategoryHandler) DeleteCategory(c *gin.Context) {
	categoryId, isExist := c.Params.Get("categoryId")
	if !isExist {
		WriteJsonResponseGin(c, view.ErrBadRequest("category not found"))
		return
	}

	resp := m.svc.DeleteCategory(categoryId)

	WriteJsonResponseGin(c, resp)
}

func (p *CategoryHandler) GetCategoryById(c *gin.Context) {
	categoryId, isExist := c.Params.Get("categoryId")
	if !isExist {
		WriteJsonResponseGin(c, view.ErrBadRequest("categoryId not found in params"))
		return
	}

	resp := p.svc.GetCategoryById(categoryId)
	WriteJsonResponseGin(c, resp)
}
