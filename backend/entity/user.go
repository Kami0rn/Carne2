package entity

import (
	// "golang.org/x/text/internal/number"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName   string `valid:"required~Please fill FirstName"`
	LastName    string `valid:"required~Please fill FirstName"`
	PhoneNumber int    `valid:"required~Pease fill phonenumber, matches(^[0]\\d{9}$)~Pease fill valid phonenumber"`
	Email       string `gorm:"uniqueIndex;" valid:"required~Please fill FirstName, email~Please fill email"`
	Password    string `valid:"required~Please fill Lastname"`

	//give fk
	Wallet []Wallet `gorm:"foreignKey:UserID"`
	Result []Result `gorm:"foreignKey:UserID"`
	Wagyu  []Wagyu  `gorm:"foreignKey:UserID"`
}
