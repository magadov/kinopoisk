package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/magadov/kinopoisk-app/handlers"
)

func SetupGenres(router *gin.Engine) {

	{
		groupGenres := router.Group("/genres")

		groupGenres.POST("/", handlers.AddGenres)
		groupGenres.GET("/", handlers.GetAllGenres)
		groupGenres.PATCH("/:id", handlers.UpdateGenres)
		groupGenres.DELETE("/:id", handlers.Remove)
	}
}
