package repository

import (
	"github.com/charmbracelet/log"
	"gorm.io/gorm"
	"ricr.dev/site-manager/models"
)

type UserRepo struct {
	Db *gorm.DB
}

func (repo *UserRepo) GetOne(key string, items ...interface{}) error {
	user := items[0].(*models.User)
	isEmail := items[1].(bool)

	if isEmail {
		if err := repo.Db.First(user, "email = ?", key).Error; err != nil {
			return err
		}
		log.Infof("Found 1 user record with the email %s", user.Email)
	} else {
		if err := repo.Db.First(user, "username = ?", key).Error; err != nil {
			return err
		}
		log.Infof("Found 1 user record with the username %s", user.Username)
	}

	return nil
}

func (repo *UserRepo) CreateOne(item interface{}) error {
	user := item.(*models.User)

	err := repo.Db.Create(user).Error

	if err != nil {
		log.Errorf("Could not create a new user with the username %s", user.Username)
		return err
	}

	log.Infof("Created a new user with the username %s", user.Username)

	return nil
}

func (repo *UserRepo) UpdateOne(item interface{}) error {
	// query to update the user
	return notImplemented()
}
