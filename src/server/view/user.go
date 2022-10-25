package view

import "mini_e_commerce/src/server/model"

type UserCreateResponse struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

func NewUserCreateResponse(user *model.User) *UserCreateResponse {
	return &UserCreateResponse{
		Fullname: user.Fullname,
		Email:    user.Email,
	}
}

type UserFindAllResponse struct {
	Id       string `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Gender   string `json:"gender"`
	Contact  string `json:"contact"`
	Street   string `json:"street"`
}

func NewUserFindAllResponse(users *[]model.User) []UserFindAllResponse {
	var usersFind []UserFindAllResponse
	for _, user := range *users {
		usersFind = append(usersFind, *parseModelToUserFind(&user))
	}
	return usersFind
}

func parseModelToUserFind(user *model.User) *UserFindAllResponse {
	return &UserFindAllResponse{
		Fullname: user.Fullname,
		Email:    user.Email,
		Gender:   user.Gender,
		Contact:  user.Contact,
		Street:   user.Street}
}
