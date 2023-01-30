package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email         string `gorm:"type:varchar(100);unique_index"`
	Password      []byte `gorm:"type:bytea"`
	Participation []Participation
	Event         Event
	Name          string    `gorm:"type:varchar(100)"`
	Birth         time.Time `gorm:"type:date"`
	Address       string    `gorm:"type:varchar(100)"`
	Gender        string    `gorm:"type:varchar(100)"`
	Bio           string    `gorm:"type:varchar(100)"`
}
