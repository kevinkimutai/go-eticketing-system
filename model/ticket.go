package model

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	UserId        uint
	Quantity      uint
	TicketType    TicketType
	EventId       uint
	TransactionId uint
}
