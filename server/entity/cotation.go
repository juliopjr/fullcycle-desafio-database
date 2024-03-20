package entity

import "gorm.io/gorm"

type Quotation struct {
	gorm.Model
	Bid string `gorm:"not null"`
}
