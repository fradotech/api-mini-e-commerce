package params

import (
	"errors"
	"mini_e_commerce/src/server/model"

	"github.com/go-playground/validator/v10"
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

func Validate(u interface{}) error {
	err := validator.New().Struct(u)
	if err == nil {
		return nil
	}
	myErr := err.(validator.ValidationErrors)
	errString := ""
	for _, e := range myErr {
		errString += e.Field() + " is " + e.Tag()
	}
	return errors.New(errString)
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
