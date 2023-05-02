package models

import "gorm.io/gorm"

type Tweet struct {
	gorm.Model
	Text      string
	Profile   Profile
	ProfileID uint
	Likes     uint
	Topic     Topic
	TopicID   uint
	Mention   Profile `gorm:"foreignKey:MentionID"`
	MentionID uint
}

type Tweets []Tweet
