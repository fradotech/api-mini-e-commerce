package view

import "mini_e_commerce/src/server/model"

type CategoryCreateResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoryUpdateResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewCategoryCreateResponse(category *model.Category) *CategoryCreateResponse {
	return &CategoryCreateResponse{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
}

func NewCategoryUpdateResponse(category *model.Category) *CategoryUpdateResponse {
	return &CategoryUpdateResponse{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
}

type CategoryGetAllResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewCategoryGetAllResponse(categories *[]model.Category) *[]CategoryGetAllResponse {
	var categoriesResponse []CategoryGetAllResponse

	for _, category := range *categories {
		categoriesResponse = append(categoriesResponse, CategoryGetAllResponse{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	return &categoriesResponse
}
