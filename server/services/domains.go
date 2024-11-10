package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"io"
	"net/http"
	"ricr.dev/site-manager/config"
	"ricr.dev/site-manager/models"
	"ricr.dev/site-manager/repository"
	"ricr.dev/site-manager/utils"
	"slices"
	"strconv"
	"strings"
	"time"
)

type (
	DomainsService struct {
		SettingsRepo *repository.SettingsRepo
		Context      echo.Context
	}

	Pagination struct {
		Total     int `json:"total_records"`
		PageCount int `json:"page_count"`
	}

	Domains struct {
		Domains []Domain `json:"domains"`
		Pagination
	}

	Domain struct {
		OrderId     string `json:"order_id"`
		Domain      string `json:"name"`
		RenewalDate string `json:"renewal_at"`
		CreatedDate string `json:"created_at"`
	}

	DomainRaw struct {
		OrderId     string `json:"orders.orderid"`
		Domain      string `json:"entity.description"`
		RenewalDate string `json:"orders.endtime"`
		CreatedDate string `json:"orders.creationdt"`
	}

	Records struct {
		Records []Record `json:"records"`
		Pagination
	}

	Record struct {
		Id            int    `json:"record_id"`
		TimeToLiveRaw string `json:"timetolive,omitempty"`
		TimeToLive    string `json:"ttl"`
		Status        string `json:"status"`
		Type          string `json:"type"`
		Host          string `json:"host"`
		Value         string `json:"value"`
	}
)

var recordTypes = map[string]func(string) string{
	"A":     func(action string) string { return fmt.Sprintf("%s-ipv4-record", action) },
	"AAAA":  func(action string) string { return fmt.Sprintf("%s-ipv6-record", action) },
	"CNAME": func(action string) string { return fmt.Sprintf("%s-cname-record", action) },
	"TXT":   func(action string) string { return fmt.Sprintf("%s-txt-record", action) },
}

func NewDomainsService(db *gorm.DB) *DomainsService {
	settingsRepo := &repository.SettingsRepo{Db: db}

	return &DomainsService{
		SettingsRepo: settingsRepo,
	}
}

func (d *DomainsService) GetDomains() (error, *Domains) {
	response, err := domainsHttpApi(d.Context, d.SettingsRepo, "GET", "domains/search.json", "&no-of-records=20&page-no=1", nil)

	if err != nil {
		return err, nil
	}

	domains := &Domains{}

	responseBuf, _ := io.ReadAll(response.Body)

	for _, domainData := range unmarshalGeneric(responseBuf, &domains.Pagination) {
		domain := &Domain{}
		domainRaw := DomainRaw{}
		_ = json.Unmarshal(domainData, &domainRaw)

		domain.Domain = domainRaw.Domain
		domain.OrderId = domainRaw.OrderId
		domain.RenewalDate = domainRaw.RenewalDate
		domain.CreatedDate = domainRaw.CreatedDate

		domains.Domains = append(domains.Domains, *domain)
	}

	slices.SortFunc(domains.Domains, func(i, j Domain) int {
		orderId1, _ := strconv.Atoi(i.OrderId)
		orderId2, _ := strconv.Atoi(j.OrderId)

		return orderId1 - orderId2
	})

	return nil, domains
}

func (d *DomainsService) GetDomain(domain string) (error, interface{}) {
	unmarshal := func(data []byte) interface{} {
		var result map[string]json.RawMessage
		_ = json.Unmarshal(data, &result)

		return result
	}

	query := fmt.Sprintf("&domain-name=%s&options=%s", domain, "all")
	response, err := domainsHttpApi(d.Context, d.SettingsRepo, "GET", "domains/details-by-name.json", query, nil)

	if err != nil {
		return err, nil
	}

	responseBuf, _ := io.ReadAll(response.Body)
	result := unmarshal(responseBuf)

	return nil, result
}

func (d *DomainsService) GetRecords(domain string, recordType string) (error, *Records) {
	query := fmt.Sprintf("&domain-name=%s&type=%s&no-of-records=%d&page-no=%d", domain, recordType, 20, 1)
	response, err := domainsHttpApi(d.Context, d.SettingsRepo, "GET", "dns/manage/search-records.json", query, nil)

	if err != nil {
		return err, nil
	}

	records := &Records{}

	responseBuf, _ := io.ReadAll(response.Body)

	for i, recordData := range unmarshalGeneric(responseBuf, &records.Pagination) {
		record := Record{}
		_ = json.Unmarshal(recordData, &record)

		record.Id = i + 1
		record.TimeToLive = record.TimeToLiveRaw
		record.TimeToLiveRaw = "" // clear the raw value for ignoring it in the response

		records.Records = append(records.Records, record)
	}

	slices.SortFunc(records.Records, func(i, j Record) int {
		return strings.Compare(i.Host, j.Host)
	})

	return nil, records
}

func (d *DomainsService) AddRecord(recordType string, query string) (error, interface{}) {
	url := fmt.Sprintf("dns/manage/%s.json", recordTypes[strings.ToUpper(recordType)]("add"))
	response, err := domainsHttpApi(d.Context, d.SettingsRepo, "POST", url, query, nil)

	if err != nil {
		return err, nil
	}

	return nil, response
}

func (d *DomainsService) DeleteRecord(recordType string, query string) (error, interface{}) {
	url := fmt.Sprintf("dns/manage/%s.json", recordTypes[strings.ToUpper(recordType)]("delete"))
	response, err := domainsHttpApi(d.Context, d.SettingsRepo, "POST", url, query, nil)

	if err != nil {
		return err, nil
	}

	return nil, response
}

func unmarshalGeneric(data []byte, pagination *Pagination) []json.RawMessage {
	var result map[string]json.RawMessage
	_ = json.Unmarshal(data, &result)

	var dataArray []json.RawMessage

	for key, value := range result {
		if key == "recsonpage" {
			var pageCount string
			_ = json.Unmarshal(value, &pageCount)
			pagination.PageCount, _ = strconv.Atoi(pageCount)
		} else if key == "recsindb" {
			var totalCount string
			_ = json.Unmarshal(value, &totalCount)
			pagination.Total, _ = strconv.Atoi(totalCount)
		} else {
			rawJsonData := json.RawMessage{}
			_ = json.Unmarshal(value, &rawJsonData)

			dataArray = append(dataArray, rawJsonData)
		}
	}

	return dataArray
}

func domainsHttpApi(ctx echo.Context, repo *repository.SettingsRepo, method string, url string, query string, payload []byte) (*http.Response, error) {
	session := ctx.Get("user").(*config.Session)

	apiKey := &models.Settings{}
	referrerId := &models.Settings{}

	if err := repo.GetOne("DOMAINS", apiKey, session.UserID); err != nil {
		return nil, err
	}
	if err := repo.GetOne("DOMAINS_ID", referrerId, session.UserID); err != nil {
		return nil, err
	}

	api := fmt.Sprintf("%s%s", utils.DomainsApiUrl(url, referrerId.Value, apiKey.Value), query)

	var payloadReader io.Reader

	if payload != nil {
		payloadReader = strings.NewReader(string(payload))
	}

	log.Infof("Sending request to: %s", api)

	domainsApiRequestStart := time.Now()

	req, _ := http.NewRequest(method, api, payloadReader)

	client := &http.Client{}
	response, err := client.Do(req)

	log.Infof("Request took: %s", time.Since(domainsApiRequestStart))

	if err != nil {
		return nil, err
	}

	if response.StatusCode > 299 {
		return nil, errors.New("Request failed with status code: " + response.Status)
	}

	return response, nil
}
