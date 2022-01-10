package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// database that will be used across different packages
var DB *gorm.DB

// InitDatabase creates a db
func InitDatabase() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/manager?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	return
}
