package models

import "ricr.dev/site-manager/config"

type Site struct {
	config.BaseModel
	Domain     string `json:"domain" gorm:"size:255;unique"`
	ConfigName string `json:"config_name" gorm:"size:255;unique"`
	User       uint   `json:"user"`
	Enabled    bool   `json:"enabled" gorm:"default:0"`
	Content    string `json:"content,omitempty" gorm:"-"`
}
