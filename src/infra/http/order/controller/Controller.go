package orderController

import "github.com/labstack/echo/v4"

type OrderController struct{}

func (p *OrderController) Create(c echo.Context) error {
	return c.JSON(200, "Create order")
}
