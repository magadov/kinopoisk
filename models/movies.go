package models

import "gorm.io/gorm"

type Movies struct {
	gorm.Model
	Title         string
	OriginalTitle string
	ReleaseYear   int
	DurationMin   int
	Rating        float64
	PgRating      string
	Country       string
	Description   string
	GenreID       uint
}

type MoviesRequestDTO struct {
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

type MoviesUpdateDTO struct {
	Title         *string  `json:"title"`
	OriginalTitle *string  `json:"original_title"`
	ReleaseYear   *int     `json:"release_year"`
	DurationMin   *int     `json:"duration_min"`
	Rating        *float64 `json:"rating"`
	PgRating      *string  `json:"pg_rating"`
	Country       *string  `json:"country"`
	Description   *string  `json:"description"`
}
