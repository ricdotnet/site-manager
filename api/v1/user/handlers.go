package user

import (
	"github.com/alexedwards/argon2id"
	"github.com/labstack/echo/v4"
	"net/http"
	"ricr.dev/site-manager/config"
	"ricr.dev/site-manager/models"
	"ricr.dev/site-manager/utils"
)

type User = models.User

type Response struct {
	config.ApiResponse
	Token string `json:"token,omitempty"`
}

func (a *API) login(ctx echo.Context) error {
	a.logger.Info("Entering /login handler")

	user := new(User)
	_ = ctx.Bind(user)

	username := user.Username
	if username == "" {
		a.logger.Warning("A user tried to login without username")
		return ctx.JSON(http.StatusBadRequest, config.ApiResponse{
			Code:    400,
			Message: "Username cannot be missing",
		})
	}

	password := user.Password
	if password == "" {
		a.logger.Warningf("User %s tried to login without password", username)
		return ctx.JSON(http.StatusBadRequest, config.ApiResponse{
			Code:    400,
			Message: "Password cannot be missing",
		})
	}

	result, _ := a.repository.GetOne(username, false)
	if result == nil {
		return ctx.JSON(http.StatusNotFound, config.ApiResponse{
			Code:    404,
			Message: "A user with the details provided does not exist",
		})
	}

	correctPassword, _, _ := argon2id.CheckHash(password, result.Password)
	if !correctPassword {
		a.logger.Warningf("User %s tried to login with an incorrect password", username)
		return ctx.JSON(http.StatusBadRequest, config.ApiResponse{
			Code:    400,
			Message: "Incorrect password used",
		})
	}

	token := utils.MakeToken(result)

	a.logger.Info("Exiting /login handler")
	return ctx.JSON(http.StatusOK, Response{
		Token: token,
		ApiResponse: config.ApiResponse{
			Code:    200,
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
			Code:    400,
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
		Code:    200,
		Message: "Registered with success",
	})
}

func (a *API) update(ctx echo.Context) error {
	a.logger.Info("Entering the /update handler")

	// do the update stuff

	a.logger.Info("Exiting the /update handler")
	return ctx.JSON(http.StatusNotImplemented, config.ApiResponse{
		Code:    501,
		Message: "Endpoint not implemented",
	})
}
