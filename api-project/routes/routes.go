package routes

import (
	"manager/controller"
	"manager/middleware"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {

	// create gin app
	r := gin.Default()

	// company api group
	c := r.Group("/company")
	{
		c.Use(middleware.Authz())
		c.GET("/", controller.GetAllCompanies)
		c.POST("/", controller.CreateCompany)
		c.GET("/:id", controller.GetCompanyById)
		c.PUT("/:id", controller.UpdateCompany)
		c.DELETE("/:id", controller.DeleteCompany)
	}

	// department api group
	d := r.Group("/department")
	{
		d.Use(middleware.Authz())
		d.GET("/", controller.GetAllDepartments)
		d.POST("/", controller.CreateDepartment)
		d.GET("/:id", controller.GetDepartmentById)
		d.PUT("/:id", controller.UpdateDepartment)
		d.DELETE("/:id", controller.DeleteDepartment)
	}

	// team api group
	t := r.Group("/team")
	{
		t.Use(middleware.Authz())
		t.GET("/", controller.GetAllTeams)
		t.POST("/", controller.CreateTeam)
		t.GET("/:id", controller.GetTeamById)
		t.PUT("/:id", controller.UpdateTeam)
		t.DELETE("/:id", controller.DeleteTeam)
	}

	// employee api group
	e := r.Group("/employee")
	{
		e.Use(middleware.Authz())
		e.GET("/", controller.GetAllEmployees)
		e.POST("/", controller.CreateEmployee)
		e.GET("/:id", controller.GetEmployeeById)
		e.PUT("/:id", controller.UpdateEmployee)
		e.DELETE("/:id", controller.DeleteEmployee)
	}

	// auth api
	r.POST("/signup", controller.Signup)
	r.POST("/login", controller.Login)

	return r
}
