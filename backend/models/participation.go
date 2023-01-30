package models

import "gorm.io/gorm"

type Participation struct {
	gorm.Model
	UserID  uint `gorm:"type:integer"`
	EventID uint `gorm:"type:integer"`
}
