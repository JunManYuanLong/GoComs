package request

type (
	AddUserRequest struct {
		Nickname        string `json:"nickname"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
		Name            string `json:"name"`
		Email           string `json:"email"`
		Telephone       string `json:"telephone"`
		Picture         string `json:"picture"`
	}

	UpdateUserRequest struct {
		Id        int    `json:"id"`
		Nickname  string `json:"nickname"`
		Email     string `json:"email"`
		Telephone string `json:"telephone"`
		Picture   string `json:"picture"`
	}

	UserBindGroupRequest struct {
		UserId  int `json:"user_id"`
		GroupId int `json:"group_id"`
	}

	ParsePassword struct {
		
	}

	UpdatePasswordRequest struct {
		UserId      int    `json:"user_id"`
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	ResetPasswordRequest struct {
		UserId      int    `json:"user_id"`
		NewPassword string `json:"new_password"`
	}
)
