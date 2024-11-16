package user

import (
	"fmt"
	"github.com/ricdotnet/goenvironmental"
	"net/http"
	"net/smtp"
	"regexp"
	"ricr.dev/site-manager/utils"
	"strconv"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"

	"ricr.dev/site-manager/config"
	"ricr.dev/site-manager/models"
)

type (
	Response struct {
		config.ApiResponse
		ID             uint              `json:"id,omitempty"`
		Username       string            `json:"username,omitempty"`
		Email          string            `json:"email,omitempty"`
		ActiveSessions *[]models.Session `json:"active_sessions,omitempty"`
	}

	SignIn struct {
		Email string `json:"email"`
	}

	Authenticate struct {
		Code string `json:"code"`
	}
)

func (u *UserAPI) authUser(ctx echo.Context) error {
	session := ctx.Get("user").(*config.Session)

	user := &models.User{}
	_ = u.repo.GetOneByID(session.UserID, user)

	return ctx.JSON(http.StatusOK, Response{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		ApiResponse: config.ApiResponse{
			Code:        http.StatusOK,
			MessageCode: "valid_token",
		},
	})
}

func (u *UserAPI) signIn(ctx echo.Context) error {
	log.Info("Sending magic link email")

	smtpUser, err := goenvironmental.Get("SMTP_USER")
	smtpPassword, err := goenvironmental.Get("SMTP_PASSWORD")
	smtpFrom, err := goenvironmental.Get("SMTP_FROM")
	smtpUrl, err := goenvironmental.Get("SMTP_URL")
	smtpPort, err := goenvironmental.Get("SMTP_PORT")

	if err != nil {
		log.Error(err.Error())
		return ctx.JSON(http.StatusBadRequest, &config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: "smtp_invalid_env",
		})
	}

	body := &SignIn{}
	_ = ctx.Bind(body)

	if body.Email == "" {
		return ctx.JSON(http.StatusBadRequest, &config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: "email_not_present",
		})
	}

	user := &models.User{}
	_ = u.repo.GetOne(body.Email, user, true)

	if user.Email == "" {
		return ctx.JSON(http.StatusNotFound, &config.ApiResponse{
			Code:        http.StatusNotFound,
			MessageCode: "email_not_found",
		})
	}

	code := utils.MakeToken()

	loginCode := &models.LoginCode{
		Code:      code,
		Email:     user.Email,
		CreatedAt: time.Now(),
	}
	_ = u.loginCodeRepo.CreateOne(loginCode)

	to := []string{
		body.Email,
	}

	message := []byte("Subject: Login code for Site-Manager!\r\n\r\n" +
		fmt.Sprintf("Use this code to `%s` login into Site-Manager.\r\n", code))

	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpUrl)
	err = smtp.SendMail(fmt.Sprintf("%s:%s", smtpUrl, smtpPort), auth, smtpFrom, to, message)
	if err != nil {
		log.Errorf("Failed to send email: %s", err.Error())

		return ctx.JSON(http.StatusBadRequest, &config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: "failed_to_send_email",
		})
	}

	return ctx.JSON(http.StatusOK, &config.ApiResponse{
		Code:        http.StatusOK,
		MessageCode: "email_sent",
	})
}

func (u *UserAPI) verifyCode(ctx echo.Context) error {
	body := &Authenticate{}
	_ = ctx.Bind(body)

	loginCode := &models.LoginCode{}
	_ = u.loginCodeRepo.GetOne(body.Code, loginCode)

	if loginCode.CreatedAt.Add(5 * time.Minute).Before(time.Now()) {
		return ctx.JSON(http.StatusForbidden, &config.ApiResponse{
			Code:        http.StatusForbidden,
			MessageCode: "login_code_expired",
		})
	}

	_ = u.loginCodeRepo.DeleteOne(loginCode.Code)

	user := &models.User{}
	err := u.repo.GetOne(loginCode.Email, user, true)
	if err != nil {
		log.Warnf("User %s does not exist", loginCode.Email)

		return ctx.JSON(http.StatusBadRequest, &Response{
			ApiResponse: config.ApiResponse{
				Code:        http.StatusBadRequest,
				MessageCode: "email_not_found",
			},
		})
	}

	token := utils.MakeToken()

	userAgent := ctx.Request().Header.Get("User-Agent")

	session := &models.Session{}
	session.Token = token
	session.UserID = user.ID
	session.UserAgent = userAgent
	session.LastActive = time.Now()
	session.IpAddress = ctx.RealIP()

	expiresAt := time.Now()
	expiresAt = expiresAt.AddDate(0, 0, 30)

	session.ExpiresAt = expiresAt
	_ = u.sessionRepo.CreateOne(session)

	cookieName, _ := goenvironmental.Get("COOKIE_NAME")
	cookieDomain, _ := goenvironmental.Get("COOKIE_DOMAIN")

	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    token,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		Domain:   cookieDomain,
		Expires:  expiresAt,
	}

	ctx.SetCookie(cookie)

	user.LastLogin = time.Now()

	_ = u.repo.UpdateOne(user)

	log.Info("Exiting /login handler")

	return ctx.JSON(http.StatusOK, Response{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		ApiResponse: config.ApiResponse{
			Code:        http.StatusOK,
			MessageCode: "login_success",
		},
	})
}

