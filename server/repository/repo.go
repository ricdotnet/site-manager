package repository

import "errors"

type GetOne interface {
	GetOne(string, ...interface{}) error
}

type GetOneByID interface {
	GetOneByID(uint, ...interface{}) error
}

type GetAll interface {
	GetAll(...interface{}) error
}

type GetAllByID interface {
	GetAllByID(...interface{}) error
}

type CreateOne interface {
	CreateOne(interface{}) error
}

type UpdateOne interface {
	UpdateOne(interface{}) error
}

type CreateOrUpdateOne interface {
	CreateOrUpdateOne(interface{}) (interface{}, error)
}

type DeleteOne interface {
	DeleteOne(string) error
}

type DeleteManyByID interface {
	DeleteManyByID([]uint) error
}

type UserRepository interface {
	GetOne
	CreateOne
}

type SitesRepository interface {
	GetAll
	GetOneByID
	CreateOne
	UpdateOne
	DeleteManyByID
}

type SettingsRepository interface {
	GetOne
	CreateOrUpdateOne
	DeleteOne
}

func notImplemented() error {
	return errors.New("not implemented")
}
