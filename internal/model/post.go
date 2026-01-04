package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID  uint   `gorm:"not null;index"`
	Title   string `gorm:"not null;size:200"`
	Content string `gorm:"not null;type:text"`
	Views   int    `gorm:"index"`
}

func (Post) TableName() string {
	return "post"
}
