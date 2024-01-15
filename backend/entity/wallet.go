package entity

import (
	// "golang.org/x/text/internal/number"
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Total float64

	//FK
	UserID *uint
	User   User `gorm:"foreignKey:UserID"`
}
