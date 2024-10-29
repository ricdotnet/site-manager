package domains

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"ricr.dev/site-manager/api/middlewares"
)

func Routes(v1 *echo.Group, db *gorm.DB) {
	domainsApi := New(db)

	domains := v1.Group("/domains", middlewares.AuthMiddleware())

	domains.GET("/all", domainsApi.getAllDomains)
	domains.GET("/:domain", domainsApi.getDomain)
	domains.GET("/dns/:domain/:type", domainsApi.getRecords)
	domains.PUT("/dns/:domain/:type", domainsApi.addRecord)
	domains.DELETE("/dns/:domain/:type", domainsApi.deleteRecord)
}
