package models

import "gorm.io/gorm"

type Genres struct {
	gorm.Model
	Name string `json:"name"`
}

type GenresRequestDTO struct {
	Name string
}

type GenresUpdateDTO struct {
	Name string
}
