package movies

import (
	"net/http"

	"github.com/go-pg/pg"
    "github.com/labstack/echo"
    "github.com/sjl2/goyagi/pkg/application"
    "github.com/sjl2/goyagi/pkg/model"
)

type handler struct {
    app application.App
}

func (h *handler) listHandler(c echo.Context) error {
    var movies []*model.Movie

    err := h.app.DB.
        Model(&movies).
        Order("id DESC").
        Select()
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, movies)
}

func (h *handler) retrieveHandler(c echo.Context) error {
    id := c.Param("id")

    var movie model.Movie

    err := h.app.DB.Model(&movie).Where("id = ?", id).First()
    if err != nil {
        if err == pg.ErrNoRows {
            return echo.NewHTTPError(http.StatusNotFound, "movie not found")
        }
        return err
    }

    return c.JSON(http.StatusOK, movie)
}
