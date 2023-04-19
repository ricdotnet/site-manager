package user

import (
	"github.com/op/go-logging"
	"gorm.io/gorm"
)

type Repository struct {
	db     *gorm.DB
	logger *logging.Logger
}

func NewRepository(db *gorm.DB, logger *logging.Logger) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
	}
}

func (r *Repository) GetOne(username string) (*User, error) {

	user := new(User)
	if err := r.db.First(user, "username = ?", username).Error; err != nil {
		return nil, err
	}

	r.logger.Infof("Found 1 user record for username %s", user.Username)

	return user, nil
}

func (r *Repository) CreateOne(user *User) {

	r.db.Create(user)

	r.logger.Infof("Created a new user with the username %s", user.Username)
}

func (r *Repository) UpdateOne(user *User) {
	// query to update the user
}
