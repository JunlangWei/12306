package models

import (
	"github.com/mamachengcheng/12306/app/utils"
	"gorm.io/gorm"
)

func init() {
	if utils.MysqlDBErr == nil && !utils.MysqlDB.Migrator().HasTable(&User{}) {
		_ = utils.MysqlDB.Migrator().CreateTable(&User{})
	}
}

type User struct {
	gorm.Model
	Username string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	IDCard   string `gorm:"not null;unique"`
	Name     string `gorm:"not null"`
	Phone    string `gorm:"not null;unique"`
	Type     string `gorm:"not null"`
	Email    string `gorm:"not null;unique"`
}
