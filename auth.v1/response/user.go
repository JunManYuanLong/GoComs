package response

import "ict.com/public.v1/model"

type QueryUserResponse struct {
	model.EntityModel
	NickName  string `json:"nickname"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	Picture   string `json:"picture"`
}
