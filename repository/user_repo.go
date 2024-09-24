package repository

import (
	"github.com/charmbracelet/log"
	"gorm.io/gorm"
	"ricr.dev/site-manager/models"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (repo *UserRepo) FindFirst(value string, isEmail bool) (*models.User, error) {
	user := new(models.User)

	if isEmail {
		if err := repo.db.First(user, "email = ?", value).Error; err != nil {
			return nil, err
		}
		log.Infof("Found 1 user record with the email %s", user.Email)
	} else {
		if err := repo.db.First(user, "username = ?", value).Error; err != nil {
			return nil, err
		}
		log.Infof("Found 1 user record with the username %s", user.Username)
	}

	return user, nil
}

func (repo *UserRepo) InsertOne(user *models.User) {
	err := repo.db.Create(user).Error

	if err != nil {
		log.Errorf("Could not create a new user with the username %s", user.Username)
	}

	log.Infof("Created a new user with the username %s", user.Username)
}

func (repo *UserRepo) UpdateOne(user *models.User) {
	// query to update the user
}
