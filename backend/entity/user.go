package entity

import (
	// "golang.org/x/text/internal/number"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"uniqueIndex"`
	LastName string
	PhoneNumber int
	Email       string `gorm:"uniqueIndex;"`
	Password string


	//give fk
	Wallet []Wallet `gorm:"foreignKey:UserID"`
	Result []Result `gorm:"foreignKey:UserID"`
	Wagyu  []Wagyu  `gorm:"foreignKey:UserID"`
}
