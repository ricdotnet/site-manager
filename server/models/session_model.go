package models

import (
	"ricr.dev/site-manager/config"
	"time"
)

type UserAgent struct {
	Browser        string `json:"name"`
	BrowserVersion string `json:"version"`
	Os             string `json:"os"`
	OsVersion      string `json:"os_version"`
	IsDesktop      bool   `json:"is_desktop"`
	IsMobile       bool   `json:"is_mobile"`
	IsTablet       bool   `json:"is_tablet"`
}

type Session struct {
	config.BaseModel
	Token           string    `json:"token" gorm:"primaryKey" length:"15"`
	ExpiresAt       time.Time `json:"expires_at"`
	UserID          uint      `json:"user_id"`
	UserAgent       string    `json:"user_agent"`
	ParsedUserAgent UserAgent `json:"parsed_user_agent" gorm:"-"`
	ThisDevice      bool      `json:"this_device" gorm:"-"`
	LastActive      time.Time `json:"last_active"`
	IpAddress       string    `json:"ip_address"`
}
