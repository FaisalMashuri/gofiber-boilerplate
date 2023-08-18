package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       string `gorm:"primary_key"`
	Name     string
	Email    string
	Password string
	phone    string
}
