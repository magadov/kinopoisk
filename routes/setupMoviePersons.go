package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/magadov/kinopoisk-app/handlers"
)

func SetupMoviePersons(router *gin.Engine) {
	{
		groupMoviePersons := router.Group("/movie-persons")

		groupMoviePersons.POST("/:id", handlers.AddMoviePersons)
		groupMoviePersons.GET("/", handlers.GetMoviePersons)
		groupMoviePersons.PATCH("/:id", handlers.UpdateMoviePersons)
		groupMoviePersons.DELETE("/:id", handlers.DeleteMoviePersons)
		groupMoviePersons.GET("/:id", handlers.GetMoviePersonsByID)
	}
}
