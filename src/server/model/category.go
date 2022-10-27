package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	UserId      string `json:"user_id"`

	User User `json:"user"`
}

var Categories = []Category{}
