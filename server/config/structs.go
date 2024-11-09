package config

import (
	"time"
)

// list of global structs used in the app
type (
	ApiResponse struct {
		Code        int    `json:"code"`
		Message     string `json:"message,omitempty"`
		MessageCode string `json:"message_code,omitempty"`
	}

	BaseModel struct {
		ID        uint      `json:"id" gorm:"primaryKey"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Session struct {
		UserID   uint
		Username string
		Email    string
		Role     string
	}
)
