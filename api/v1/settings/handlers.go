package settings

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"ricr.dev/site-manager/config"
	"ricr.dev/site-manager/models"
)

type Setting = models.Settings

func (api *API) getApiKey(c echo.Context) error {
	setting := &Setting{}
	api.findFirst(setting, c.Param("key"))

	return c.JSON(http.StatusOK, setting)
}

func (api *API) createOrUpdateApiKey(c echo.Context) error {
	body := &Setting{}
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: "setting_missing_payload",
		})
	}

	if body.Key == "" || body.Value == "" {
		return c.JSON(http.StatusBadRequest, config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: "setting_invalid_payload",
		})
	}

	if err := api.insertOrUpdate(body); err != nil {
		return c.JSON(http.StatusInternalServerError, config.ApiResponse{
			Code:        http.StatusInternalServerError,
			MessageCode: "setting_save_error",
		})
	}

	return c.NoContent(http.StatusAccepted)
}

func (api *API) deleteApiKey(c echo.Context) error {
	if err := api.delete(c.Param("key")); err != nil {
		return c.JSON(http.StatusInternalServerError, config.ApiResponse{
			Code:        http.StatusInternalServerError,
			MessageCode: "setting_delete_error",
		})
	}

	return c.NoContent(http.StatusAccepted)
}
