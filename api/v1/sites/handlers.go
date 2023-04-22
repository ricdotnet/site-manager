package sites

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/ricdotnet/goenvironmental"
	"net/http"
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
	userSites, err := a.repository.GetAll(2)
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
	site, err := a.repository.GetOne(id, userCtx.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Something went wrong when trying to find the site")
	}

	// if the site doesn't exist it will set the site to 0?
	if site.ID == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "The site you tried to look for does not exist")
	}

	apacheDir, _ := goenvironmental.Get("APACHE_DIR")
	vhosts, err := a.sitesService.ReadSingle(apacheDir+"sites-available/", site.ConfigName)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, &Response{
			ApiResponse: config.ApiResponse{
				Code:    400,
				Message: "We have a database record for that name but could not read the file",
			},
		})
	}
	_vhosts := fmt.Sprintf("%s", vhosts)

	return ctx.JSON(http.StatusOK, Response{
		ApiResponse: config.ApiResponse{
			Code:    200,
			Message: "We found the site you are looking for",
		},
		Site:   site,
		Vhosts: _vhosts,
	})
}

// create
// get the body content and add a new site to the db and vhosts file
func (a *API) create(ctx echo.Context) error {
	site := new(Site)
	if err := ctx.Bind(site); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	create, err := a.repository.Create(site)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, Response{
			ApiResponse: config.ApiResponse{
				Code:    400,
				Message: "Unable to create a new site",
			},
		})
	}

	return ctx.JSON(http.StatusCreated, Response{
		ApiResponse: config.ApiResponse{
			Code:    201,
			Message: "Site created with success",
		},
		Site: create,
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

	id, _ := strconv.Atoi(ctx.Param("id"))
	if site.ConfigName != "" {
		oldSite, _ := a.repository.GetOne(id, 1)
		err = a.sitesService.UpdateName(oldSite.ConfigName, site.ConfigName)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Unable to find a file to rename")
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
func (a *API) enable(ctx echo.Context) error {
	site := &Site{}
	err := ctx.Bind(site)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong when binding the interface to the context")
	}

	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	site.ID = uint(id)
	if err := a.repository.Enable(site); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Unable to update site status")
	}

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

	if err := a.repository.Delete(_id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Unable to delete site")
	}

	return ctx.JSON(http.StatusOK, Response{
		ApiResponse: config.ApiResponse{
			Code:    200,
			Message: "Record deleted from the system",
		},
	})
}
