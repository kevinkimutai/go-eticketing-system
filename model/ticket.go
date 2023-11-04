package model

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	UserID       uint
	User         User `json:"user"`
	Quantity     uint `json:"quantity"`
	TicketTypeID uint

	EventID       uint
	TransactionID uint
}
