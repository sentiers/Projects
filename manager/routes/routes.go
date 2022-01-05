package routes

import (
	"manager/controller"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.Default()

	// company
	r.GET("company", controller.GetAllCompanies)
	r.POST("company", controller.CreateCompany)
	r.GET("company/:id", controller.GetCompanyById)
	r.PUT("company/:id", controller.UpdateCompany)
	r.DELETE("company/:id", controller.DeleteCompany)

	// department
	r.GET("department", controller.GetAllDepartments)
	r.POST("department", controller.CreateDepartment)
	r.GET("department/:id", controller.GetDepartmentById)
	r.PUT("department/:id", controller.UpdateDepartment)
	r.DELETE("department/:id", controller.DeleteDepartment)

	// team
	r.GET("team", controller.GetAllTeams)
	r.POST("team", controller.CreateTeam)
	r.GET("team/:id", controller.GetTeamById)
	r.PUT("team/:id", controller.UpdateTeam)
	r.DELETE("team/:id", controller.DeleteTeam)

	// employee
	r.GET("employee", controller.GetAllEmployees)
	r.POST("employee", controller.CreateEmployee)
	r.GET("employee/:id", controller.GetEmployeeById)
	r.PUT("employee/:id", controller.UpdateEmployee)
	r.DELETE("employee/:id", controller.DeleteEmployee)

	return r
}
