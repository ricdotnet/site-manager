package sites

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/ricdotnet/goenvironmental"
	"net/http"
	"ricr.dev/site-manager/config"
	"ricr.dev/site-manager/services/sites"
	"strconv"
)

type Response struct {
	config.ApiResponse
	Site   *Site   `json:"site,omitempty"`
	Sites  *[]Site `json:"sites,omitempty"`
	Vhosts any     `json:"vhosts,omitempty"`
}

// all
// get all sites that belong to a specific user
func (a *API) all(ctx echo.Context) error {
	userSites := a.repository.GetAll(2)

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
	id, _ := strconv.Atoi(ctx.Param("id"))
	site := a.repository.GetOne(id)

	// if the site doesn't exist it will set the site to 0?
	if site.ID == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "The site you tried to look for does not exist")
	}

	vhosts, err := sites.ReadSingle(goenvironmental.Get("APACHE_DIR")+"sites-available/", site.ConfigName)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
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

	a.repository.Create(site)

	return ctx.JSON(http.StatusCreated, Response{
		ApiResponse: config.ApiResponse{
			Code:    201,
			Message: "Site created with success",
		},
	})
}

// update
// endpoint to change an existing site data (not the vhosts file)
func (a *API) update(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, Response{
		ApiResponse: config.ApiResponse{
			Code:    200,
			Message: "Endpoint not implemented yet",
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

	a.repository.Delete(_id)

	return ctx.JSON(http.StatusOK, Response{
		ApiResponse: config.ApiResponse{
			Code:    200,
			Message: "Record deleted from the system",
		},
	})
}
