package model

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name        string     `json:"name"`
	Venue       string     `json:"venue"`
	Location    string     `json:"location"`
	DateOfEvent *time.Time `json:"date_of_event"`
	Price       uint       `json:"price"`
	BannerUrl   string     `json:"banner_url"`
	VendorID    uint
	Tickets     []Ticket
	TicketType  []TicketType `json:"ticket_type"`
}
