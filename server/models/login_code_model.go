package models

import "time"

type LoginCode struct {
	Code      string `gorm:"unique"`
	Email     string
	CreatedAt time.Time
}
