package entity

import (
	// "golang.org/x/text/internal/number"
	"gorm.io/gorm"
)

type Breed struct {
	gorm.Model
	BreedName   string
	BreedDetail string

	//give fk
	Wagyu []Wagyu `gorm:"foreignKey:BreedID"`
}
