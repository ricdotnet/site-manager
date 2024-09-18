package models

type Settings struct {
	Key   string `json:"key" gorm:"size:255;unique"`
	Value string `json:"value" gorm:"size:255;"`
}
