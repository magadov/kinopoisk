package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/magadov/kinopoisk-app/database"
	"github.com/magadov/kinopoisk-app/models"
)

func AddGenres(ctx *gin.Context) {
	var input models.GenresRequestDTO

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "некорректный формат JSON"})
		return
	}

	var genre = models.Genres{
		Name: input.Name,
	}

	if err := database.DB.Create(&genre).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось создать запись"})
		return
	}

	ctx.JSON(http.StatusOK, genre)
}

func GetAllGenres(ctx *gin.Context) {
	var genres []models.Genres

	if err := database.DB.Find(&genres).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось вывести жанры"})
		return
	}

	ctx.JSON(http.StatusOK, genres)
}

func UpdateGenres(ctx *gin.Context) {
	var input models.GenresUpdateDTO

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error":   "некорректный формат JSON",
			"details": err.Error(),
		})
		return
	}

	var updated = models.Genres{
		Name: input.Name,
	}

	if err := database.DB.Model(&models.Genres{}).
		Where("id = ?", ctx.Param("id")).
		Update("Name", input.Name).
		Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "не удалось сохранить изменения",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, updated)
}

func Remove(ctx *gin.Context) {
	if err := database.DB.Unscoped().Delete(&models.Genres{}, ctx.Param("id")).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "не удалось удалить жанр",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "жанр удален!"})
}
