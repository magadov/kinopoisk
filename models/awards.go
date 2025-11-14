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
	Year     uint   `json:"year"`
	Result   string `json:"result"`
}

type MovieAwardsDTO struct {
	MovieID   uint   `json:"movie_id"`
	AwardName string `json:"award_name"`
	Category  string `json:"category"`
	Year      uint   `json:"year"`
	Result    string `json:"result"`
}

type PersonAwardsDTO struct {
	PersonID uint   `json:"person_id"`
	Category string `json:"category"`
	Year     uint   `json:"year"`
	Result   string `json:"result"`
}
