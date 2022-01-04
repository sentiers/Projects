package controller

import (
	"log"
	"manager/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCompanies(c *gin.Context) {
	var company []models.Company
	err := models.GetAllCompanies(&company)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, company)
	}
}

func CreateCompany(c *gin.Context) {
	var company models.Company

	if err := c.BindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := models.CreateCompany(&company); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, company)
}
