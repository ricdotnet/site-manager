package sites

import (
	"ricr.dev/site-manager/config"
)

type Site struct {
	config.BaseModel
	Domain     string `json:"domain" gorm:"size:255"`
	ConfigName string `json:"config_name" gorm:"unique;size:255"`
	User       uint   `json:"user"`
	Enabled    bool   `json:"enabled" gorm:"default:0"`
}
