package gorm_mysql

import (
	"mini_e_commerce/src/server/model"
	"mini_e_commerce/src/server/repository"
	"time"

	"gorm.io/gorm"
)

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) repository.ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (p *productRepo) GetProducts() (*[]model.Product, error) {
	var products []model.Product

	err := p.db.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return &products, nil
}

func (p *productRepo) CreateProduct(product *model.Product) error {
	return p.db.Create(product).Error
}

func (p *productRepo) UpdateProduct(id string, product *model.Product) error {
	return p.db.Where("id = ?", id).Updates(product).Error
}

func (p *productRepo) DeleteProduct(id string, product *model.Product) error {
	return p.db.Model(&product).Where("id = ?", id).Update("deleted_at", time.Now()).Error
}

func (p *productRepo) GetProductById(productId string) (*model.Product, error) {
	var product model.Product
	err := p.db.Joins("Category").Where("products.id=?", productId).First(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}
