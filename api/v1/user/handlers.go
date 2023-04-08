package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ricr.dev/site-manager/config"
)

func (a *API) login(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, config.ApiResponse{
		Code:    501,
		Message: "Endpoint not implemented",
	})
}

func (a *API) register(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, config.ApiResponse{
		Code:    501,
		Message: "Endpoint not implemented",
	})
}

func (a *API) update(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotImplemented, config.ApiResponse{
		Code:    501,
		Message: "Endpoint not implemented",
	})
}
