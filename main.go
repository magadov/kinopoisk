package main

import (
	"github.com/gin-gonic/gin"
	"github.com/magadov/kinopoisk-app/database"
	"github.com/magadov/kinopoisk-app/routes"
)

func main() {
    database.Init()
	database.Migrate(database.DB)

	router := gin.Default()

	routes.SetupAwards(router)
	routes.SetupPersonAw(router)

	router.Run(":8080")
}
