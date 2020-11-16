package utils

import (
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetMysqlDB() *gorm.DB {
	cfg, _ := ini.Load("app/conf/config.ini")

	mysqlCfg := cfg.Section("mysql")
	username := mysqlCfg.Key("username").String()
	password := mysqlCfg.Key("password").String()
	address := mysqlCfg.Key("http").String() + ":" + mysqlCfg.Key("port").String()
	dbname := mysqlCfg.Key("dbname").String()

	dsn := username + ":" + password + "@tcp("+ address + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	client, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return client
}