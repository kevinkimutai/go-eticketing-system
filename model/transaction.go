package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID  uint
	User    User
	Tickets []Ticket
}
