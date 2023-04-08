package user

import "gorm.io/gorm"

type API struct {
	repository *Repository
}

func New(db *gorm.DB) *API {
	return &API{
		repository: NewRepository(db),
	}
}
