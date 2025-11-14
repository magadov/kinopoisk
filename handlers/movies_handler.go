package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/magadov/kinopoisk-app/database"
	"github.com/magadov/kinopoisk-app/models"
)

func AddMovie(ctx *gin.Context) {
	var input models.MoviesRequestDTO

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "некорректный формат JSON",
			"details": err.Error(),
		})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id должно быть числовым значением!"})
		return
	}

	var movie = models.Movies{
		Title:         input.Title,
		OriginalTitle: input.OriginalTitle,
		ReleaseYear:   input.ReleaseYear,
		DurationMin:   input.DurationMin,
		Rating:        input.Rating,
		PgRating:      input.PgRating,
		Country:       input.Country,
		Description:   input.Description,
		GenreID:       uint(id),
	}

	if err := database.DB.Create(&movie).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "не удалось добавить фильм",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, movie)
}

func GetAllMovies(ctx *gin.Context) {
	var movies []models.Movies

	if err := database.DB.Find(&movies).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "не удалось вывести список фильмов",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, movies)
}

func GetAllMoviesByGenreID(ctx *gin.Context) {
	var movies []models.Movies

	if err := database.DB.Model(&models.Movies{}).Where("genre_id = ?", ctx.Param("genre_id")).Find(&movies).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "не удалось вывести список фильмов",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, movies)
}

func RemoveMovie(ctx *gin.Context) {
	if err := database.DB.Unscoped().Delete(&models.Movies{}, ctx.Param("id")).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "не удалось удалить фильм",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "фильм успешно удален!"})
}

func UpdateMovie(ctx *gin.Context) {
	var input models.MoviesUpdateDTO

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(&models.Movies{})

	// if err := database.DB.Model(&models.Movies{}).
	// 	Where("id = ?", ctx.Param("id")).
	// 	Updates(&updated).Error; err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"error":   "не удалось сохранить изменения",
	// 		"details": err.Error(),
	// 	})
	// 	return
	// }

	// ctx.JSON(http.StatusOK, updated)
}

func FilteredByQueryParameters(ctx *gin.Context) {
	var movies []models.Movies
	var query = database.DB.Model(&models.Movies{})

	var requests = map[string]bool{
		"release_year": true,
		"duration_min": true,
		"rating":       true,
		"pg_rating":    true,
		"country":      true,
		"genre_id":     true,
	}

	for key := range ctx.Request.URL.Query() {
		if !requests[key] {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "передан неверный параметр"})
			return
		}
	}

	for key, _ := range requests {
		if request := ctx.Query(key); request != "" {
			query = query.Where(fmt.Sprintf("%s=?", key), request)
		}

	}

	if err := query.Find(&movies).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, movies)
}
