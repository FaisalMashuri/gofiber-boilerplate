package service

import (
	"gofiber-boilerplate/internal/domain/user/model"
	"gofiber-boilerplate/internal/domain/user/repository"
)

type UserService interface {
	//	Todo : register function
	Create(data model.User) (err error)
}

type userServiceImpl struct {
	repo repository.UserRepository
}

func NewUserServicee(repo *repository.UserRepository) UserService {
	return &userServiceImpl{
		repo: *repo,
	}
}

func (service *userServiceImpl) Create(data model.User) (err error) {
	// write the logic
	err = service.repo.Create(&data)
	if err != nil {
		return
	}

	return
}
