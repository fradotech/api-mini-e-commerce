package params

import (
	"mini_e_commerce/src/server/model"
)

type AuthRegister struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

type UserCreate struct {
	Fullname   string
	Gender     string
	Contact    string
	Street     string
	CityId     string
	ProvinceId string
}

type UserUpdate struct {
	Email    string
	Password string
}

type UserLogin struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

func (u *AuthRegister) ParseToModelRegister() *model.User {
	return &model.User{
		Email:    u.Email,
		Password: u.Password,
	}
}

func (u *UserCreate) ParseToModelCreate() *model.User {
	return &model.User{
		Fullname:   u.Fullname,
		Gender:     u.Gender,
		Contact:    u.Contact,
		Street:     u.Street,
		CityId:     u.CityId,
		ProvinceId: u.ProvinceId,
	}
}
