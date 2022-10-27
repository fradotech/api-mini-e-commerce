package service

import (
	"database/sql"
	"log"
	"mini_e_commerce/src/server/params"
	"mini_e_commerce/src/server/repository"
	"mini_e_commerce/src/server/view"
)

type UserServices struct {
	repo repository.UserRepo
}

func NewUserServices(repo repository.UserRepo) *UserServices {
	return &UserServices{
		repo: repo,
	}
}

func (u *UserServices) Create(req *params.UserCreate) *view.Response {
	user := req.ParseToModelCreate()

	user.Role = "member"

	var err error

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

	success.Message = "CREATED_USER_SUCCESS"
	success.GeneralInfo = user.Email

	return view.SuccessCreated(success)
}

func (u *UserServices) FindAll() *view.Response {

	users, err := u.repo.FindAll()

	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessFindAll(view.NewUserFindAllResponse(users))
}

func (u *UserServices) FindByEmail(email string) *view.Response {
	user, err := u.repo.FindByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer(err.Error())
	}
	return view.SuccessFindAll(user)
}
