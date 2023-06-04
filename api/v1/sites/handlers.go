package sites

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os/exec"
	"ricr.dev/site-manager/config"
	"ricr.dev/site-manager/models"
	"ricr.dev/site-manager/utils"
	"strconv"
)

type Site = models.Site
type User = models.User

type Response struct {
	config.ApiResponse
	Site   *Site   `json:"site,omitempty"`
	Sites  *[]Site `json:"sites,omitempty"`
	Vhosts any     `json:"vhosts,omitempty"`
}

// all
// get all sites that belong to a specific user
func (a *API) all(ctx echo.Context) error {
	userCtx := utils.GetTokenClaims(ctx)

	userSites, err := a.repository.GetAll(userCtx.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong when trying to find all sites")
	}

	return ctx.JSON(http.StatusOK, Response{
		ApiResponse: config.ApiResponse{
			Code:    200,
			Message: "Here are all your sites",
		},
		Sites: userSites,
	})
}

// single
// read the contents of the specified file
func (a *API) single(ctx echo.Context) error {
	userCtx := utils.GetTokenClaims(ctx)

	id, _ := strconv.Atoi(ctx.Param("id"))
	site, err := a.repository.GetOne(id, userCtx.UserID)
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
func (a *API) create(ctx echo.Context) error {
	a.logger.Info("Entering create a site")

	userCtx := utils.GetTokenClaims(ctx)

	site := new(Site)
	if err := ctx.Bind(site); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	site.UserID = userCtx.UserID

	newSite, err := a.repository.Create(site)
	if err != nil {
		a.logger.Warningf("Failed to create a site with the domain %s", site.Domain)
		a.logger.Info("Exiting create a site")
		return ctx.JSON(http.StatusBadRequest, Response{
			ApiResponse: config.ApiResponse{
				Code:        http.StatusBadRequest,
				MessageCode: "site_create_error",
			},
		})
	}

	a.logger.Info("Exiting create a site")
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
func (a *API) update(ctx echo.Context) error {
	site := &Site{}
	err := ctx.Bind(site)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong when binding the interface to the context")
	}

	userCtx := utils.GetTokenClaims(ctx)

	id, _ := strconv.Atoi(ctx.Param("id"))
	site.ID = uint(id)
	if site.ConfigName != "" {
		oldSite, _ := a.repository.GetOne(id, userCtx.UserID)
		err = a.sitesService.UpdateName(oldSite.ConfigName, site.ConfigName)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Something went wrong when updating this site")
		}
	}

	updated, _ := a.repository.Update(site)

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
func (a *API) status(ctx echo.Context) error {
	site := &Site{}
	err := ctx.Bind(site)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong when binding the interface to the context")
	}

	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	site.ID = uint(id)
	if err = a.repository.Enable(site); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Unable to update site status")
	}

	cmd := exec.Command("systemctl", "reload", "apache2")
	stdout, _ := cmd.Output()
	a.logger.Info(stdout)

	return ctx.JSON(http.StatusOK, Response{
		ApiResponse: config.ApiResponse{
			Code:    200,
			Message: "Site updated successfully",
		},
	})
}

// delete
// remove an entry from the database as well as the actual vhosts file and disable the site too
func (a *API) delete(ctx echo.Context) error {
	id := ctx.Param("id")

	_id, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "The id must be of type integer")
	}

	if err = a.repository.Delete(_id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Unable to delete site")
	}

	return ctx.JSON(http.StatusOK, Response{
		ApiResponse: config.ApiResponse{
			Code:    200,
			Message: "Record deleted from the system",
		},
	})
}
