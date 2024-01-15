package entity

import (
	// "golang.org/x/text/internal/number"
	"gorm.io/gorm"
	"time"
)

type Result struct {
	gorm.Model
	Percent float32
	Quality float32
	Date time.Time

	//FK
	UserID *uint
	User   User `gorm:"foreignKey:UserID"`

	WagyuID *uint
	Wagyu  Wagyu `gorm:"foreignKey:WagyuID"`
}
