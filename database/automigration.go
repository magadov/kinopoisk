package database

import (
	"github.com/magadov/kinopoisk-app/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.Genres{}, &models.MoviePersons{}, &models.MovieAwards{}, &models.Movies{},
		&models.PersonAwards{}, &models.Persons{}); err != nil {
		return err
	}
	return nil
}
