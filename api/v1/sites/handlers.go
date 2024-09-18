package sites

import (
	"net/http"
	"os/exec"
	"regexp"
	"strconv"

	"github.com/labstack/echo/v4"
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
	Vhosts any     `json:"vhosts,omitempty"`
}

type DeleteSites struct {
	Sites *[]uint
}

// all
// get all sites that belong to a specific user
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

// single
// read the contents of the specified file
func (api *API) getSite(ctx echo.Context) error {
	userCtx := utils.GetTokenClaims(ctx)

	id, _ := strconv.Atoi(ctx.Param("id"))
	site, err := api.findFirst(id, userCtx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "site_not_found")
	}

	// TODO: the below code will be removed and vhosts files will be mapped to a site-data table
	// .... using a stub to generate new vhosts when the user updates or adds a new site
	//apacheDir, _ := goenvironmental.Get("APACHE_DIR")
	//vhosts, err := a.sitesService.ReadSingle(apacheDir+"sites-available/", site.ConfigName)
	//if err != nil {
	//	return echo.NewHTTPError(http.StatusBadRequest, &Response{
	//		ApiResponse: config.ApiResponse{
	//			Code:    400,
	//			Message: "We have a database record for that name but could not read the file",
	//		},
	//	})
	//}
	//_vhosts := fmt.Sprintf("%s", vhosts)

	return ctx.JSON(http.StatusOK, Response{
		ApiResponse: config.ApiResponse{
			Code:    http.StatusOK,
			Message: "site_found",
		},
		Site: site,
		//Vhosts: _vhosts,
	})
}

// create
// get the body content and add a new site to the db and vhosts file
func (api *API) createSite(ctx echo.Context) error {
	api.logger.Info("Entering create a site")

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
		api.logger.Warningf("Failed to create a site with the domain %s", site.Domain)
		api.logger.Info("Exiting create a site")
		return ctx.JSON(http.StatusBadRequest, Response{
			ApiResponse: config.ApiResponse{
				Code:        http.StatusBadRequest,
				MessageCode: "site_create_error",
			},
		})
	}

	api.logger.Info("Exiting create a site")
	return ctx.JSON(http.StatusCreated, Response{
		ApiResponse: config.ApiResponse{
			Code:    http.StatusCreated,
			Message: "site_create_success",
		},
		Site: newSite,
	})
}

// update
// endpoint to change an existing site data (not the vhosts file)
func (api *API) updateSite(ctx echo.Context) error {
	site := &Site{}
	err := ctx.Bind(site)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong when binding the interface to the context")
	}

	userCtx := utils.GetTokenClaims(ctx)

	id, _ := strconv.Atoi(ctx.Param("id"))
	site.ID = uint(id)
	if site.ConfigName != "" {
		oldSite, _ := api.findFirst(id, userCtx)
		err = api.sitesService.UpdateName(oldSite.ConfigName, site.ConfigName)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Something went wrong when updating this site")
		}
	}

	updated, _ := api.update(site)

	return ctx.JSON(http.StatusOK, Response{
		ApiResponse: config.ApiResponse{
			Code:    200,
			Message: "Site updated successfully",
		},
		Site: updated,
	})
}

// enable
// endpoint to enable disable a site
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
	api.logger.Info(stdout)

	return ctx.JSON(http.StatusOK, Response{
		ApiResponse: config.ApiResponse{
			Code:    200,
			Message: "Site updated successfully",
		},
	})
}

// delete
// remove an entry from the database as well as the actual vhosts file and disable the site too
func (api *API) deleteSite(ctx echo.Context) error {
	api.logger.Info("Entering delete sites handler")

	sites := &DeleteSites{}
	err := ctx.Bind(sites)
	if err != nil {
		api.logger.Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "delete_sites_failed")
	}

	if err = api.delete(sites.Sites); err != nil {
		api.logger.Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "delete_sites_failed")
	}

	api.logger.Info("Exiting delete sites handler")
	return ctx.JSON(http.StatusOK, Response{
		ApiResponse: config.ApiResponse{
			Code:    http.StatusOK,
			Message: "delete_sites_success",
		},
	})
}

func (a *API) validateSite(site *Site) *string {
	domainRegex := "^[A-Za-z0-9-]+(.[A-Za-z0-9-]+)*\\.[A-Za-z]{2,}$"
	configRegex := "^[A-Za-z0-9-]+(.[A-Za-z0-9-]+)*.[A-Za-z]{2,}\\.conf$"

	if site.Domain == "" {
		a.logger.Warning("Tried to add a site with no domain")
		errCode := "site_missing_domain"
		return &errCode
	}

	if match, _ := regexp.MatchString(domainRegex, site.Domain); !match {
		a.logger.Warningf("Tried to add a site with an invalid domain %s", site.Domain)
		errCode := "site_invalid_domain"
		return &errCode
	}

	if site.ConfigName == "" {
		a.logger.Warning("Tried to add a site with no config name")
		errCode := "site_missing_config"
		return &errCode
	}

	if match, _ := regexp.MatchString(configRegex, site.ConfigName); !match {
		a.logger.Warningf("Tried to add a site with an invalid config name %s", site.ConfigName)
		errCode := "site_invalid_config"
		return &errCode
	}

	return nil
}
