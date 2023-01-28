package http

import (
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"gitlab.ozon.dev/bagdatov/homework-3/service-create-order/internal/usecase"
)

type controller struct {
	uc usecase.Order
}

func NewController(uc usecase.Order) *echo.Echo {
	e := echo.New()

	p := prometheus.NewPrometheus("service-create-order", nil)
	p.Use(e)
	c := &controller{uc: uc}

	orders := e.Group("/api/v1/orders")
	orders.POST("/create", c.create)
	orders.POST("/cancel", c.cancel)

	return e
}
