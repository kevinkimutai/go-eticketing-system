package model

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Venue       string     `json:"venue"`
	Location    string     `json:"location"`
	DateOfEvent *time.Time `json:"date_of_event"`
	VendorId    uint
	UserId      uint
	Ticket      []Ticket
}
