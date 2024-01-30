package entity

import (
	// "golang.org/x/text/internal/number"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName   string `valid:"required~Please fill FirstName"`
	LastName    string `valid:"required~Please fill LastName"`
	PhoneNumber int    `valid:"required~Pease fill phonenumber, numeric~Please enter a valid phone number"`
	Email       string `gorm:"uniqueIndex;" valid:"required~Please fill email, email~Email is invalid"`
	Password    string `valid:"required~Please fill Lastname"`

	//give fk
	Wallet []Wallet `gorm:"foreignKey:UserID"`
	Result []Result `gorm:"foreignKey:UserID"`
	Wagyu  []Wagyu  `gorm:"foreignKey:UserID"`
}
