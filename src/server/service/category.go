package service

import (
	"database/sql"
	"fmt"
	"mini_e_commerce/src/server/params"
	"mini_e_commerce/src/server/repository"
	"mini_e_commerce/src/server/view"
)

type CategoryService struct {
	repo repository.CategoryRepo
}

func NewCategoryServices(repo repository.CategoryRepo) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (c *CategoryService) GetCategories() *view.Response {
	categories, err := c.repo.GetCategories()
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessFindAll(view.NewCategoryGetAllResponse(categories))
}

func (c *CategoryService) CreateCategory(req *params.CategoryCreate) *view.Response {
	category := req.ParseToModel()
	fmt.Println(category)
	err := c.repo.CreateCategory(category)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessCreated(view.NewCategoryCreateResponse(category))
}

func (c *CategoryService) UpdateCategory(categoryId string, req *params.CategoryUpdate) *view.Response {
	category := req.ParseToModel()

	getCategory, err := c.repo.GetCategoryById(categoryId)
	if err != nil {
		return view.ErrBadRequest("category doesn't exists")
	}

	category.ID = getCategory.ID
	err = c.repo.UpdateCategory(categoryId, category)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessUpdated(view.NewCategoryUpdateResponse(category))
}

func (p *CategoryService) DeleteCategory(categoryId string) *view.Response {
	category, err := p.repo.GetCategoryById(categoryId)
	if err != nil {
		return view.ErrBadRequest("category doesn't exists")
	}

	err = p.repo.DeleteCategory(categoryId, category)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessDeleted(view.NewCategoryUpdateResponse(category))
}

func (c *CategoryService) GetCategoryById(categoryId string) *view.Response {
	category, err := c.repo.GetCategoryById(categoryId)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessFindAll(category)
}
