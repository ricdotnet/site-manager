package user

import (
	"github.com/alexedwards/argon2id"
	"github.com/golang-jwt/jwt/v4"
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
	Username string `json:"username,omitempty"`
	Token    string `json:"token,omitempty"`
}

func (a *API) auth(ctx echo.Context) error {

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.JwtCustomClaims)

	return ctx.JSON(http.StatusOK, Response{
		Username: claims.Username,
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
		Username: username,
		Token:    token,
		ApiResponse: config.ApiResponse{
			Code:        http.StatusOK,
			MessageCode: "login_success",
		},
	})
}

func (a *API) register(ctx echo.Context) error {
	a.logger.Info("Entering the /register handler")

	usernameRegex := "^[A-Za-z0-9_$]+$"
	emailRegex := "^[A-Za-z0-9._%+-]+@[A-Za-z0-9-.]+\\.[A-Za-z]{2,}$"

	user := new(User)
	_ = ctx.Bind(user)

	if user.Username == "" {
		return ctx.JSON(http.StatusBadRequest, config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: "missing_username",
		})
	}

	if user.Email == "" {
		return ctx.JSON(http.StatusBadRequest, config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: "missing_email",
		})
	}

	if user.Password == "" {
		return ctx.JSON(http.StatusBadRequest, config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: "missing_password",
		})
	}

	if len(user.Username) < 5 || len(user.Username) > 30 {
		return ctx.JSON(http.StatusBadRequest, config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: "invalid_username_length",
		})
	}

	if match, _ := regexp.MatchString(usernameRegex, user.Username); !match {
		a.logger.Warningf("Username %s does not pass regex validation", user.Username)
		return ctx.JSON(http.StatusBadRequest, config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: "invalid_username",
		})
	}

	if match, _ := regexp.MatchString(emailRegex, user.Email); !match {
		a.logger.Warningf("Email address %s dot not pass regex validation", user.Email)
		return ctx.JSON(http.StatusBadRequest, config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: "invalid_email",
		})
	}

	if len(user.Password) <= 8 {
		return ctx.JSON(http.StatusBadRequest, config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: "invalid_password_length",
		})
	}

	if user.Password != user.PasswordConfirm {
		return ctx.JSON(http.StatusBadRequest, config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: "passwords_not_match",
		})
	}

	existingUsername, _ := a.repository.GetOne(user.Username, false)
	if existingUsername != nil {
		a.logger.Infof("A user with the username %s already exists", user.Username)
		return ctx.JSON(http.StatusBadRequest, config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: "username_exists",
		})
	}

	existingEmail, _ := a.repository.GetOne(user.Email, true)
	if existingEmail != nil {
		a.logger.Infof("A user with the email %s already exists", user.Email)
		return ctx.JSON(http.StatusBadRequest, config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: "email_exists",
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
