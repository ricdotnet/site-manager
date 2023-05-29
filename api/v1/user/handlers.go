package user

import (
	"github.com/alexedwards/argon2id"
	"github.com/labstack/echo/v4"
	"net/http"
	"regexp"
	"ricr.dev/site-manager/config"
	"ricr.dev/site-manager/models"
	"ricr.dev/site-manager/utils"
)

type User = models.User

type Response struct {
	config.ApiResponse
	ID       uint   `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Token    string `json:"token,omitempty"`
}

func (a *API) auth(ctx echo.Context) error {
	userCtx := utils.GetTokenClaims(ctx)

	return ctx.JSON(http.StatusOK, Response{
		ID:       userCtx.UserID,
		Username: userCtx.Username,
		ApiResponse: config.ApiResponse{
			Code:        http.StatusOK,
			MessageCode: "valid_token",
		},
	})
}

func (a *API) login(ctx echo.Context) error {
	a.logger.Info("Entering /login handler")

	user := new(User)
	_ = ctx.Bind(user)

	username := user.Username
	if username == "" {
		a.logger.Warning("A user tried to login without username")
		return ctx.JSON(http.StatusBadRequest, config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: "missing_username",
		})
	}

	password := user.Password
	if password == "" {
		a.logger.Warningf("User %s tried to login without password", username)
		return ctx.JSON(http.StatusBadRequest, config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: "missing_password",
		})
	}

	result, _ := a.repository.GetOne(username, false)
	if result == nil {
		return ctx.JSON(http.StatusBadRequest, &Response{
			ApiResponse: config.ApiResponse{
				Code:        http.StatusBadRequest,
				MessageCode: "username_not_found",
			},
		})
	}

	correctPassword, _, _ := argon2id.CheckHash(password, result.Password)
	if !correctPassword {
		a.logger.Warningf("User %s tried to login with an incorrect password", username)
		return ctx.JSON(http.StatusBadRequest, &Response{
			ApiResponse: config.ApiResponse{
				Code:        http.StatusBadRequest,
				MessageCode: "incorrect_password",
			},
		})
	}

	token := utils.MakeToken(result)

	a.logger.Info("Exiting /login handler")
	return ctx.JSON(http.StatusOK, Response{
		ID:       result.ID,
		Username: result.Username,
		Token:    token,
		ApiResponse: config.ApiResponse{
			Code:        http.StatusOK,
			MessageCode: "login_success",
		},
	})
}

func (a *API) register(ctx echo.Context) error {
	a.logger.Info("Entering the /register handler")

	user := new(User)
	_ = ctx.Bind(user)

	messageCode := a.registerValidationHelper(user)

	if messageCode != "" {
		a.logger.Info("Exiting the /register handler")
		return ctx.JSON(http.StatusBadRequest, config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: messageCode,
		})
	}

	password, _ := argon2id.CreateHash(user.Password, argon2id.DefaultParams)
	user.Password = password
	a.repository.CreateOne(user)

	a.logger.Info("Exiting the /register handler")
	return ctx.JSON(http.StatusOK, config.ApiResponse{
		Code:        http.StatusOK,
		MessageCode: "register_success",
	})
}

func (a *API) update(ctx echo.Context) error {
	a.logger.Info("Entering the /update handler")

	// do the update stuff

	a.logger.Info("Exiting the /update handler")
	return ctx.JSON(http.StatusNotImplemented, config.ApiResponse{
		Code:        http.StatusNotImplemented,
		MessageCode: "not_implemented",
	})
}

func (a *API) registerValidationHelper(user *User) string {
	usernameRegex := "^[A-Za-z0-9_$]+$"
	emailRegex := "^[A-Za-z0-9._%+-]+@[A-Za-z0-9-.]+\\.[A-Za-z]{2,}$"

	if user.Username == "" {
		a.logger.Warning("A user tried to register without username")
		return "missing_username"
	}

	if user.Email == "" {
		a.logger.Warning("A user tried to register without email address")
		return "missing_email"
	}

	if user.Password == "" {
		a.logger.Warning("A user tried to register without password")
		return "missing_password"
	}

	if len(user.Username) < 5 || len(user.Username) > 30 {
		a.logger.Warningf("Username %s is of invalid length", user.Username)
		return "invalid_username_length"
	}

	if match, _ := regexp.MatchString(usernameRegex, user.Username); !match {
		a.logger.Warningf("Username %s does not pass regex validation", user.Username)
		return "invalid_username"
	}

	if match, _ := regexp.MatchString(emailRegex, user.Email); !match {
		a.logger.Warningf("Email address %s does not pass regex validation", user.Email)
		return "invalid_email"
	}

	if len(user.Password) <= 8 {
		a.logger.Warningf("User %s tried to register with a password of invalid length", user.Username)
		return "invalid_password_length"
	}

	if user.Password != user.PasswordConfirm {
		a.logger.Warningf("User %s tried to register without confirming their password", user.Username)
		return "passwords_not_match"
	}

	existingUsername, _ := a.repository.GetOne(user.Username, false)
	if existingUsername != nil {
		a.logger.Warningf("User %s tried to register with an existing username", user.Username)
		return "username_exists"
	}

	existingEmail, _ := a.repository.GetOne(user.Email, true)
	if existingEmail != nil {
		a.logger.Infof("User %s tried to register with an existing email address", user.Email)
		return "email_exists"
	}

	return ""
}
