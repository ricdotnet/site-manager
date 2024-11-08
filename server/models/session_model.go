package models

import "time"

type Session struct {
	Token      string    `json:"token" gorm:"primaryKey" length:"15"`
	ExpiresAt  time.Time `json:"expires_at"`
	UserID     uint      `json:"user_id"`
	UserAgent  string    `json:"user_agent"`
	LastActive time.Time `json:"last_active"`
}
