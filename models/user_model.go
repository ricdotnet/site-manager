package models

import "ricr.dev/site-manager/config"

type User struct {
	config.BaseModel
	Username        string `json:"username" gorm:"unique;size:255"`
	Password        string `json:"password" gorm:"size:255"`
	PasswordConfirm string `json:"password_confirm" gorm:"-"`
	Email           string `json:"email" gorm:"unique;size:255"`
	FirstName       string `json:"first_name" gorm:"size:255"`
	LastName        string `json:"last_name" gorm:"size:255"`
	Sites           []Site `gorm:"foreignKey:User"`
}
