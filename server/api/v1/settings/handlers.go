package settings

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ricr.dev/site-manager/config"
	"ricr.dev/site-manager/models"
)

type Setting = models.Settings

func (s *SettingsAPI) getAllApiKeys(ctx echo.Context) error {
	settings := &[]Setting{}
	userCtx := ctx.Get("user").(*config.Session)

	err := s.repository.SettingsRepository.GetAll(settings, userCtx.UserID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, config.ApiResponse{
			Code:        http.StatusInternalServerError,
			MessageCode: "setting_get_error",
		})
	}

	return ctx.JSON(http.StatusOK, settings)
}

func (s *SettingsAPI) getApiKey(ctx echo.Context) error {
	setting := &Setting{}
	err := s.repository.SettingsRepository.GetOne(ctx.Param("key"), setting)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, config.ApiResponse{
			Code:        http.StatusNotFound,
			MessageCode: "setting_not_found",
		})
	}

	return ctx.JSON(http.StatusOK, setting)
}

func (s *SettingsAPI) createOrUpdateApiKey(ctx echo.Context) error {
	body := &Setting{}
	if err := ctx.Bind(body); err != nil {
		return ctx.JSON(http.StatusBadRequest, config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: "setting_missing_payload",
		})
	}

	if body.Key == "" || body.Value == "" {
		return ctx.JSON(http.StatusBadRequest, config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: "setting_invalid_payload",
		})
	}

	userCtx := ctx.Get("user").(*config.Session)
	body.UserID = userCtx.UserID

	setting, err := s.repository.SettingsRepository.CreateOrUpdateOne(body)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, config.ApiResponse{
			Code:        http.StatusInternalServerError,
			MessageCode: "setting_save_error",
		})
	}

	return ctx.JSON(http.StatusAccepted, setting)
}

func (s *SettingsAPI) deleteApiKey(ctx echo.Context) error {
	userCtx := ctx.Get("user").(*config.Session)

	if err := s.repository.SettingsRepository.DeleteOne(ctx.Param("key"), userCtx.UserID); err != nil {
		return ctx.JSON(http.StatusInternalServerError, config.ApiResponse{
			Code:        http.StatusInternalServerError,
			MessageCode: "setting_delete_error",
		})
	}

	return ctx.NoContent(http.StatusAccepted)
}
