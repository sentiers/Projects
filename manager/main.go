package main

import (
	"fmt"
	"manager/config"
	"manager/models"
	"manager/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func main() {

	// open database
	dsn := "root:root@tcp(127.0.0.1:3306)/manager?charset=utf8mb4&parseTime=True&loc=Local"
	config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Status: ", err)
	}
	err = config.DB.AutoMigrate(&models.Company{}, &models.Department{}, &models.Team{}, &models.Employee{}, &models.User{})

	r := routes.Routers()
	r.Use(gin.Logger())

	err = r.Run()
}
