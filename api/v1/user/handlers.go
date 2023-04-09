package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ricr.dev/site-manager/config"
)

func (a *API) login(ctx echo.Context) error {
	a.logger.Info("Entering /login handler")

	// do the login stuff

	a.logger.Info("Exiting /login handler")
	return ctx.JSON(http.StatusNotImplemented, config.ApiResponse{
		Code:    501,
		Message: "Endpoint not implemented",
	})
}

func (a *API) register(ctx echo.Context) error {
	a.logger.Info("Entering the /register handler")

	// do the register stuff

	a.logger.Info("Exiting the /register handler")
	return ctx.JSON(http.StatusNotImplemented, config.ApiResponse{
		Code:    501,
		Message: "Endpoint not implemented",
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
