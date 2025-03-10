package repository

import (
	"errors"
	"gorm.io/gorm"
)

type (
	GetOne interface {
		GetOne(string, ...interface{}) error
	}

	GetOneByID interface {
		GetOneByID(uint, ...interface{}) error
	}

	GetAll interface {
		GetAll(...interface{}) error
	}

	GetAllByID interface {
		GetAllByID(...interface{}) error
	}

	CreateOne interface {
		CreateOne(interface{}) error
	}

	UpdateOne interface {
		UpdateOne(interface{}) error
	}

	CreateOrUpdateOne interface {
		CreateOrUpdateOne(interface{}) (interface{}, error)
	}

	DeleteOne interface {
		DeleteOne(...interface{}) error
	}

	DeleteOneByID interface {
		DeleteOneByID(uint) error
	}

	DeleteManyByID interface {
		DeleteManyByID([]uint, ...interface{}) error
	}

	UserRepository interface {
		GetOne
		CreateOne
		GetOneByID
		UpdateOne
	}

	SitesRepository interface {
		GetAll
		GetOneByID
		CreateOne
		UpdateOne
		DeleteManyByID
	}

	SettingsRepository interface {
		GetAll
		GetOne
		CreateOrUpdateOne
		DeleteOne
	}

	SessionRepository interface {
		GetAll
		GetOne
		CreateOne
		DeleteOne
		UpdateOne
		DeleteOneByID
	}

	LoginCodeRepository interface {
		GetOne
		CreateOne
		DeleteOne
	}

	Repository struct {
		UserRepository      UserRepository
		SitesRepository     SitesRepository
		SettingsRepository  SettingsRepository
		SessionRepository   SessionRepository
		LoginCodeRepository LoginCodeRepository
	}
)

func notImplemented() error {
	return errors.New("not implemented")
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository:      &UserRepo{Db: db},
		SitesRepository:     &SitesRepo{Db: db},
		SettingsRepository:  &SettingsRepo{Db: db},
		SessionRepository:   &SessionRepo{Db: db},
		LoginCodeRepository: &LoginCodeRepo{Db: db},
	}
}
