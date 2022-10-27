package params

import (
	"mini_e_commerce/src/server/model"
)

type CategoryCreate struct {
	Name        string `validate:"required"`
	Description string `validate:"required"`
	UserId      string
}

type CategoryUpdate struct {
	CategoryCreate
}

func (c *CategoryCreate) ParseToModel() *model.Category {
	return &model.Category{
		Name:        c.Name,
		Description: c.Description,
		UserId:      c.UserId,
	}
}
