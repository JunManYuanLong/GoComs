package model

import (
	"ict.com/public.v1/model"
)

type AkSK struct {
	model.EntityModel
	AK     string `json:"ak"`
	SK     string `json:"sk"`
	UserId uint   `json:"user_id"`
}
