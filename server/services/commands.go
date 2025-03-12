package services

import (
	"encoding/json"
	"errors"
	"github.com/ricdotnet/goenvironmental"
	"gorm.io/gorm"
	"io"
	"net/http"
	"regexp"
	"ricr.dev/site-manager/config"
	"ricr.dev/site-manager/models"
	"ricr.dev/site-manager/repository"
	"strconv"
	"strings"
)

type (
	CommandsService struct {
		settingsRepo *repository.SettingsRepo
	}

	CommandsResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Failed  bool   `json:"failed"`
	}

	Certificate struct {
		Name       string `json:"name"`
		ExpiryDate string `json:"expiry_date"`
		ExpiryDays int    `json:"expiry_days"`
		Domains    string `json:"domains"`
	}
)

func NewCommandsService(db *gorm.DB) *CommandsService {
	return &CommandsService{
		settingsRepo: &repository.SettingsRepo{Db: db},
	}
}

func commandApiCall(sr *repository.SettingsRepo, endpoint string, userId uint) (*CommandsResponse, error) {
	commandServiceUrl, _ := goenvironmental.Get("APP_COMMAND_SERVICE_URL")

	setting := &models.Settings{}
	err := sr.GetOne("COMMAND_SERVICE", setting, userId)

	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("POST", commandServiceUrl+endpoint, nil)
	req.Header.Set("Authorization", setting.Value)

	client := &http.Client{}
	response, _ := client.Do(req)

	responseBuf, _ := io.ReadAll(response.Body)

	result := &CommandsResponse{}
	_ = json.Unmarshal(responseBuf, &result)

	return result, nil
}

func (cs *CommandsService) ReloadNginx(userCtx interface{}) error {
	user := userCtx.(*config.Session)

	result, err := commandApiCall(cs.settingsRepo, "reload-nginx", user.UserID)

	if err != nil {
		return err
	}

	if result.Failed {
		return errors.New("failed to reload nginx")
	}

	return nil
}

func (cs *CommandsService) GetCertificates(userCtx interface{}) ([]Certificate, error) {
	user := userCtx.(*config.Session)

	result, err := commandApiCall(cs.settingsRepo, "certificates", user.UserID)
	certificates := make([]Certificate, 0)

	if err != nil {
		return certificates, err
	}

	certificateParts := strings.Split(result.Message, "\n")

	certificateHeader := regexp.MustCompile("^.+Certificate Name: .+")
	certificateDomains := regexp.MustCompile("^.+Domains: .+")
	certificateExpiry := regexp.MustCompile("^.+Expiry Date: .+")
	certificateExpiryDate := regexp.MustCompile("\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}\\+00:00")
	certificateExpiryDays := regexp.MustCompile("(\\((VALID|INVALID): .+)")
	certificateMatchDigits := regexp.MustCompile("\\d{2}")

	for i, certificate := range certificateParts {
		if certificateHeader.MatchString(certificate) {
			cert := &Certificate{}
			name := strings.Split(certificate, " Name: ")[1]

			cert.Name = name

			for j := 1; j < 7; j++ {
				if certificateDomains.MatchString(certificateParts[i+j]) {
					domains := strings.Split(certificateParts[i+j], ": ")[1]
					cert.Domains = domains
				} else if certificateExpiry.MatchString(certificateParts[i+j]) {
					expiryDateString := strings.Split(certificateParts[i+j], "Expiry Date: ")[1]
					expiryDate := certificateExpiryDate.FindStringSubmatch(expiryDateString)
					expiryDays := certificateExpiryDays.FindStringSubmatch(expiryDateString)

					cert.ExpiryDate = expiryDate[0]

					if strings.Contains(expiryDays[0], "INVALID") {
						cert.ExpiryDays = -1
						continue
					}

					days, _ := strconv.Atoi(certificateMatchDigits.FindStringSubmatch(expiryDays[0])[0])
					cert.ExpiryDays = days
				}
			}

			certificates = append(certificates, *cert)
			continue
		}
	}

	return certificates, nil
}
