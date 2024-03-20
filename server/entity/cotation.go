package entity

import "gorm.io/gorm"

type Cotation struct {
	gorm.Model
	Bid string `gorm:"not null"`
}
