package sites

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"ricr.dev/site-manager/api/middlewares"
)

// for the user id we are assuming it will always be present on the authorisation token

func Routes(v1 *echo.Group, db *gorm.DB) {
	sitesApi := New(db)

	sites := v1.Group("/site", middlewares.CookieMiddleware)

	sites.GET("/all", sitesApi.getAllSites)
	sites.GET("/:id", sitesApi.getSite)
	sites.POST("/", sitesApi.createSite)
	sites.PATCH("/:id", sitesApi.updateSite)
	sites.PATCH("/:id/status", sitesApi.updateSiteStatus) // a2ensite / a2dissite
	sites.DELETE("/", sitesApi.deleteSite)
}
