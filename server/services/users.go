package services

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/mailjet/mailjet-apiv3-go"
	"github.com/ricdotnet/goenvironmental"
	"strconv"
)

type UsersService struct {
}

func (s *UsersService) SendMail(to string, code string) error {
	mjApiKey, err := goenvironmental.Get("MJ_API_KEY")
	mjApiSecret, err := goenvironmental.Get("MJ_API_SECRET")
	appUrl, err := goenvironmental.Get("APP_URL")
	mjLoginTmpl, err := goenvironmental.Get("MJ_LOGIN")

	if err != nil {
		log.Errorf("Failed to get ENV variable: %s", err.Error())
		return err
	}

	mjClient := mailjet.NewMailjetClient(mjApiKey, mjApiSecret)

	templateId, _ := strconv.Atoi(mjLoginTmpl)

	messagesInfo := []mailjet.InfoMessagesV31{
		{
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: to,
				},
			},
			TemplateID:       templateId,
			TemplateLanguage: true,
			Variables: map[string]interface{}{
				"loginLink":    fmt.Sprintf("%s/verify-code?code=%s", appUrl, code),
				"loginCode":    code,
				"emailAddress": to,
			},
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	_, err = mjClient.SendMailV31(&messages)

	if err != nil {
		log.Errorf("Failed to send email: %s", err.Error())
		return err
	}

	return nil
}
