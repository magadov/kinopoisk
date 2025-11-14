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
