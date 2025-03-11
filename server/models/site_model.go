package models

import (
	"ricr.dev/site-manager/config"
)

type Site struct {
	config.BaseModel
	Domain     string `json:"domain" gorm:"size:255;"`
	ConfigName string `json:"config_name" gorm:"size:255;"`
	Enabled    bool   `json:"enabled" gorm:"default:0"`
	HasSSL     bool   `json:"has_ssl" gorm:"-"`
	UserID     uint   `json:"user_id" gorm:"foreignKey:User"`
	ConfigOnly bool   `json:"config_only" gorm:"-"`
}

type SiteData struct {
	ServerName  string `json:"server_name" gorm:"size:255;not null;"`
	ServerAlias string `json:"server_alias" gorm:"size:255;not null;"`
	Port        uint   `json:"port" gorm:"not null;"`
	Directory   string `json:"directory" gorm:"size:255;not null;"`
	SiteID      uint   `json:"site_id" gorm:"foreignKey:Site"`
	Site        Site
}
