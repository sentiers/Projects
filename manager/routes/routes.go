package routes

import (
	"manager/controller"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.Default()

	r.GET("company", controller.GetAllCompanies)
	r.POST("company", controller.CreateCompany)

	return r
}
