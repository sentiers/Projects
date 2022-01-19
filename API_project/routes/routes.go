package routes

import (
	"manager/admin"
	"manager/controllers"
	"manager/middleware"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {

	// create gin app
	r := gin.Default()

	// root path for testing
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})

	// company api group
	c := r.Group("/company")
	{
		c.Use(middleware.Authz())
		c.GET("/", controllers.GetAllCompanies)
		c.POST("/", controllers.CreateCompany)
		c.GET("/:id", controllers.GetCompanyById)
		c.PUT("/:id", controllers.UpdateCompany)
		c.DELETE("/:id", controllers.DeleteCompany)
	}

	// department api group
	d := r.Group("/department")
	{
		d.Use(middleware.Authz())
		d.GET("/", controllers.GetAllDepartments)
		d.POST("/", controllers.CreateDepartment)
		d.GET("/:id", controllers.GetDepartmentById)
		d.PUT("/:id", controllers.UpdateDepartment)
		d.DELETE("/:id", controllers.DeleteDepartment)
	}

	// team api group
	t := r.Group("/team")
	{
		t.Use(middleware.Authz())
		t.GET("/", controllers.GetAllTeams)
		t.POST("/", controllers.CreateTeam)
		t.GET("/:id", controllers.GetTeamById)
		t.PUT("/:id", controllers.UpdateTeam)
		t.DELETE("/:id", controllers.DeleteTeam)
	}

	// employee api group
	e := r.Group("/employee")
	{
		e.Use(middleware.Authz())
		e.GET("/", controllers.GetAllEmployees)
		e.POST("/", controllers.CreateEmployee)
		e.GET("/:id", controllers.GetEmployeeById)
		e.PUT("/:id", controllers.UpdateEmployee)
		e.DELETE("/:id", controllers.DeleteEmployee)
		// search and filter
		e.GET("/name/:name", controllers.GetEmployeeByName)
		e.GET("/date/:date", controllers.GetEmployeeByDate)
		// team_emp
		e.POST("/:id/addteam/:teamid", controllers.AddEmployeeTeam)
		e.DELETE("/:id/delteam/:teamid", controllers.DeleteEmployeeTeam)
	}

	// local auth api
	r.POST("/signup", admin.Signup)
	r.POST("/login", admin.Login)

	// google auth api
	r.GET("/google/login", admin.GoogleLogin)
	r.GET("/google/redirect", admin.GoogleRedirect)

	// github auth api
	r.GET("/github/login", admin.GithubLogin)
	r.GET("/github/redirect", admin.GithubRedirect)

	return r
}
