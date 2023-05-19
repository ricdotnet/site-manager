package config

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// list of global structs used in the app
type (
	ApiResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	BaseModel struct {
		ID        uint      `gorm:"primaryKey"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	JwtCustomClaims struct {
		Id       int    `json:"id"`
		Username string `json:"username"`
		jwt.RegisteredClaims
	}
)
