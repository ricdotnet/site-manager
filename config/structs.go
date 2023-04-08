package config

import "time"

// list of global structs used in the app
type (
	ApiResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	BaseModel struct {
		ID        uint `gorm:"primaryKey"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
