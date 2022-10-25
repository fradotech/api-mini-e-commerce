package gorm_mysql

import (
	"mini_e_commerce/src/server/model"
	"mini_e_commerce/src/server/repository"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repository.UserRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) Create(user *model.User) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (u *userRepo) FindAll() (*[]model.User, error) {
	var users []model.User
	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (u *userRepo) FindById(id string) (*model.User, error) {
	var user model.User
	err := u.db.Where("id=?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := u.db.Where("email=?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) Update(user *model.User) (*model.User, error) {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(user).Error; err != nil {
			return err
		}
		return nil
	})
	return user, err
}

func (u *userRepo) Delete(id string) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(id).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
