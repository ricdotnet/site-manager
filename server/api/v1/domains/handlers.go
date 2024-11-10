package domains

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"ricr.dev/site-manager/config"
	"ricr.dev/site-manager/services"
	"strings"
)

type (
	Response struct {
		config.ApiResponse
		Domains *services.Domains `json:"domains,omitempty"`
		Domain  interface{}       `json:"domain,omitempty"` // not sure if this will be needed so no need to type now
		Records *services.Records `json:"records,omitempty"`
	}

	DNSRecord struct {
		Host  string `json:"host,omitempty"`
		Value string `json:"value,omitempty"`
		TTL   string `json:"ttl,omitempty"`
	}
)

func (d *DomainsAPI) getAllDomains(c echo.Context) error {
	err, domains := d.domainsService.GetDomains()

	if err.Error() == "record not found" {
		return c.JSON(http.StatusNoContent, &Response{
			ApiResponse: config.ApiResponse{
				Code: http.StatusNoContent,
			},
			Domains: &services.Domains{},
		})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &config.ApiResponse{
			Code:        http.StatusInternalServerError,
			MessageCode: "error_getting_domains",
		})
	}

	return c.JSON(http.StatusOK, &Response{
		ApiResponse: config.ApiResponse{
			Code: http.StatusOK,
		},
		Domains: domains,
	})
}

func (d *DomainsAPI) getDomain(c echo.Context) error {
	domainName := c.Param("domain")

	err, domain := d.domainsService.GetDomain(domainName)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &config.ApiResponse{
			Code:        http.StatusInternalServerError,
			MessageCode: "error_getting_domain",
		})
	}

	return c.JSON(http.StatusOK, &Response{
		ApiResponse: config.ApiResponse{
			Code: http.StatusOK,
		},
		Domain: domain,
	})
}

func (d *DomainsAPI) getRecords(c echo.Context) error {
	domainName := c.Param("domain")
	recordType := c.Param("type")

	_, records := d.domainsService.GetRecords(domainName, strings.ToUpper(recordType))

	return c.JSON(http.StatusOK, &Response{
		ApiResponse: config.ApiResponse{
			Code: http.StatusOK,
		},
		Records: records,
	})
}

func (d *DomainsAPI) addRecord(c echo.Context) error {
	domainName := c.Param("domain")
	recordType := c.Param("type")

	body := &DNSRecord{}

	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, &config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: "invalid_request_body",
		})
	}

	var query string

	switch strings.ToUpper(recordType) {
	case "A", "AAAA", "CNAME", "TXT":
		query = fmt.Sprintf("&domain-name=%s&host=%s&value=%s&ttl=%s", domainName, body.Host, body.Value, body.TTL)
	default:
		return c.JSON(http.StatusBadRequest, &config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: "invalid_record_type",
		})
	}

	err, _ := d.domainsService.AddRecord(recordType, query)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &config.ApiResponse{
			Code:        http.StatusInternalServerError,
			MessageCode: "error_adding_record",
		})
	}

	return c.JSON(http.StatusOK, &config.ApiResponse{
		Code:        http.StatusOK,
		MessageCode: "record_added",
	})
}

func (d *DomainsAPI) deleteRecord(c echo.Context) error {
	domainName := c.Param("domain")
	recordType := c.Param("type")

	body := &DNSRecord{}

	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, &config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: "invalid_request_body",
		})
	}

	var query string

	switch strings.ToUpper(recordType) {
	case "A", "AAAA", "CNAME", "TXT":
		query = fmt.Sprintf("&domain-name=%s&host=%s&value=%s&ttl=%s", domainName, body.Host, body.Value, body.TTL)
	default:
		return c.JSON(http.StatusBadRequest, &config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: "invalid_record_type",
		})
	}

	err, _ := d.domainsService.DeleteRecord(recordType, query)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &config.ApiResponse{
			Code:        http.StatusInternalServerError,
			MessageCode: "error_deleting_record",
		})
	}

	return c.JSON(http.StatusOK, &config.ApiResponse{
		Code:        http.StatusOK,
		MessageCode: "record_deleted",
	})
}
