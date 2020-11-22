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
	Name        string
	InitialName string
	Pinyin      string
	CityId      string
	CityName    string
	ShowName    string
	NameType    string
}
