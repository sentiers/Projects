package routes

import (
	"manager/controller"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.Default()

	r.GET("company", controller.GetAllCompanies)
	r.POST("company", controller.CreateCompany)
	r.GET("company/:id", controller.GetCompanyById)
	r.PUT("company/:id", controller.UpdateCompany)
	r.DELETE("company/:id", controller.DeleteCompany)
	return r
}
