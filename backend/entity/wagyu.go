package entity

import (
	// "golang.org/x/text/internal/number"
	"gorm.io/gorm"
)

type Wagyu struct {
	gorm.Model
	Weight float64
	Age    float32

	//FK
	UserID *uint
	User   User `gorm:"foreignKey:UserID"`

	BreedID *uint
	Breed   Breed `gorm:"foreignkey:BreedID"`


	//give fk
	Result []Result `gorm:"foreignKey:WagyuID"`
}
