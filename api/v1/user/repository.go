package user

import (
	"github.com/charmbracelet/log"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetOne(value string, isEmail bool) (*User, error) {
	user := new(User)

	if isEmail {
		if err := r.db.First(user, "email = ?", value).Error; err != nil {
			return nil, err
		}
		log.Infof("Found 1 user record with the email %s", user.Email)
	} else {
		if err := r.db.First(user, "username = ?", value).Error; err != nil {
			return nil, err
		}
		log.Infof("Found 1 user record with the username %s", user.Username)
	}

	return user, nil
}

func (r *Repository) CreateOne(user *User) {
	err := r.db.Create(user).Error

	if err != nil {
		log.Errorf("Could not create a new user with the username %s", user.Username)
	}

	log.Infof("Created a new user with the username %s", user.Username)
}

func (r *Repository) UpdateOne(user *User) {
	// query to update the user
}
