package domains

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"ricr.dev/site-manager/api/middlewares"
)

type Response struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

type DomainsResponse struct {
	Id        json.RawMessage
	Total     int `json:"totalrecords"`
	PageCount int `json:"pagecount"`
}

type DomainData struct {
	OrderId string `json:"orders.orderid"`
	Domain  string `json:"entity.description"`
}

func Routes(v1 *echo.Group, db *gorm.DB) {
	domainsApi := New(db)

	domains := v1.Group("/domains", middlewares.AuthMiddleware())

	domains.GET("/all", domainsApi.getAllDomains)
	domains.GET("/:domain", domainsApi.getDomain)
	domains.GET("/:domain/:type", domainsApi.getRecords)
}
