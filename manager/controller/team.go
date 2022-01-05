package controller

import (
	"log"
	"manager/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTeams(c *gin.Context) {
	var team []models.Team
	err := models.GetAllTeams(&team)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, team)
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
