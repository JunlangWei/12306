package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string
	PassWord string
	IDCard   string
	Name     string
	Phone    string
	Type     string
	Email    *string
}
