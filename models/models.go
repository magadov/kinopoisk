package models

import "gorm.io/gorm"

type MovieAwards struct {
	gorm.Model
	MovieID   uint   `json:"movie_id"`
	AwardName string `json:"award_name"`
	Category  string `json:"category"`
	Year      uint   `json:"year"`
	Result    string `json:"result"`
}

type PersonAwards struct {
	gorm.Model
	PersonID uint   `json:"person_id"`
	Category string `json:"category"`
	Year     string `json:"year"`
	Result   string `json:"result"`
}

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
