package server

import (
	orderRoutes "go-service/src/infra/http/order/routes"
	productRoutes "go-service/src/infra/http/product/routes"

	"github.com/labstack/echo/v4"
)

func StartServer() {
	e := echo.New()

	productGroup := e.Group("/product")
	productRoutes.ProductRoutes(productGroup)

	orderGroup := e.Group("/order")
	orderRoutes.OrderRoutes(orderGroup)

	e.Logger.Fatal(e.Start(":3000"))
}
