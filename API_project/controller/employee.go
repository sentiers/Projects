package controller

import (
	"log"
	"manager/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
