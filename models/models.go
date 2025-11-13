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


