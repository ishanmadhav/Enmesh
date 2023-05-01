package models

import "gorm.io/gorm"

type Tweet struct {
	gorm.Model
	Text     string
	Author   Profile
	Likes    uint
	Hashtags []Topic
	Mentions []Profile
}
