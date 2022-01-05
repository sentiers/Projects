package controller

import (
	"log"
	"manager/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllDepartments(c *gin.Context) {
	var department []models.Department
	err := models.GetAllDepartments(&department)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, department)
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
