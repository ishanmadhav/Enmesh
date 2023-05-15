package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	Name     string
	Username string
	Age      uint
	Tweets   []Tweet
}
