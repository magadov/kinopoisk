package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/magadov/kinopoisk-app/handlers"
)

func SetupAwards(r *gin.Engine){
	{
		groupAwards := r.Group("/awards")

		groupAwards.GET("/", handlers.AllAwards)
		groupAwards.GET("/winners/:year", handlers.WinMovieOfYear)
		// groupAwards.
		groupAwards.POST("/", handlers.CreateAwards)
		groupAwards.PATCH("/:id", handlers.UpdateAward)
		groupAwards.DELETE("/:id", handlers.DeleteAward)
	}
}
