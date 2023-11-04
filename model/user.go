package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName    string        `json:"first_name" gorm:"not null"`
	LastName     string        `json:"last_name" gorm:"not null"`
	Email        string        `json:"email" gorm:"unique"`
	Role         string        `gorm:"type:ENUM('admin', 'user');default:'user'" json:"role"`
	Password     string        `json:"password" gorm:"not null"`
	Tickets      []Ticket      `gorm:"foreignKey:UserID" json:"tickets"`
	Transactions []Transaction `gorm:"foreignKey:UserID" json:"transactions"`
	Events       []Event       `gorm:"foreignKey:VendorID" json:"events"`
}
