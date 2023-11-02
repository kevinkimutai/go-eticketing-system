package model

import "gorm.io/gorm"

type TicketType struct {
	gorm.Model
	Type              string `gorm:"type:ENUM('regular','vvip','vip');default:'regular'"`
	Quantity          uint
	RemainingQuantity uint
}
