package user

import (
	"github.com/charmbracelet/log"
)


func (api *API) findFirst(value string, isEmail bool) (*User, error) {
	user := new(User)

	if isEmail {
		if err := api.db.First(user, "email = ?", value).Error; err != nil {
			return nil, err
		}
		log.Infof("Found 1 user record with the email %s", user.Email)

	} else {
		if err := api.db.First(user, "username = ?", value).Error; err != nil {
			return nil, err
		}
		log.Infof("Found 1 user record with the username %s", user.Username)

	}

	return user, nil
}

func (api *API) insert(user *User) {
	err := api.db.Create(user).Error

	if err != nil {
		log.Errorf("Could not create a new user with the username %s", user.Username)
	}

	log.Infof("Created a new user with the username %s", user.Username)
}

func (api *API) UpdateOne(user *User) {
	// query to update the user
}
