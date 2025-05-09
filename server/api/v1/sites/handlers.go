package sites

import (
	"net/http"
	"regexp"
	"ricr.dev/site-manager/services"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"github.com/ricdotnet/goenvironmental"

	"ricr.dev/site-manager/config"
	"ricr.dev/site-manager/models"
)

type Site = models.Site
type User = models.User

type Response struct {
	config.ApiResponse
	Site         *Site                  `json:"site,omitempty"`
	Sites        *[]Site                `json:"sites,omitempty"`
	Config       string                 `json:"config,omitempty"`
	Certificates []services.Certificate `json:"certificates,omitempty"`
}

type RequestBody struct {
	Site   *Site  `json:"site"`
	Config string `json:"config"`
}

type DeleteSites struct {
	Sites []uint `json:"sites"`
}

func (s *SitesAPI) getAllSites(ctx echo.Context) error {
	sites := &[]Site{}
	userCtx := ctx.Get("user")

	err := s.repository.SitesRepository.GetAll(sites, userCtx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &config.ApiResponse{
			Code:        http.StatusInternalServerError,
			MessageCode: "sites_fetch_error",
		})
	}

	_ = s.sitesService.FindFileOnly(sites)

	return ctx.JSON(http.StatusOK, &Response{
		ApiResponse: config.ApiResponse{
			Code:        http.StatusOK,
			MessageCode: "sites_fetch_success",
		},
		Sites: sites,
	})
}

func (s *SitesAPI) getSite(ctx echo.Context) error {
	site := &Site{}
	userCtx := ctx.Get("user")

	id, _ := strconv.Atoi(ctx.Param("id"))
	err := s.repository.SitesRepository.GetOneByID(uint(id), site, userCtx)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, &config.ApiResponse{
			Code:        http.StatusNotFound,
			MessageCode: "site_not_found",
		})
	}

	nginxDir, _ := goenvironmental.Get("NGINX_PATH")
	conf, err := s.sitesService.ReadSingle(nginxDir, site.ConfigName)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, &config.ApiResponse{
			Code:        400,
			MessageCode: "database_record_exists_but_file_not_found",
		})
	}

	return ctx.JSON(http.StatusOK, &Response{
		ApiResponse: config.ApiResponse{
			Code:    http.StatusOK,
			Message: "site_found",
		},
		Site:   site,
		Config: string(conf),
	})
}

func (s *SitesAPI) createSite(ctx echo.Context) error {
	log.Info("Entering create a site")

	site := new(Site)
	if err := ctx.Bind(site); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if errorCode := s.validateSite(site); errorCode != "" {
		return ctx.JSON(http.StatusBadRequest, &config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: errorCode,
		})
	}

	userCtx := ctx.Get("user").(*config.Session)
	site.UserID = userCtx.UserID

	transaction := s.db.Begin()

	err := transaction.Create(site).Error
	if err != nil {
		log.Warnf("Failed to create a site with the domain %s", site.Domain)

		transaction.Rollback()

		return ctx.JSON(http.StatusBadRequest, &config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: "site_create_error",
		})
	}

	if !site.ConfigOnly {
		err = s.sitesService.WriteSingle(site.ConfigName, "")
		if err != nil {
			log.Errorf("Failed to write the new site config: %s", err.Error())
			transaction.Rollback()
			return echo.NewHTTPError(http.StatusBadRequest, &config.ApiResponse{
				Code:        http.StatusBadRequest,
				MessageCode: "failed_to_write_file_config",
			})
		}
	}

	transaction.Commit()
	site.ConfigOnly = false

	log.Info("Exiting create a site")

	return ctx.JSON(http.StatusCreated, &Response{
		ApiResponse: config.ApiResponse{
			Code:    http.StatusCreated,
			Message: "site_create_success",
		},
		Site: site,
	})
}

