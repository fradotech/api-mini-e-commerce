package service

import (
	"database/sql"
	"log"
	"mini_e_commerce/src/adaptor"
	"mini_e_commerce/src/server/params"
	"mini_e_commerce/src/server/repository"
	"mini_e_commerce/src/server/view"
)

type UserServices struct {
	repo            repository.UserRepo
	typicodeAdaptor *adaptor.TypicodeAdaptor
}

func NewUserServices(repo repository.UserRepo, typicodeAdaptor *adaptor.TypicodeAdaptor) *UserServices {
	return &UserServices{
		repo:            repo,
		typicodeAdaptor: typicodeAdaptor,
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

	return view.SuccessCreated(user)
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
		return view.ErrNotFound()
	}
	return view.SuccessFindAll(user)
}

func (c *UserServices) Update(email string, req *params.UserUpdate) *view.Response {
	user := req.ParseToModelCreate()

	getUser, err := c.repo.FindByEmail(email)
	if err != nil {
		return view.ErrBadRequest("user doesn't exists")
	}

	user.ID = getUser.ID
	err = c.repo.Update(email, user)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessUpdated(view.NewUserUpdateResponse(user))
}
