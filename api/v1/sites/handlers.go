package sites

import (
	"net/http"
	"os/exec"
	"regexp"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"github.com/ricdotnet/goenvironmental"

	"ricr.dev/site-manager/config"
	"ricr.dev/site-manager/models"
	"ricr.dev/site-manager/utils"
)

type Site = models.Site
type User = models.User

type Response struct {
	config.ApiResponse
	Site   *Site   `json:"site,omitempty"`
	Sites  *[]Site `json:"sites,omitempty"`
	Config string  `json:"config,omitempty"`
}

type RequestBody struct {
	Site   *Site  `json:"site"`
	Config string `json:"config"`
}

type DeleteSites struct {
	Sites *[]uint
}

// TODO: Add pagination
func (api *API) getAllSites(ctx echo.Context) error {
	userCtx := utils.GetTokenClaims(ctx)

	sites, err := api.findAll(userCtx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, config.ApiResponse{
			Code:        http.StatusInternalServerError,
			MessageCode: "sites_fetch_error",
		})
	}

	return ctx.JSON(http.StatusOK, Response{
		ApiResponse: config.ApiResponse{
			Code:        http.StatusOK,
			MessageCode: "sites_fetch_success",
		},
		Sites: sites,
	})
}

func (api *API) getSite(ctx echo.Context) error {
	userCtx := utils.GetTokenClaims(ctx)

	id, _ := strconv.Atoi(ctx.Param("id"))
	site, err := api.findFirst(id, userCtx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "site_not_found")
	}

	nginxDir, _ := goenvironmental.Get("SITES_AVAILABLE_PATH")
	conf, err := api.sitesService.ReadSingle(nginxDir, site.ConfigName)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, &Response{
			ApiResponse: config.ApiResponse{
				Code:    400,
				Message: "We have a database record for that name but could not read the file",
			},
		})
	}

	return ctx.JSON(http.StatusOK, Response{
		ApiResponse: config.ApiResponse{
			Code:    http.StatusOK,
			Message: "site_found",
		},
		Site:   site,
		Config: string(conf),
	})
}

func (api *API) createSite(ctx echo.Context) error {
	log.Info("Entering create a site")

	site := new(Site)
	if err := ctx.Bind(site); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if errorCode := api.validateSite(site); errorCode != nil {
		return ctx.JSON(http.StatusBadRequest, &config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: *errorCode,
		})
	}

	userCtx := utils.GetTokenClaims(ctx)
	site.UserID = userCtx.UserID

	newSite, err := api.insert(site)
	if err != nil {
		log.Warnf("Failed to create a site with the domain %s", site.Domain)

		return ctx.JSON(http.StatusBadRequest, Response{
			ApiResponse: config.ApiResponse{
				Code:        http.StatusBadRequest,
				MessageCode: "site_create_error",
			},
		})
	}

	if err = api.sitesService.WriteSingle(site.ConfigName, ""); err != nil {
		log.Errorf("Failed to write the new site config: %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to write the new site config")
	}

	log.Info("Exiting create a site")

	return ctx.JSON(http.StatusCreated, Response{
		ApiResponse: config.ApiResponse{
			Code:    http.StatusCreated,
			Message: "site_create_success",
		},
		Site: newSite,
	})
}

func (api *API) updateSite(ctx echo.Context) error {
	body := &RequestBody{}
	err := ctx.Bind(body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong when binding the interface to the context")
	}

	oldSite, err := api.findFirst(int(body.Site.ID), utils.GetTokenClaims(ctx))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "This site does not exist")
	}

	newSite, err := api.update(body.Site)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update site")
	}

	if oldSite.ConfigName != newSite.ConfigName {
		err = api.sitesService.UpdateName(oldSite.ConfigName, newSite.ConfigName)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update site")
		}
	}

	err = api.sitesService.WriteSingle(body.Site.ConfigName, body.Config)
	if err != nil {
		println(err.Error())
	}

	return ctx.JSON(http.StatusOK, Response{
		ApiResponse: config.ApiResponse{
			Code:    200,
			Message: "Site updated successfully",
		},
	})
}

func (api *API) updateSiteStatus(ctx echo.Context) error {
	site := &Site{}
	err := ctx.Bind(site)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong when binding the interface to the context")
	}

	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	site.ID = uint(id)
	if err = api.updateEnabled(site); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Unable to update site status")
	}

	cmd := exec.Command("systemctl", "reload", "apache2")
	stdout, _ := cmd.Output()

	log.Info(stdout)

	return ctx.JSON(http.StatusOK, Response{
		ApiResponse: config.ApiResponse{
			Code:    200,
			Message: "Site updated successfully",
		},
	})
}

func (api *API) deleteSite(ctx echo.Context) error {
	log.Info("Entering delete sites handler")

	sites := &DeleteSites{}
	err := ctx.Bind(sites)
	if err != nil {
		log.Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "delete_sites_failed")
	}

	if err = api.delete(sites.Sites); err != nil {
		log.Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "delete_sites_failed")
	}

	log.Info("Exiting delete sites handler")

	return ctx.JSON(http.StatusOK, Response{
		ApiResponse: config.ApiResponse{
			Code:    http.StatusOK,
			Message: "delete_sites_success",
		},
	})
}

func (api *API) validateSite(site *Site) *string {
	domainRegex := "^[A-Za-z0-9-]+(.[A-Za-z0-9-]+)*\\.[A-Za-z]{2,}$"
	configRegex := "^[A-Za-z0-9-]+(.[A-Za-z0-9-]+)*.[A-Za-z]{2,}\\.conf$"

	if site.Domain == "" {
		log.Warn("Tried to add a site with no domain")
		errCode := "site_missing_domain"
		return &errCode
	}

	if match, _ := regexp.MatchString(domainRegex, site.Domain); !match {
		log.Warnf("Tried to add a site with an invalid domain %s", site.Domain)
		errCode := "site_invalid_domain"
		return &errCode
	}

	if site.ConfigName == "" {
		log.Warn("Tried to add a site with no config name")
		errCode := "site_missing_config"
		return &errCode
	}

	if match, _ := regexp.MatchString(configRegex, site.ConfigName); !match {
		log.Warnf("Tried to add a site with an invalid config name %s", site.ConfigName)
		errCode := "site_invalid_config"
		return &errCode
	}

	return nil
}
