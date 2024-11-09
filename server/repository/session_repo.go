package repository

import (
	"github.com/charmbracelet/log"
	"gorm.io/gorm"
	"ricr.dev/site-manager/config"
	"ricr.dev/site-manager/models"
)

type SessionRepo struct {
	Db *gorm.DB
}

type Session = models.Session

func (repo *SessionRepo) CreateOne(session interface{}) error {
	return repo.Db.Create(session).Error
}

func (repo *SessionRepo) UpdateOne(session interface{}) error {
	return repo.Db.Save(session).Error
}

func (repo *SessionRepo) GetOne(token string, items ...interface{}) error {
	log.Info("Will look for a user session")

	session := items[0].(*Session)
	err := repo.Db.Where(&Session{Token: token}).First(session).Error

	if err != nil {
		return err
	}

	log.Infof("Found session for userid %d", session.UserID)

	return nil
}

func (repo *SessionRepo) GetAll(items ...interface{}) error {
	sessions := items[0].(*[]Session)
	session := items[1].(*config.Session)

	return repo.Db.Where("`user_id` = ?", session.UserID).Order("created_at DESC").Find(sessions).Error
}

func (repo *SessionRepo) DeleteOne(opts ...interface{}) error {
	token := opts[0].(string)

	return repo.Db.Delete(&Session{}, "`token` = ?", token).Error
}

func (repo *SessionRepo) DeleteOneByID(id uint) error {
	return repo.Db.Delete(&Session{}, "`id` = ?", id).Error
}
