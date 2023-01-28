package http

import (
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"gitlab.ozon.dev/bagdatov/homework-3/service-store/internal/usecase"
)

type controller struct {
	uc usecase.Storage
}

func NewController(uc usecase.Storage) *echo.Echo {
	e := echo.New()

	p := prometheus.NewPrometheus("service-store", nil)
	p.Use(e)
	c := &controller{uc: uc}

	orders := e.Group("/api/v1/items")
	orders.GET("/find", c.find)
	orders.POST("/add", c.add)

	return e
}
