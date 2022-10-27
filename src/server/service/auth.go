package service

import (
	"database/sql"
	"log"
	"mini_e_commerce/src/helper"
	"mini_e_commerce/src/server/params"
	"mini_e_commerce/src/server/repository"
	"mini_e_commerce/src/server/view"
)

type AuthServices struct {
	repo repository.UserRepo
}

func NewAuthServices(repo repository.UserRepo) *AuthServices {
	return &AuthServices{
		repo: repo,
	}
}

func (u *AuthServices) Register(req *params.AuthRegister) *view.Response {
	user := req.ParseToModelRegister()

	user.Role = "member"

	hash, err := helper.GeneratePassword(user.Password)
	if err != nil {
		log.Printf("get error when try to generate password %v\n", "")
		return view.ErrInternalServer(err.Error())
	}

	user.Password = hash

	err = u.repo.Create(user)
	if err != nil {
		log.Printf("get error register user with error %v\n", "")
		return view.ErrInternalServer(err.Error())
	}

	type Succcess struct {
		Message     string
		GeneralInfo string
	}

	var success Succcess

	success.Message = "REGISTER_SUCCESS"
	success.GeneralInfo = user.Email

	return view.SuccessCreated(success)
}

func (u *AuthServices) Login(req *params.UserLogin) *view.Response {
	user, err := u.repo.FindByEmail(req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer(err.Error())
	}

	err = helper.ValidatePassword(user.Password, req.Password)
	if err != nil {
		return view.ErrUnauthorized()
	}

	token := helper.Token{
		Email: user.Email,
	}

	tokString, err := helper.CreateToken(&token)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessCreated(tokString)
}
