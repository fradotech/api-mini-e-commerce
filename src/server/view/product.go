package view

import "mini_e_commerce/src/server/model"

type ProductCreateResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Price       uint32 `json:"price"`
	Description string `json:"description"`
	Stock       uint16 `json:"stock"`
}

func NewProductCreateResponse(product *model.Product) *ProductCreateResponse {
	return &ProductCreateResponse{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}
}

type ProductUpdateResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Price       uint32 `json:"price"`
	Description string `json:"description"`
	Stock       uint16 `json:"stock"`
}

func NewProductUpdateResponse(product *model.Product) *ProductUpdateResponse {
	return &ProductUpdateResponse{
		Id:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Stock:       product.Stock,
	}
}

type ProductGetAllResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"namdescription"`
	Price       uint32 `json:"price"`
	Stock       uint16 `json:"stock"`
}

func NewProductGetAllResponse(products *[]model.Product) *[]ProductGetAllResponse {
	var productsResponse []ProductGetAllResponse

	for _, product := range *products {
		productsResponse = append(productsResponse, ProductGetAllResponse{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
		})
	}

	return &productsResponse
}
