package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"io"
	"net/http"
	"ricr.dev/site-manager/models"
	"ricr.dev/site-manager/repository"
	"strconv"
)

type (
	DomainsService struct {
		SettingsRepo *repository.SettingsRepo
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

func NewDomainsService(db *gorm.DB) *DomainsService {
	settingsRepo := &repository.SettingsRepo{Db: db}

	return &DomainsService{
		SettingsRepo: settingsRepo,
	}
}

func (d *DomainsService) GetDomains() (error, *Domains) {
	response, _ := domainsHttpApi(d.SettingsRepo, "GET", "domains/search.json", "&no-of-records=20&page-no=1")
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

	return nil, domains
}

func (d *DomainsService) GetDomain(domain string) interface{} {
	unmarshal := func(data []byte) interface{} {
		var result map[string]json.RawMessage
		_ = json.Unmarshal(data, &result)

		return result
	}

	query := fmt.Sprintf("&domain-name=%s&options=%s", domain, "all")
	response, _ := domainsHttpApi(d.SettingsRepo, "GET", "domains/details-by-name.json", query)

	responseBuf, _ := io.ReadAll(response.Body)
	result := unmarshal(responseBuf)

	return result
}

func (d *DomainsService) GetRecords(domain string, recordType string) *Records {
	query := fmt.Sprintf("&domain-name=%s&type=%s&no-of-records=%d&page-no=%d", domain, recordType, 20, 1)
	response, _ := domainsHttpApi(d.SettingsRepo, "GET", "dns/manage/search-records.json", query)
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

	return records
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

func domainsHttpApi(repo *repository.SettingsRepo, method string, url string, query string) (*http.Response, error) {
	apiKey := &models.Settings{}
	referrerId := &models.Settings{}

	_ = repo.GetOne("DOMAINS", apiKey)
	_ = repo.GetOne("DOMAINS_ID", referrerId)

	api := fmt.Sprintf("https://test.httpapi.com/api/%s?auth-userid=%s&api-key=%s%s", url, referrerId.Value, apiKey.Value, query)

	req, _ := http.NewRequest(method, api, nil)

	client := &http.Client{}
	response, err := client.Do(req)

	if response.StatusCode > 299 {
		return nil, errors.New("Request failed with status code: " + response.Status)
	}

	return response, err
}
