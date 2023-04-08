package user

import "ricr.dev/site-manager/config"

type User struct {
	config.BaseModel
	Username  string `json:"username" gorm:"unique;size:255"`
	Password  string `json:"password" gorm:"size:255"`
	Email     string `json:"email" gorm:"size:255"`
	FirstName string `json:"first_name" gorm:"size:255"`
	LastName  string `json:"last_name" gorm:"size:255"`
}
