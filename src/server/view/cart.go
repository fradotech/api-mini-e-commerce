package view

import "mini_e_commerce/src/server/model"

type CartCreateResponse struct {
	Id        uint   `json:"id"`
	UserId    string `json:"user_id"`
	ProductId string `json:"product_id"`
	Qty       uint16 `json:"qty"`
	Subtotal  uint32 `json:"subtotal"`
}

func NewCartCreateResponse(cart *model.Cart) *CartCreateResponse {
	return &CartCreateResponse{
		Id:        cart.ID,
		UserId:    cart.UserId,
		ProductId: cart.ProductId,
		Qty:       cart.Qty,
		Subtotal:  cart.Subtotal,
	}
}

type CartUpdateResponse struct {
	Id        uint   `json:"id"`
	UserId    string `json:"user_id"`
	ProductId string `json:"product_id"`
	Qty       uint16 `json:"qty"`
	Subtotal  uint32 `json:"subtotal"`
}

func NewCartUpdateResponse(cart *model.Cart) *CartUpdateResponse {
	return &CartUpdateResponse{
		Id:        cart.ID,
		UserId:    cart.UserId,
		ProductId: cart.ProductId,
		Qty:       cart.Qty,
		Subtotal:  cart.Subtotal,
	}
}

type CartGetResponse struct {
	Id        uint   `json:"id"`
	UserId    string `json:"user_id"`
	ProductId string `json:"product_id"`
	Qty       uint16 `json:"qty"`
	Subtotal  uint32 `json:"subtotal"`
}

func NewCartGetResponse(carts *[]model.Cart) *[]CartGetResponse {
	var cartsResponse []CartGetResponse

	for _, cart := range *carts {
		cartsResponse = append(cartsResponse, CartGetResponse{
			Id:        cart.ID,
			UserId:    cart.UserId,
			ProductId: cart.ProductId,
			Qty:       cart.Qty,
			Subtotal:  cart.Subtotal,
		})
	}

	return &cartsResponse
}
