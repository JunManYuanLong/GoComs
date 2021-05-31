package model

import "ict.com/public.v1/model"

type UserGroup struct {
	model.EntityModel
	GroupName string `json:"group_name"`
	Comment   string `json:"comment"`
}
