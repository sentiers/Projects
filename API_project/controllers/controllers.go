package controllers

import (
	"log"
	"manager/models"
	"net/http"

	"github.com/gin-gonic/gin"

	"strconv"
)

// Company Controller ==============================================

func GetAllCompanies(c *gin.Context) {
	var company []models.Company
	err := models.GetAllCompanies(&company)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, company) // return all data
	}
}

func CreateCompany(c *gin.Context) {
	var company models.Company

	// BindJSON -> ShouldBindJSON because of the [WARNING]
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := models.CreateCompany(&company); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, company)
}

func GetCompanyById(c *gin.Context) {
	var company models.Company
	id := c.Params.ByName("id")

	if err := models.GetCompanyById(&company, id); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, company)
}

func UpdateCompany(c *gin.Context) {
	var company models.Company
	id := c.Params.ByName("id")

	if err := models.GetCompanyById(&company, id); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	if err := c.BindJSON(&company); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	if err := models.UpdateCompany(&company); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, company)
}

func DeleteCompany(c *gin.Context) {
	var company models.Company
	id := c.Params.ByName("id")

	if err := models.DeleteCompany(&company, id); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
}

// Department Controller ==============================================

func GetAllDepartments(c *gin.Context) {
	var department []models.Department
	err := models.GetAllDepartments(&department)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, department) // return all data
	}
}

func CreateDepartment(c *gin.Context) {
	var department models.Department

	// BindJSON -> ShouldBindJSON because of the [WARNING]
	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := models.CreateDepartment(&department); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, department)
}

func GetDepartmentById(c *gin.Context) {
	var department models.Department
	id := c.Params.ByName("id")

	if err := models.GetDepartmentById(&department, id); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, department)
}

func UpdateDepartment(c *gin.Context) {
	var department models.Department
	id := c.Params.ByName("id")

	if err := models.GetDepartmentById(&department, id); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	if err := c.BindJSON(&department); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	if err := models.UpdateDepartment(&department); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, department)
}

func DeleteDepartment(c *gin.Context) {
	var department models.Department
	id := c.Params.ByName("id")

	if err := models.DeleteDepartment(&department, id); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
}

// Team Controller ==============================================
func GetAllTeams(c *gin.Context) {
	var team []models.Team
	err := models.GetAllTeams(&team)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, team) // return all data
	}
}

func CreateTeam(c *gin.Context) {
	var team models.Team

	// BindJSON -> ShouldBindJSON because of the [WARNING]
	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := models.CreateTeam(&team); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, team)
}

func GetTeamById(c *gin.Context) {
	var team models.Team
	id := c.Params.ByName("id")

	if err := models.GetTeamById(&team, id); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, team)
}

func UpdateTeam(c *gin.Context) {
	var team models.Team
	id := c.Params.ByName("id")

	if err := models.GetTeamById(&team, id); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	if err := c.BindJSON(&team); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	if err := models.UpdateTeam(&team); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, team)
}

func DeleteTeam(c *gin.Context) {
	var team models.Team
	id := c.Params.ByName("id")

	if err := models.DeleteTeam(&team, id); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
}

// Employee Controller ==============================================
func GetAllEmployees(c *gin.Context) {
	pagination := GeneratePagination(c)
	var employee []models.Employee
	if err := models.GetAllEmployees(&employee, &pagination); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, employee)
}

func CreateEmployee(c *gin.Context) {
	var employee models.Employee

	// BindJSON -> ShouldBindJSON because of the [WARNING]
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := models.CreateEmployee(&employee); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, employee)
}

func GetEmployeeById(c *gin.Context) {
	var employee models.Employee
	id := c.Params.ByName("id")

	if err := models.GetEmployeeById(&employee, id); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, employee)
}

func UpdateEmployee(c *gin.Context) {
	var employee models.Employee
	id := c.Params.ByName("id")

	if err := models.GetEmployeeById(&employee, id); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	if err := c.BindJSON(&employee); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	if err := models.UpdateEmployee(&employee); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, employee)
}

func DeleteEmployee(c *gin.Context) {
	var employee models.Employee
	id := c.Params.ByName("id")

	if err := models.DeleteEmployee(&employee, id); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
}

// =========================================

// search
func GetEmployeeByName(c *gin.Context) {
	var employee []models.Employee
	name := c.Params.ByName("name") // ex) John

	if err := models.GetEmployeeByName(&employee, name); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, employee)
}

// filter
func GetEmployeeByDate(c *gin.Context) {

	var employee []models.Employee
	date := c.Params.ByName("date") // ex) 20200125

	if err := models.GetEmployeeByDate(&employee, date); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, employee)
}

// =========================================

// pagination ex) ?page=2&limit=10&sort=created_at
func GeneratePagination(c *gin.Context) models.Pagination {
	// Initializing default ?page=1&limit=10
	limit := 10
	page := 1
	sort := "id"
	query := c.Request.URL.Query() // map[limit:[10] page:[2]]

	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break
		}
	}
	return models.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}
