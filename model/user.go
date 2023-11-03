package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName    string        `json:"first_name" gorm:"not null"`
	LastName     string        `json:"last_name" gorm:"not null"`
	Email        string        `json:"email" gorm:"unique"`
	Role         string        `gorm:"type:ENUM('admin', 'user');default:'user'"`
	Password     string        `json:"password" gorm:"not null"`
	Tickets      []Ticket      `gorm:"foreignKey:UserID"`
	Transactions []Transaction `gorm:"foreignKey:UserID"`
	Events       []Event       `gorm:"foreignKey:UserID"`
}
