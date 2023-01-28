package http

import (
	"github.com/labstack/echo/v4"
	"log"
	"strconv"
)

func (c *controller) cancel(ctx echo.Context) error {

	req := ctx.QueryParam("order_id")

	id, err := strconv.ParseUint(req, 10, 64)
	if err != nil || id == 0 {
		return echo.NewHTTPError(400, "Bad request")
	}

	if err := c.uc.Cancel(ctx.Request().Context(), id); err != nil {
		log.Printf("cancel handler -> id: %v, error: %v", req, err)
		return echo.NewHTTPError(500, err.Error())
	}

	return ctx.String(200, "cancellation request submitted")
}
