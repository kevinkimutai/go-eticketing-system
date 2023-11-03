package model

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	UserID        uint
	User          User
	Quantity      uint
	TicketTypeID  uint
	TicketType    TicketType
	EventID       uint
	TransactionID uint
}
