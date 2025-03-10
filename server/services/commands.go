package services

import (
	"encoding/json"
	"github.com/ricdotnet/goenvironmental"
	"gorm.io/gorm"
	"io"
	"net/http"
	"ricr.dev/site-manager/config"
	"ricr.dev/site-manager/models"
	"ricr.dev/site-manager/repository"
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
)

func NewCommandsService(db *gorm.DB) *CommandsService {
	return &CommandsService{
		settingsRepo: &repository.SettingsRepo{Db: db},
	}
}

func commandApiCall(sr *repository.SettingsRepo, userId uint) (*CommandsResponse, error) {
	commandServiceUrl, _ := goenvironmental.Get("APP_COMMAND_SERVICE_URL")

	setting := &models.Settings{}
	err := sr.GetOne("COMMAND_SERVICE", setting, userId)

	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("POST", commandServiceUrl+"certificates", nil)
	req.Header.Set("Authorization", setting.Value)

	client := &http.Client{}
	response, _ := client.Do(req)

	responseBuf, _ := io.ReadAll(response.Body)

	result := &CommandsResponse{}
	_ = json.Unmarshal(responseBuf, &result)

	return result, nil
}

func (cs *CommandsService) GetCertificates(userCtx interface{}) ([]string, error) {
	user := userCtx.(*config.Session)

	result, err := commandApiCall(cs.settingsRepo, user.UserID)

	if err != nil {
		return nil, err
	}

	certificateParts := strings.Split(result.Message, "\n")

	return certificateParts, nil
}
