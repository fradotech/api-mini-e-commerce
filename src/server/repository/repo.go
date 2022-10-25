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
