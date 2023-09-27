package productRoutes

import (
	productController "go-service/src/infra/http/product/controller"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(group *echo.Group) {
	product := productController.ProductController{}
	group.GET("/:slug", product.Get)
	group.POST("/", product.Create)
	group.DELETE("/:id", product.Delete)
	group.PUT("/", product.Update)
}
