package models

import (
	"github.com/mamachengcheng/12306/app/utils"
	"gorm.io/gorm"
)

func init() {
	if utils.MysqlDBErr == nil && !utils.MysqlDB.Migrator().HasTable(&Station{}) {
		err := utils.MysqlDB.Migrator().CreateTable(&Station{})
		if err != nil {
			panic(err)
		}
	}
}

type Station struct {
	gorm.Model
}