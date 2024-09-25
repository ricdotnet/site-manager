package sites

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"ricr.dev/site-manager/api/middlewares"
)

// for the user id we are assuming it will always be present on the authorisation token

func Routes(v1 *echo.Group, db *gorm.DB) {
	api := New(db)

	sites := v1.Group("/site", middlewares.AuthMiddleware())

	sites.GET("/all", api.getAllSites)
	sites.GET("/:id", api.getSite)
	sites.POST("/", api.createSite)
	sites.PATCH("/:id", api.updateSite)
	sites.PATCH("/:id/status", api.updateSiteStatus) // a2ensite / a2dissite
	sites.DELETE("/", api.deleteSite)
}
