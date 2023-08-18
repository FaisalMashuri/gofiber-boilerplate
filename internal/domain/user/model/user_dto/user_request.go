package user_dto

import "gofiber-boilerplate/internal/domain/user/model"

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserRequest) ToModeluser() model.User {
	return model.User{
		Email:    u.Email,
		Password: u.Password,
	}
}
