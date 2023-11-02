package model

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	EventId       int
	TransactionId int
}
