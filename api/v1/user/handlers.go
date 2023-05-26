package user

import (
	"github.com/alexedwards/argon2id"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"ricr.dev/site-manager/config"
	"ricr.dev/site-manager/models"
	"ricr.dev/site-manager/utils"
)

type User = models.User

type Response struct {
	config.ApiResponse
	Username    string `json:"username,omitempty"`
	Token       string `json:"token,omitempty"`
	MessageCode string `json:"message_code,omitempty"`
}

func (a *API) auth(ctx echo.Context) error {

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.JwtCustomClaims)

	return ctx.JSON(http.StatusOK, Response{
		Username: claims.Username,
		ApiResponse: config.ApiResponse{
			Code:    http.StatusOK,
			Message: "A valid token is being used",
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
			Code:    http.StatusBadRequest,
			Message: "Username cannot be missing",
		})
	}

	password := user.Password
	if password == "" {
		a.logger.Warningf("User %s tried to login without password", username)
		return ctx.JSON(http.StatusBadRequest, config.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Password cannot be missing",
		})
	}

	result, _ := a.repository.GetOne(username, false)
	if result == nil {
		return ctx.JSON(http.StatusBadRequest, &Response{
			ApiResponse: config.ApiResponse{
				Code:    http.StatusBadRequest,
				Message: "A user with the details provided does not exist",
			},
			MessageCode: "username_not_found",
		})
	}

	correctPassword, _, _ := argon2id.CheckHash(password, result.Password)
	if !correctPassword {
		a.logger.Warningf("User %s tried to login with an incorrect password", username)
		return ctx.JSON(http.StatusBadRequest, &Response{
			ApiResponse: config.ApiResponse{
				Code:    http.StatusBadRequest,
				Message: "Incorrect password used",
			},
			MessageCode: "incorrect_password",
		})
	}

	token := utils.MakeToken(result)

	a.logger.Info("Exiting /login handler")
	return ctx.JSON(http.StatusOK, Response{
		Username: username,
		Token:    token,
		ApiResponse: config.ApiResponse{
			Code:    http.StatusOK,
			Message: "Logged in with success",
		},
	})
}

func (a *API) register(ctx echo.Context) error {
	a.logger.Info("Entering the /register handler")

	user := new(User)
	_ = ctx.Bind(user)

	if user.Password != user.PasswordConfirm {
		return ctx.JSON(http.StatusBadRequest, config.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "The passwords do not match",
		})
	}

	existingUsername, _ := a.repository.GetOne(user.Username, false)
	if existingUsername != nil {
		a.logger.Infof("A user with the username %s already exists", user.Username)
		return ctx.JSON(http.StatusBadRequest, config.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "A user with that username already exists",
		})
	}

	existingEmail, _ := a.repository.GetOne(user.Email, true)
	if existingEmail != nil {
		a.logger.Infof("A user with the email %s already exists", user.Email)
		return ctx.JSON(http.StatusBadRequest, config.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "A user with that email already exists",
		})
	}

	password, _ := argon2id.CreateHash(user.Password, argon2id.DefaultParams)
	user.Password = password
	a.repository.CreateOne(user)

	a.logger.Info("Exiting the /register handler")
	return ctx.JSON(http.StatusOK, config.ApiResponse{
		Code:    http.StatusOK,
		Message: "Registered with success",
	})
}

func (a *API) update(ctx echo.Context) error {
	a.logger.Info("Entering the /update handler")

	// do the update stuff

	a.logger.Info("Exiting the /update handler")
	return ctx.JSON(http.StatusNotImplemented, config.ApiResponse{
		Code:    http.StatusNotImplemented,
		Message: "Endpoint not implemented",
	})
}
