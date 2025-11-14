package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/magadov/kinopoisk-app/handlers"
)

func SetupPersons(router *gin.Engine) {

	{
		groupPersons := router.Group("/persons")

		groupPersons.POST("/", handlers.AddPersons)
		groupPersons.GET("/", handlers.GetPersons)
		groupPersons.PATCH("/:id", handlers.UpdatePersons)
		groupPersons.DELETE("/:id", handlers.DeletePersons)
	}
}
