package main

import (
	"fmt"
	"manager/admin"
	"manager/config"
	"manager/models"
	"manager/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// initialize database
	err := config.InitDatabase()
	if err != nil {
		fmt.Println("Status: ", err)
	}
	// table created automatically
	config.DB.AutoMigrate(&models.Company{}, &models.Department{}, &models.Team{}, &models.Employee{}, &admin.User{})

	r := routes.Routers() // routers
	r.Use(gin.Logger())   // use logger middleware

	r.Run() // run in default PORT 8080
	// r.Run(":4000") // run in PORT 4000
}
