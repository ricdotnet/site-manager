package domains

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"ricr.dev/site-manager/api/middlewares"
)

func Routes(v1 *echo.Group, db *gorm.DB) {
	domainsApi := New(db)

	domains := v1.Group("/domains", middlewares.CookieMiddleware(db))

	domains.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			domainsApi.domainsService.Context = c
			return next(c)
		}
	})

	domains.GET("/all", domainsApi.getAllDomains)
	domains.GET("/:domain", domainsApi.getDomain)
	domains.GET("/dns/:domain/:type", domainsApi.getRecords)
	domains.PUT("/dns/:domain/:type", domainsApi.addRecord)
	domains.DELETE("/dns/:domain/:type", domainsApi.deleteRecord)
}
