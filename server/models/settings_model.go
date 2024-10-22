package models

import (
	"ricr.dev/site-manager/config"
)

type Settings struct {
	config.BaseModel
	Key      string `json:"key" gorm:"size:255;unique"`
	Value    string `json:"value" gorm:"size:255;not null"`
	IsApiKey bool   `json:"is_api_key" gorm:"default:false"`
	UserID   uint   `json:"user_id" gorm:"not null"`
}
