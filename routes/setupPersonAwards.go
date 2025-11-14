package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/magadov/kinopoisk-app/handlers"
)

func SetupPersonAw(r *gin.Engine){
	{
		groupPersonAw := r.Group("/person_aw")
		
		groupPersonAw.GET("/", handlers.AllPersonAw)
		groupPersonAw.GET("/winners/:year", handlers.WinnersOfYear)
		groupPersonAw.POST("/", handlers.CreatePersonAw)
		groupPersonAw.PATCH("/:id", handlers.UpdatePersonAw)
		groupPersonAw.DELETE("/:id", handlers.DeletePersonAw)
	}
}