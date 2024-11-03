package repository

import (
	"github.com/charmbracelet/log"
	"gorm.io/gorm"
	"ricr.dev/site-manager/models"
)

type SessionRepo struct {
	Db *gorm.DB
}

type Session = models.Session

func (repo *SessionRepo) CreateOne(session interface{}) error {
	return repo.Db.Create(session).Error
}

func (repo *SessionRepo) GetOne(token string, items ...interface{}) error {
	log.Info("Will look for a user session")

	session := items[0].(*Session)
	err := repo.Db.Where(&Session{Token: token}).First(session).Error

	if err != nil {
		return err
	}

	log.Infof("Found session for user %d", session.UserId)

	return nil
}

func (repo *SessionRepo) DeleteOne(opts ...interface{}) error {
	token := opts[0].(string)

	return repo.Db.Delete(&Session{}, "`token` = ?", token).Error
}
