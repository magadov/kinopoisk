
package models

import "gorm.io/gorm"


type Movie_awards struct {
	gorm.Model
	MovieID   int    `json:"movie_id"`
	Movie     Movies `json:"movie" gorm:"foreignKey:MovieID"`
	AwardName string `json:"award_name"`
	Category  string `json:"category"`
	Year      int    `json:"year"`
	Result    string `json:"result"`
}

type Person_awards struct {
	gorm.Model
	PersonID int `json: "person_id"`
	Person   Persons `json:"person" gorm:"foreignKey:PersonID"`
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
	Genre         Genres  `gorm:"foreignKey:GenreID"`
}

type Genres struct {
	gorm.Model
	Name string `json:"name" gorm:"unique;not null"`

type Persons struct {
	gorm.Model
	FullName  string `json:"full_name" `
	BirthDate string `json:"birth_date"`
	Country   string `json:"country"`
}

type MoviePersons struct {
	gorm.Model
	MovieID  uint    `json:"movie_id"`
	Movies   Movies  `gorm:"foreignKey:MovieID"`
	PersonID uint    `json:"person_id"`
	Persons  Persons `gorm:"foreignKey:PersonID"`
}
