package models

import (
	"ricr.dev/site-manager/config"
)

type ThirdParty struct {
	config.BaseModel
	Name   string `json:"name" gorm:"size:255;"`
	ApiKey string `json:"api_key" gorm:"size:255;"`
}
