package models

import (
	"ricr.dev/site-manager/config"
	"time"
)

type User struct {
	config.BaseModel
	Username        string    `json:"username" gorm:"unique;size:255"`
	Password        string    `json:"password" gorm:"size:255"`
	PasswordConfirm string    `json:"password_confirm" gorm:"-"`
	Email           string    `json:"email" gorm:"unique;size:255"`
	FirstName       string    `json:"first_name" gorm:"size:255;default:"`
	LastName        string    `json:"last_name" gorm:"size:255;default:"`
	LastLogin       time.Time `json:"last_login" gorm:"default:"`
	Sites           []Site    `gorm:"foreignKey:UserID"`
}
