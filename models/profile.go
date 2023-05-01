package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	Name  string
	Email string
	Age   uint8
}
