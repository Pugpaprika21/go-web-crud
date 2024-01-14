package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username      string `gorm:"uniqueIndex;type:varchar(50);not null"`
	Password      string `gorm:"uniqueIndex;type:varchar(50);not null"`
	FirstName     string `gorm:"type:varchar(100)"`
	LastName      string `gorm:"type:varchar(100)"`
	Email         string `gorm:"type:varchar(100)"`
	Address       string `gorm:"type:varchar(255)"`
	PhoneNumber   string `gorm:"type:varchar(20)"`
	Profile       string `gorm:"type:varchar(255)"`
	RememberToken string `gorm:"type:varchar(100)"`
}
