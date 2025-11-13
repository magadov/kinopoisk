package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/magadov/kinopoisk-app/database"
	"github.com/magadov/kinopoisk-app/models"
)

func AddMoviePersons(c *gin.Context) {

	var input models.MoviePersonsRequestDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка преобразования данных."})
		return
	}

	var id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID фильма"})
		return
	}

	if input.PersonID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PersonID не может быть пустым!"})
		return
	}

	var moviePersons = models.MoviePersons{
		MovieID:  uint(id),
		PersonID: input.PersonID,
	}

	if err := database.DB.Create(&moviePersons).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при добавлении."})
		return
	}

	if err := database.DB.
		Preload("Movies").
		First(&moviePersons, moviePersons.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при загрузке фильма."})
		return
	}
	c.JSON(http.StatusOK, moviePersons)
}

func GetMoviePersons(c *gin.Context) {
	var moviePersons []models.MoviePersons

	if err := database.DB.Preload("Movies").
		Find(&moviePersons).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось загрузить данные."})
		return
	}
	c.JSON(http.StatusOK, moviePersons)
}

func UpdateMoviePersons(c *gin.Context) {
	var input models.MoviePersonsRequestDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка преобразования данных."})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID фильма"})
	}

	var moviePersons models.MoviePersons

	if err := database.DB.Where("movie_id = ?", uint(id)).First(&moviePersons).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Запись с таким ID не найдена!"})
		return
	}

	moviePersons.ID = input.PersonID

	if err := database.DB.Save(&moviePersons).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при обновлении!"})
		return
	}

	if err := database.DB.Preload("Movies").
		First(&moviePersons, moviePersons.ID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при загрузке фильма!"})
		return
	}

	c.JSON(http.StatusOK, moviePersons)

}

func DeleteMoviePersons(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID!"})
		return
	}
	var moviePersons models.MoviePersons
	if err := database.DB.First(&moviePersons, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Запись не найдена"})
		return
	}

	if err := database.DB.Unscoped().Delete(&moviePersons).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при удалении!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Запись успешно удалена!"})

}

func GetMoviePersonsByID(c *gin.Context) {
	var moviePersons []models.MoviePersons
	if err := database.DB.Preload("Movies").
		Where("person_id = ?", c.Param("id")).
		Find(&moviePersons).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при загрузке данных."})
		return
	}

	if len(moviePersons) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Персонажа с таким ID нет!"})
		return
	}

	c.JSON(http.StatusOK, moviePersons)

}