func (u *UserAPI) registerUser(ctx echo.Context) error {
	log.Info("Entering the /register handler")

	user := new(models.User)
	_ = ctx.Bind(user)

	messageCode := u.registerValidationHelper(user)

	if messageCode != "" {
		log.Info("Exiting the /register handler")

		return ctx.JSON(http.StatusBadRequest, config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: messageCode,
		})
	}

	password, _ := argon2id.CreateHash(user.Password, argon2id.DefaultParams)
	user.Password = password
	_ = u.repo.CreateOne(user)

	log.Info("Exiting the /register handler")

	return ctx.JSON(http.StatusOK, config.ApiResponse{
		Code:        http.StatusOK,
		MessageCode: "register_success",
	})
}

func (u *UserAPI) updateUser(ctx echo.Context) error {
	log.Info("Entering the /update handler")

	// do the update stuff

	log.Info("Exiting the /update handler")

	return ctx.JSON(http.StatusNotImplemented, config.ApiResponse{
		Code:        http.StatusNotImplemented,
		MessageCode: "not_implemented",
	})
}

func (u *UserAPI) getSessions(ctx echo.Context) error {
	log.Info("Entering the /sessions handler")

	session := ctx.Get("user").(*config.Session)

	activeSessions := &[]models.Session{}
	_ = u.sessionRepo.GetAll(activeSessions, session)

	newActiveSessions := make([]models.Session, 0)

	for _, s := range *activeSessions {
		userAgent := &models.UserAgent{}
		utils.ParseUserAgent(s.UserAgent, userAgent)

		s.ParsedUserAgent = *userAgent
		s.ThisDevice = ctx.Get("token") == s.Token

		newActiveSessions = append(newActiveSessions, s)
	}

	log.Info("Exiting the /sessions handler")

	return ctx.JSON(http.StatusOK, &Response{
		ApiResponse: config.ApiResponse{
			Code:        http.StatusOK,
			MessageCode: "found_sessions",
		},
		ActiveSessions: &newActiveSessions,
	})
}

func (u *UserAPI) deleteSession(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	_ = u.sessionRepo.DeleteOneByID(uint(id))

	return ctx.JSON(http.StatusOK, config.ApiResponse{
		Code:        http.StatusOK,
		MessageCode: "session_deleted",
	})
}

func (u *UserAPI) logout(ctx echo.Context) error {
	log.Info("Entering the /logout handler")

	_ = u.sessionRepo.DeleteOne(ctx.Get("token"))
	ctx.SetCookie(utils.MakeEmptyCookie())

	return ctx.JSON(http.StatusOK, config.ApiResponse{
		Code:        http.StatusOK,
		MessageCode: "logout_success",
	})
}

func (u *UserAPI) registerValidationHelper(user *models.User) string {
	usernameRegex := "^[A-Za-z0-9_$]+$"
	emailRegex := "^[A-Za-z0-9._%+-]+@[A-Za-z0-9-.]+\\.[A-Za-z]{2,}$"

	if user.Username == "" {
		log.Warn("A user tried to register without username")

		return "missing_username"
	}

	if user.Email == "" {
		log.Warn("A user tried to register without email address")

		return "missing_email"
	}

	if user.Password == "" {
		log.Warn("A user tried to register without password")

		return "missing_password"
	}

	if len(user.Username) < 5 || len(user.Username) > 30 {
		log.Warnf("Username %s is invalid: %s", user.Username, "INVALID_LENGTH")

		return "invalid_username_length"
	}

	if match, _ := regexp.MatchString(usernameRegex, user.Username); !match {
		log.Warnf("Username %s is invalid", user.Username)

		return "invalid_username"
	}

	if match, _ := regexp.MatchString(emailRegex, user.Email); !match {
		log.Warnf("Email address %s is invalid", user.Email)

		return "invalid_email"
	}

	if len(user.Password) <= 8 {
		log.Warnf("User %s entered an invalid password: %s", user.Username, "INVALID_LENGTH")

		return "invalid_password_length"
	}

	if user.Password != user.PasswordConfirm {
		log.Warnf("User %s tried to register without matching passwords", user.Username)

		return "passwords_not_match"
	}

	_userUsername := &models.User{}
	_ = u.repo.GetOne(user.Username, _userUsername, false)
	if _userUsername.Username != "" {
		log.Warnf("Username %s is already registered", user.Username)

		return "username_exists"
	}

	_userEmail := &models.User{}
	_ = u.repo.GetOne(user.Email, _userEmail, true)
	if _userEmail.Email != "" {
		log.Warnf("Email %s is already registered", user.Email)

		return "email_exists"
	}

	return ""
}
