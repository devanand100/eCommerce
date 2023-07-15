package user

import 	"gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string
	LastName string
	MobileNumber string `gorm:"unique;not null"`
	Email string	`gorm:"unique"`
	Password string  `gorm:"not null"`
}


