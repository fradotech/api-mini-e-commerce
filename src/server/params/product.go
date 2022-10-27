package params

import (
	"mini_e_commerce/src/server/model"
)

type ProductCreate struct {
	Name        string `validate:"required"`
	Description string `validate:"required"`
	Price       uint32 `validate:"required"`
	Stock       uint16 `validate:"required"`
	CategoryId  string `validate:"required" json:"category_id"`
	UserId      string `json:"user_id"`
}

type ProductUpdate struct {
	ProductCreate
}

func (p *ProductCreate) ParseToModel() *model.Product {
	return &model.Product{
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Stock:       p.Stock,
		CategoryId:  p.CategoryId,
		UserId:      p.UserId,
	}
}
