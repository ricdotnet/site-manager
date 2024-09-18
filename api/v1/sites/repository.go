package sites

import (
	"gorm.io/gorm"
	"ricr.dev/site-manager/config"
)

func (api *API) findAll(user *config.JwtCustomClaims) (*[]Site, error) {
	sites := new([]Site)
	if err := api.db.Find(&sites, "user_id = ?", user.UserID).Error; err != nil {
		return nil, err
	}

	api.logger.Infof("Found %d site records for user %s", len(*sites), user.Username)
	return sites, nil
}

func (api *API) findFirst(id int, user *config.JwtCustomClaims) (*Site, error) {
	site := &Site{}
	if err := api.db.First(site, "id = ? AND user_id = ?", id, user.UserID).Error; err != nil {
		return nil, err
	}
	if site != nil {
		api.logger.Infof("Found 1 site record with id %d for user %s", id, user.Username)
	}

	return site, nil
}

func (api *API) insert(site *Site) (*Site, error) {
	if err := api.db.Create(site).Error; err != nil {
		return nil, err
	}
	return site, nil
}

func (api *API) update(site *Site) (*Site, error) {
	updated := &Site{}
	var err error
	if err = api.db.Updates(site).Error; err != nil {
		return nil, err
	}
	if err = api.db.First(updated, "id = ?", site.ID).Error; err != nil {
		return nil, err
	}
	return updated, nil
}

// TODO: maybe merge this with above?
func (api *API) updateEnabled(site *Site) error {
	if err := api.db.Model(site).
		Where("id = ?", site.ID).
		Update("enabled", site.Enabled).Error; err != nil {
		return err
	}

	return nil
}

func (api *API) delete(sites *[]uint) error {
	err := api.db.Transaction(func(tx *gorm.DB) error {
		for _, site := range *sites {
			api.logger.Debugf("Deleting site with id %d", site)

			if err := api.db.Delete(&Site{}, site).Error; err != nil {
				return err
			}
		}

		return nil
	})

	return err
}
