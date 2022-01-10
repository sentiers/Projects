package main

import (
	"fmt"
	"manager/admin"
	"manager/config"
	"manager/models"
	"manager/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func main() {

	// initialize database
	dsn := "root:root@tcp(127.0.0.1:3306)/manager?charset=utf8mb4&parseTime=True&loc=Local"
	config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Status: ", err)
	}
	// table created automatically
	err = config.DB.AutoMigrate(&models.Company{}, &models.Department{}, &models.Team{}, &models.Employee{}, &admin.User{})

	r := routes.Routers() // routers
	r.Use(gin.Logger())   // use logger middleware

	err = r.Run() // run
}
