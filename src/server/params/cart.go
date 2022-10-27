package params

import (
	"mini_e_commerce/src/server/model"
)

type CartCreate struct {
	UserId    string
	ProductId string `validate:"required" json:"product_id"`
	Qty       uint16 `validate:"required"`
	Subtotal  uint32
}

type CartUpdate struct {
	CartCreate
}

func (p *CartCreate) ParseToModel() *model.Cart {
	return &model.Cart{
		ProductId: p.ProductId,
		Qty:       p.Qty,
	}
}
