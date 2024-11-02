package config

import (
	"github.com/golang-jwt/jwt/v4"
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
		ID        uint      `gorm:"primaryKey"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Session struct {
		UserID uint
		Role   string
	}

	JwtCustomClaims struct {
		UserID   uint   `json:"id"`
		Username string `json:"username"`
		jwt.RegisteredClaims
	}
)
