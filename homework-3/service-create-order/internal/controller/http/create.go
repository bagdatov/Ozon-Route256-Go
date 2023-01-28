package http

import (
	"github.com/labstack/echo/v4"
	"gitlab.ozon.dev/bagdatov/homework-3/service-create-order/internal/entity"
	"log"
	"time"
)

type createRequest struct {
	ItemID   uint64 `json:"item_id"`
	SellerID uint64 `json:"seller_id"`
	ClientID uint64 `json:"client_id"`
}

type createResponse struct {
	OrderID uint64 `json:"order_id"`
}

func (c *controller) create(ctx echo.Context) error {
	var req createRequest

	if err := ctx.Bind(&req); err != nil {
		return err
	}

	if req.ItemID == 0 || req.SellerID == 0 || req.ClientID == 0 {
		return echo.NewHTTPError(400, "Bad request")
	}

	id, err := c.uc.Create(ctx.Request().Context(), entity.Order{
		ItemID:   req.ItemID,
		SellerID: req.SellerID,
		ClientID: req.ClientID,
		Date:     time.Now(),
	})

	if err != nil {
		log.Printf("create handler -> req: <%+v>, error: %v", req, err)
		return echo.NewHTTPError(500, err.Error())
	}

	return ctx.JSON(200, createResponse{id})
}
