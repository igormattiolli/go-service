package orderRoutes

import (
	orderController "go-service/src/infra/http/order/controller"

	"github.com/labstack/echo/v4"
)

func OrderRoutes(group *echo.Group) {
	order := orderController.OrderController{}
	group.POST("/", order.Create)
}
