package service

import (
	"database/sql"
	"mini_e_commerce/src/server/params"
	"mini_e_commerce/src/server/repository"
	"mini_e_commerce/src/server/view"
	"strconv"
)

type CartService struct {
	repo        repository.CartRepo
	productRepo repository.ProductRepo
	userRepo    repository.UserRepo
}

func NewCartServices(repo repository.CartRepo, productRepo repository.ProductRepo, userRepo repository.UserRepo) *CartService {
	return &CartService{
		repo:        repo,
		productRepo: productRepo,
		userRepo:    userRepo,
	}
}

func (p *CartService) GetCarts() *view.Response {
	carts, err := p.repo.GetCarts()
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessFindAll(view.NewCartGetResponse(carts))
}

func (p *CartService) GetUserCarts(userId string) *view.Response {
	carts, err := p.repo.GetUserCarts(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessFindAll(view.NewCartGetResponse(carts))
}

func (p *CartService) GetCartById(cartId string) *view.Response {
	cart, err := p.repo.GetCartById(cartId)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessFindAll(cart)
}

func (p *CartService) CreateCart(userId string, req *params.CartCreate) *view.Response {
	cart := req.ParseToModel()
	user, err := p.userRepo.FindById(userId)
	if err != nil {
		return view.ErrBadRequest("user id doesn't exists")
	}

	product, errProduct := p.productRepo.GetProductById(cart.ProductId)
	if errProduct != nil {
		return view.ErrBadRequest("product id doesn't exists")
	}

	productInCart, _ := p.repo.GetCartByProductId(userId, cart.ProductId)
	if productInCart != nil {
		productInCart.Qty = productInCart.Qty + req.Qty
		productInCart.Subtotal = uint32(productInCart.Qty) * product.Price
		cartId := strconv.FormatUint(uint64(productInCart.ID), 10)
		err = p.repo.UpdateCart(cartId, productInCart)
		if err != nil {
			return view.ErrInternalServer(err.Error())
		}
		return view.SuccessUpdated(view.NewCartUpdateResponse(productInCart))
	}

	cart.Subtotal = product.Price * uint32(cart.Qty)
	cart.User = *user
	cart.Product = *product
	err = p.repo.CreateCart(cart)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessCreated(view.NewCartCreateResponse(cart))
}

func (p *CartService) UpdateCart(userId string, cartId string, req *params.CartUpdate) *view.Response {
	cart := req.ParseToModel()
	user, err := p.userRepo.FindById(userId)
	if err != nil {
		return view.ErrBadRequest("user id doesn't exists")
	}

	product, errProduct := p.productRepo.GetProductById(cart.ProductId)
	if errProduct != nil {
		return view.ErrBadRequest("product id doesn't exists")
	}

	getCart, errCart := p.repo.GetCartById(cartId)
	if errCart != nil {
		return view.ErrBadRequest("cart doesn't exists")
	}

	cart.User = *user
	cart.Product = *product
	cart.ID = getCart.ID
	cart.Subtotal = product.Price * uint32(cart.Qty)
	err = p.repo.UpdateCart(cartId, cart)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessUpdated(view.NewCartUpdateResponse(cart))
}

func (p *CartService) DeleteCart(cartId string) *view.Response {
	cart, err := p.repo.GetCartById(cartId)
	if err != nil {
		return view.ErrBadRequest("cart doesn't exists")
	}

	err = p.repo.DeleteCart(cartId, cart)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessDeleted(view.NewCartUpdateResponse(cart))
}
