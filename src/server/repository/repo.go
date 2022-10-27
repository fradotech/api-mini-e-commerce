package repository

import "mini_e_commerce/src/server/model"

type UserRepo interface {
	Create(user *model.User) error
	FindAll() (*[]model.User, error)
	FindById(id string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(id string) error
}

type CategoryRepo interface {
	GetCategories() (*[]model.Category, error)
	CreateCategory(m *model.Category) error
	UpdateCategory(id string, m *model.Category) error
	DeleteCategory(id string, m *model.Category) error
	GetCategoryById(categoryId string) (*model.Category, error)
}

type ProductRepo interface {
	GetProducts() (*[]model.Product, error)
	CreateProduct(m *model.Product) error
	UpdateProduct(id string, m *model.Product) error
	DeleteProduct(id string, m *model.Product) error
	GetProductById(productId string) (*model.Product, error)
}
