package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Category            string `gorm:"type:varchar(256)"`
	Participations      []Participation
	UserID              uint
	Title               string    `gorm:"type:varchar(256)"`
	Detail              string    `gorm:"type:varchar(1024)"`
	StartTime           time.Time `gorm:"type:date"`
	EndTime             time.Time `gorm:"type:date"`
	Place               string    `gorm:"type:varchar(256)"`
	Address             string    `gorm:"type:varchar(256)"`
	ApplicationDeadline time.Time `gorm:"type:timestamp"`
	ImageURL            string    `gorm:"type:varchar(1024)"`
}
