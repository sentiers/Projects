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
