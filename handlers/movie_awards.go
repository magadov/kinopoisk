package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/magadov/kinopoisk-app/database"
	"github.com/magadov/kinopoisk-app/models"
)

func CreateAwards(c *gin.Context) {
	var awardInp models.MovieAwardsDTO

	if err := c.ShouldBindJSON(&awardInp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	awards := models.MovieAwards{
		MovieID:   awardInp.MovieID,
		AwardName: awardInp.AwardName,
		Category:  awardInp.Category,
		Year:      awardInp.Year,
		Result:    awardInp.Result,
	}

	if res := database.DB.Create(&awards); res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ошибка создания записи в БД"})
		return
	}

	c.JSON(http.StatusOK, awards)
}

func AllAwards(c *gin.Context) {
	var award []models.MovieAwards

	if res := database.DB.Find(&award); res.Error != nil {
		c.JSON(http.StatusBadRequest, res.Error.Error())
		return
	}

	c.JSON(http.StatusOK, award)
}

func UpdateAward(c *gin.Context) {
	var awardUpd models.MovieAwardsDTO
	id := c.Param("id")
	if err := c.ShouldBindJSON(&awardUpd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Model(&models.MovieAwards{}).Where("id = ?", id).Updates(&awardUpd)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nothing updated"})
		return
	}

	var updated models.MovieAwards

	if res := database.DB.First(&updated, id); res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)

}

func DeleteAward(c *gin.Context) {
	id := c.Param("id")
	var award models.MovieAwards

	if res := database.DB.First(&award, id); res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error.Error()})
		return
	}

	if res := database.DB.Delete(&models.MovieAwards{}, id); res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, "deleted")
}

func WinMovieOfYear(c *gin.Context) {
	var winMovies []models.MovieAwards

	if res := database.DB.Where("year = ? AND result = ?", c.Param("year"), "win").Find(&winMovies); res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error})
		return
	}

	c.JSON(http.StatusOK, winMovies)

}
