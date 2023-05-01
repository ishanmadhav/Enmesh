package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	Username string
	Name     string
	Email    string
	Age      uint8
	Tweets   []Tweet
}
