package http

import (
	"github.com/labstack/echo/v4"
	"log"
	"strconv"
)

func (c *controller) find(ctx echo.Context) error {

	req := ctx.QueryParam("item_id")

	id, err := strconv.ParseUint(req, 10, 64)
	if err != nil || id == 0 {
		return echo.NewHTTPError(400, "Bad request")
	}

	status, err := c.uc.FetchItem(ctx.Request().Context(), id)
	if err != nil {
		log.Printf("find handler -> id: %v, error: %v", req, err)
		return echo.NewHTTPError(500, err.Error())
	}

	return ctx.JSON(200, status)
}
