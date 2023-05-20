package models

import (
	"ricr.dev/site-manager/config"
)

type Site struct {
	config.BaseModel
	Domain     string `json:"domain" gorm:"size:255;unique"`
	ConfigName string `json:"config_name" gorm:"size:255;unique"`
	User       uint   `json:"user"`
	Enabled    bool   `json:"enabled" gorm:"default:0"`
	Content    string `json:"content,omitempty" gorm:"-"`
}

type SiteData struct {
	ServerName  string `json:"server_name" gorm:"size:255;not null;"`
	ServerAlias string `json:"server_alias" gorm:"size:255;not null;"`
	Port        int32  `json:"port" gorm:"not null;"`
	Directory   string `json:"directory" gorm:"size:255;not null;"`
}
