package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null;size:128"`
	Password string `gorm:"not null;size:512"`
}

func (User) TableName() string {
	return "user"
}
