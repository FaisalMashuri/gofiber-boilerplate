package repository

import (
	"gofiber-boilerplate/internal/domain/user/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	// Todo : Register function
	Create(data *model.User) (err error)
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewSubmissionPartnerRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (repo *userRepositoryImpl) Create(data *model.User) (err error) {
	err = repo.db.Debug().Save(&data).Error
	if err != nil {
		return err
	}
	return
}