func (s *SitesAPI) updateSite(ctx echo.Context) error {
	body := &RequestBody{}
	err := ctx.Bind(body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong when binding the interface to the context")
	}

	oldSite := &Site{}

	err = s.repository.SitesRepository.GetOneByID(body.Site.ID, oldSite, ctx.Get("user").(*config.Session))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, &config.ApiResponse{
			Code:        http.StatusNotFound,
			MessageCode: "site_not_found",
		})
	}

	transaction := s.db.Begin()
	err = transaction.Updates(body.Site).Error

	if err != nil {
		transaction.Rollback()
		return echo.NewHTTPError(http.StatusBadRequest, &config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: "update_site_failed",
		})
	}

	if oldSite.ConfigName != body.Site.ConfigName {
		err = s.sitesService.UpdateName(oldSite.ConfigName, body.Site.ConfigName)
		if err != nil {
			transaction.Rollback()
			return echo.NewHTTPError(http.StatusBadRequest, &config.ApiResponse{
				Code:        http.StatusBadRequest,
				MessageCode: "site_config_rename_failed",
			})
		}
	}

	err = s.sitesService.WriteSingle(body.Site.ConfigName, body.Config)
	if err != nil {
		transaction.Rollback()
		return echo.NewHTTPError(http.StatusBadRequest, &config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: "site_config_update_failed",
		})
	}

	transaction.Commit()

	return ctx.JSON(http.StatusOK, &config.ApiResponse{
		Code:        200,
		MessageCode: "site_updated",
	})
}

func (s *SitesAPI) deleteSite(ctx echo.Context) error {
	log.Info("Entering delete sites handler")

	sites := &DeleteSites{}
	err := ctx.Bind(sites)
	if err != nil {
		log.Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "delete_sites_failed")
	}

	if len(sites.Sites) == 0 {
		return ctx.JSON(http.StatusOK, &config.ApiResponse{
			Code:        http.StatusOK,
			MessageCode: "nothing_to_delete",
		})
	}

	deleted := &[]models.Site{}
	err = s.repository.SitesRepository.DeleteManyByID(sites.Sites, deleted)

	for _, site := range *deleted {
		_ = s.sitesService.DeleteSingle(site.ConfigName)
		log.Infof("Deleted site %s", site.ConfigName)
	}

	if err != nil {
		log.Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, &config.ApiResponse{
			Code:        http.StatusBadRequest,
			MessageCode: "delete_sites_failed",
		})
	}

	log.Info("Exiting delete sites handler")

	return ctx.JSON(http.StatusOK, &config.ApiResponse{
		Code:        http.StatusOK,
		MessageCode: "delete_sites_success",
	})
}

func (s *SitesAPI) validateSite(site *Site) string {
	domainRegex := "^[A-Za-z0-9-]+(.[A-Za-z0-9-]+)*\\.[A-Za-z]{2,}$"
	configRegex := "^[A-Za-z0-9-]+(.[A-Za-z0-9-]+)*.[A-Za-z]{2,}\\.conf$"

	if site.Domain == "" {
		log.Warn("Tried to add a site with no domain")
		return "site_missing_domain"
	}

	if match, _ := regexp.MatchString(domainRegex, site.Domain); !match {
		log.Warnf("Tried to add a site with an invalid domain %s", site.Domain)
		return "site_invalid_domain"
	}

	if site.ConfigName == "" {
		log.Warn("Tried to add a site with no config name")
		return "site_missing_config"
	}

	if match, _ := regexp.MatchString(configRegex, site.ConfigName); !match {
		log.Warnf("Tried to add a site with an invalid config name %s", site.ConfigName)
		return "site_invalid_config"
	}

	return ""
}

func (s *SitesAPI) reloadNginx(ctx echo.Context) error {
	userCtx := ctx.Get("user")

	err := s.commandsService.ReloadNginx(userCtx)
	if err != nil {
		log.Error(err.Error())

		return echo.NewHTTPError(http.StatusInternalServerError, &config.ApiResponse{
			Code:        http.StatusInternalServerError,
			MessageCode: "nginx_reload_failed",
			Message:     "Failed to reload nginx",
		})
	}

	return ctx.JSON(http.StatusOK, &config.ApiResponse{
		Code:        http.StatusOK,
		MessageCode: "nginx_reload_success",
		Message:     "Successfully reloaded nginx",
	})
}

func (s *SitesAPI) getCertificates(ctx echo.Context) error {
	userCtx := ctx.Get("user")

	certificates, err := s.commandsService.GetCertificates(userCtx)
	if err != nil {
		log.Error(err.Error())

		return echo.NewHTTPError(http.StatusInternalServerError, &config.ApiResponse{
			Code:        http.StatusInternalServerError,
			MessageCode: "failed_to_get_certificates",
			Message:     "Failed to get certificates",
		})
	}

	return ctx.JSON(http.StatusOK, &Response{
		ApiResponse: config.ApiResponse{
			Code:        http.StatusOK,
			MessageCode: "sites_fetch_success",
		},
		Certificates: certificates,
	})
}
