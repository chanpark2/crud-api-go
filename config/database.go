package config

import (
	"Users/witch/IdeaProjects/frameworktest/helper"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	host   = "localhost"
	port   = 3306
	user   = "witchtest"
	pass   = "witch12#"
	DBName = "witchdb"
)

func DatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.ErrorPanic(err)

	return db
}
