package models

import "gorm.io/gorm"

type Tweet struct {
	gorm.Model
	Body      string
	Likes     uint
	ProfileID uint
	HashtagID uint
}
