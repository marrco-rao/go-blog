package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	PostID  uint   `gorm:"not null;index"`
	UserID  uint   `gorm:"not null;index"`
	Content string `gorm:"not null;size:256"`
}

func (Comment) TableName() string {
	return "comment"
}
