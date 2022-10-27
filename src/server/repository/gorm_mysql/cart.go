package gorm_mysql

import (
	"mini_e_commerce/src/server/model"
	"mini_e_commerce/src/server/repository"
	"time"

	"gorm.io/gorm"
)

type cartRepo struct {
	db *gorm.DB
}

func NewCartRepo(db *gorm.DB) repository.CartRepo {
	return &cartRepo{
		db: db,
	}
}

func (p *cartRepo) GetCarts() (*[]model.Cart, error) {
	var carts []model.Cart

	err := p.db.Where("deleted_at IS NULL").Find(&carts).Error
	if err != nil {
		return nil, err
	}

	return &carts, nil
}

func (p *cartRepo) GetUserCarts(userId string) (*[]model.Cart, error) {
	var carts []model.Cart

	err := p.db.Where("user_id = ?", userId).Where("deleted_at IS NULL").Find(&carts).Error
	if err != nil {
		return nil, err
	}

	return &carts, nil
}

func (p *cartRepo) GetCartById(cartId string) (*model.Cart, error) {
	var cart model.Cart
	err := p.db.Joins("User").Joins("Product").Where("carts.id = ?", cartId).Where("carts.deleted_at IS NULL").First(&cart).Error
	if err != nil {
		return nil, err
	}

	return &cart, nil
}

func (p *cartRepo) GetCartByProductId(userId string, productId string) (*model.Cart, error) {
	var cart model.Cart
	err := p.db.Where("carts.user_id = ?", userId).Where("carts.product_id = ?", productId).Where("carts.deleted_at IS NULL").First(&cart).Error
	if err != nil {
		return nil, err
	}

	return &cart, nil
}

func (p *cartRepo) CreateCart(cart *model.Cart) error {
	return p.db.Create(cart).Error
}

func (p *cartRepo) UpdateCart(id string, cart *model.Cart) error {
	return p.db.Where("id = ?", id).Updates(cart).Error
}

func (p *cartRepo) DeleteCart(id string, cart *model.Cart) error {
	return p.db.Model(&cart).Where("id = ?", id).Update("deleted_at", time.Now()).Error
}
