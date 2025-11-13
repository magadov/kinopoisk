package main

import (
	"github.com/gin-gonic/gin"
	"github.com/magadov/kinopoisk-app/database"
)

func main() {
	db := database.Init()
	database.Migrate(db)

	router := gin.Default()

	router.Run(":5500")
}
