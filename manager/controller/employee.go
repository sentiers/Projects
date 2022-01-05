package controller

import (
	"log"
	"manager/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllEmployees(c *gin.Context) {
	var employee []models.Employee
	err := models.GetAllEmployees(&employee)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, employee)
	}
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
