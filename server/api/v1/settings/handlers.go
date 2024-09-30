package settings

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"ricr.dev/site-manager/config"
	"ricr.dev/site-manager/models"
)

type Setting = models.Settings

func (s *SettingsAPI) getApiKey(c echo.Context) error {
	setting := &Setting{}
	err := s.repo.GetOne(c.Param("key"), setting)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, config.ApiResponse{
			Code:        http.StatusNotFound,
			MessageCode: "setting_not_found",
		})
	}

	return c.JSON(http.StatusOK, setting)
}

func (s *SettingsAPI) createOrUpdateApiKey(c echo.Context) error {
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

	setting, err := s.repo.CreateOrUpdateOne(body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, config.ApiResponse{
			Code:        http.StatusInternalServerError,
			MessageCode: "setting_save_error",
		})
	}

	return c.JSON(http.StatusAccepted, setting)
}

func (s *SettingsAPI) deleteApiKey(c echo.Context) error {
	if err := s.repo.DeleteOne(c.Param("key")); err != nil {
		return c.JSON(http.StatusInternalServerError, config.ApiResponse{
			Code:        http.StatusInternalServerError,
			MessageCode: "setting_delete_error",
		})
	}

	return c.NoContent(http.StatusAccepted)
}
