package models

import "gorm.io/gorm"


type Movies struct {
	gorm.Model
	Title         string  `json:"title"`
	OriginalTitle string  `json:"original_title"`
	ReleaseYear   int     `json:"release_year"`
	DurationMin   int     `json:"duration_min"`
	Rating        float64 `json:"rating"`
	PgRating      string  `json:"pg_rating"`
	Country       string  `json:"country"`
	Description   string  `json:"description"`
	GenreID       uint    `json:"genre_id"`
}

type Genres struct {
	gorm.Model
	Name string `json:"name" gorm:"unique;not null"`
}

type Persons struct {
	gorm.Model
	FullName  string `json:"full_name" `
	BirthDate string `json:"birth_date"`
	Country   string `json:"country"`
}

type MoviePersons struct {
	gorm.Model
	MovieID  uint   `json:"movie_id"`
	Movies   Movies `gorm:"foreignKey:MovieID"`
	PersonID uint   `json:"person_id"`
}
