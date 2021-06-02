package model

import "ict.com/public/model"

type UserGroup struct {
	model.EntityModel
	GroupName string `json:"group_name"`
	Comment   string `json:"comment"`
}
