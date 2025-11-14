package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/magadov/kinopoisk-app/database"
	"github.com/magadov/kinopoisk-app/models"
)

func AddPersons(c *gin.Context) {

	var input models.PersonsRequestDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка преобразования формата."})
		return
	}

	persons := models.Persons{
		FullName:  input.FullName,
		BirthDate: input.BirthDate,
		Country:   input.Country,
	}

	if err := database.DB.Create(&persons).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка добавления персоны."})
		return
	}

	c.JSON(http.StatusOK, persons)
}

func GetPersons(c *gin.Context) {
	var persons []models.Persons
	if err := database.DB.Find(&persons).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка загрузки персонажей."})
		return
	}
	c.JSON(http.StatusOK, persons)
}

func UpdatePersons(c *gin.Context) {
	var input models.PersonsRequestDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка преобразования."})
		return
	}
	var persons = models.Persons{
		FullName:  input.FullName,
		BirthDate: input.BirthDate,
		Country:   input.Country,
	}
	if err := database.DB.Model(&models.Persons{}).Where("id = ?", c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка изменения"})
		return
	}
	c.JSON(http.StatusOK, persons)

}

func DeletePersons(c *gin.Context) {
	if err := database.DB.Delete(&models.Persons{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при удалении"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Персонаж успешно удалён."})
}
