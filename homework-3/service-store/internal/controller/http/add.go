package http

import (
	"github.com/labstack/echo/v4"
	"gitlab.ozon.dev/bagdatov/homework-3/service-store/internal/entity"
	"log"
	"time"
)

type addRequest struct {
	SellerID uint64 `json:"seller_id"`
	Price    uint64 `json:"price"`
	Name     string `json:"name"`
}

type addResponse struct {
	ItemID uint64 `json:"item_id"`
}

func (c *controller) add(ctx echo.Context) error {

	var req addRequest

	if err := ctx.Bind(&req); err != nil {
		return err
	}

	if req.Price == 0 || req.SellerID == 0 || req.Name == "" {
		return echo.NewHTTPError(400, "Bad request")
	}

	id, err := c.uc.AddItem(ctx.Request().Context(), entity.Item{
		SellerID: req.SellerID,
		Price:    req.Price,
		Name:     req.Name,
		Date:     time.Now(),
	})

	if err != nil {
		log.Printf("add handler -> id: %v, error: %v", req, err)
		return echo.NewHTTPError(500, err.Error())
	}

	return ctx.JSON(200, addResponse{ItemID: id})
}
