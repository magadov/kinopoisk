package models

import (
	"gorm.io/gorm"
)

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
