package params

import (
	"mini_e_commerce/src/server/model"
)

type AuthRegister struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

type UserCreate struct {
	Fullname   string `validate:"required"`
	Gender     string `validate:"required"`
	Contact    string `validate:"required"`
	Street     string `validate:"required"`
	CityId     string `validate:"required" json:"city_id"`
	ProvinceId string `validate:"required" json:"province_id"`
}

type UserUpdate struct {
	UserCreate
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
