package http

import (
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"gitlab.ozon.dev/bagdatov/homework-3/service-monitoring-order/internal/usecase"
)

type controller struct {
	uc usecase.Monitoring
}

func NewController(uc usecase.Monitoring) *echo.Echo {
	e := echo.New()

	p := prometheus.NewPrometheus("service-monitoring-order", nil)
	p.Use(e)
	c := &controller{uc: uc}

	orders := e.Group("/api/v1")
	orders.GET("/status", c.status)

	return e
}
