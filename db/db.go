// Copyright Korea Space Data Co., Ltd. 2020 All Rights Reserved.
package database

import (
	"fmt"
	. "ksd-grm-api/config"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var once sync.Once
var gormDB *gorm.DB

func GetDB() *gorm.DB {
	once.Do(func() {
		config := GetConfig()
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DB_USERNAME, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME)
		gormDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Printf("%+v", err)
		}
		gormDB = gormDb
	})
	return gormDB
}
