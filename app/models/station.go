package models

import (
	"github.com/mamachengcheng/12306/app/utils"
	"gorm.io/gorm"
)

func init() {
	if utils.MysqlDBErr == nil && !utils.MysqlDB.Migrator().HasTable(&Station{}) {
		_ = utils.MysqlDB.Migrator().CreateTable(&Station{})
	}
}

type Station struct {
	gorm.Model
	Name        string `gorm:"not null;unique"`
	InitialName string `gorm:"not null;unique"`
	Pinyin      string `gorm:"not null;unique"`
	CityId      string `gorm:"not null;unique"`
	CityName    string `gorm:"not null;unique"`
	ShowName    string `gorm:"not null;unique"`
	NameType    string `gorm:"not null;unique"`
}
