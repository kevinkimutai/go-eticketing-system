package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Ticket []Ticket
}
