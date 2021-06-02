package model

import "ict.com/public/model"

type User struct {
	model.EntityModel
	Nickname  string `json:"nickname"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	Picture   string `json:"picture"`
}
