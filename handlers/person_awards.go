package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/magadov/kinopoisk-app/database"
	"github.com/magadov/kinopoisk-app/models"
)

func AllPersonAw(c *gin.Context) {
	var persAwards []models.PersonAwards

	if res := database.DB.Find(&persAwards); res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, persAwards)
}

func CreatePersonAw(c *gin.Context) {
	var persAwardsInp models.PersonAwardsDTO

	if err := c.ShouldBindJSON(&persAwardsInp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	persAwards := models.PersonAwards{
		PersonID: persAwardsInp.PersonID,
		Category: persAwardsInp.Category,
		Year:     persAwardsInp.Year,
		Result:   persAwardsInp.Result,
	}

	if res := database.DB.Create(&persAwards); res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, persAwards)
}

func UpdatePersonAw(c *gin.Context) {
	id := c.Param("id")
	var personAwUpdInp models.PersonAwardsDTO

	if err := c.ShouldBindJSON(&personAwUpdInp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var personAwUpd models.PersonAwards
	if res := database.DB.Model(&models.PersonAwards{}).Where("id = ?", id).Updates(&personAwUpdInp).First(&personAwUpd, id); res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, personAwUpd)
}

func DeletePersonAw(c *gin.Context) {
	id := c.Param("id")
	var persAwardsDelete models.PersonAwards

	if res := database.DB.First(&persAwardsDelete, id); res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error.Error()})
		return
	}

	if res := database.DB.Delete(&persAwardsDelete); res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func WinnersOfYear(c *gin.Context) {
	var winPersons []models.PersonAwards

	if res := database.DB.Where("year = ? AND result = ?", c.Param("year"), "win").Find(&winPersons); res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error})
		return
	}

	c.JSON(http.StatusOK, winPersons)
}
