package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/magadov/kinopoisk-app/handlers"
)

func SetupMovies(router *gin.Engine) {

	router.POST("/movies/:genre_id", handlers.AddMovie)

	// router.GET("/movies", handlers.GetAllMovies)

	router.GET("/genres/:genre_id/movies", handlers.GetAllMoviesByGenreID)

	router.DELETE("movies/:id", handlers.RemoveMovie)

	router.PATCH("/movies/:id", handlers.UpdateMovie)

	router.GET("/movies", handlers.FilteredByQueryParameters)

}
