package models

import "time"

type Session struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	UserId    uint      `json:"user_id"`
	User      User
}
