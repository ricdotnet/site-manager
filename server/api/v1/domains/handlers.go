package domains

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func (d *DomainsAPI) getAllDomains(c echo.Context) error {
	err, data := d.domainsService.GetDomains()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &Response{
		Code:    http.StatusOK,
		Message: data,
	})
}

func (d *DomainsAPI) getDomain(c echo.Context) error {
	domainName := c.Param("domain")

	domain := d.domainsService.GetDomain(domainName)

	return c.JSON(http.StatusOK, &Response{
		Code: http.StatusOK,
		Message: map[string]interface{}{
			"details": domain,
		},
	})
}

func (d *DomainsAPI) getRecords(c echo.Context) error {
	domainName := c.Param("domain")
	recordType := c.Param("type")

	records := d.domainsService.GetRecords(domainName, strings.ToUpper(recordType))

	return c.JSON(http.StatusOK, &Response{
		Code: http.StatusOK,
		Message: map[string]interface{}{
			"records": records,
		},
	})
}
