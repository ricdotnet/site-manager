package repository

import (
	"errors"
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
)

func notImplemented() error {
	return errors.New("not implemented")
}
