package sites

import (
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

func (r *Repository) GetAll(user int) *[]Site {
	sites := new([]Site)
	r.db.Find(&sites, "user = ?", user)

	return sites
}

func (r *Repository) GetOne(id int) *Site {
	site := &Site{}
	r.db.First(site, "id = ?", id)

	return site
}

func (r *Repository) Create(site *Site) *Site {
	r.db.Create(site)
	return site
}

func (r *Repository) Update(site *Site) *Site {
	return site
}

func (r *Repository) Delete(id int) {
	r.db.Delete(&Site{}, id)
}
