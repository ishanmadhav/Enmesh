package models

import "gorm.io/gorm"

type Hashtag struct {
	gorm.Model
	Name   string
	Tweets []Tweet
}
