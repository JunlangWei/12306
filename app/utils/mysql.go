package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func GetMysqlDB() *gorm.DB {
	cfg, err := ini.Load("app/conf/config.ini")

	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	mysqlCfg := cfg.Section("mysql")
	username := mysqlCfg.Key("username").String()
	password := mysqlCfg.Key("password").String()
	address := mysqlCfg.Key("http").String() + ":" + mysqlCfg.Key("port").String()
	dbname := mysqlCfg.Key("dbname").String()

	dsn := username + ":" + password + "@tcp("+ address + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	return db
}