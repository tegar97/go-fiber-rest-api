package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password []byte
}

type UsersRegister struct {
	name     string `json:"name"`
	email    string `gorm:"unique"`
	password []byte `json:"password"`
}
