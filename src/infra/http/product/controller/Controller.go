package productController

import (
	productEntity "go-service/src/domain/product/entities"
	productFactory "go-service/src/infra/factory/product"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductController struct{}

func (p *ProductController) Get(c echo.Context) error {
	slug := c.Param("slug")

	factory := productFactory.GetProductFactory()

	response, error := factory.Execute(slug)
	if error != nil {
		return c.JSON(http.StatusNotFound, error.Error())
	}

	return c.JSON(http.StatusFound, response)
}

func (p *ProductController) Create(c echo.Context) error {
	product := productEntity.Product{}
	err := c.Bind(&product)
	if err != nil {
		return err
	}

	factory := productFactory.CreateProductFactory()

	response, error := factory.Execute(product)
	if error != nil {
		return c.JSON(http.StatusBadRequest, error.Error())
	}

	return c.JSON(http.StatusCreated, response)
}

func (p *ProductController) Delete(c echo.Context) error {
	id := c.Param("id")

	factory := productFactory.DeleteProductFactory()

	response, error := factory.Execute(id)
	if error != nil {
		return c.JSON(http.StatusBadRequest, error.Error())
	}

	return c.JSON(http.StatusOK, response)
}

func (p *ProductController) Update(c echo.Context) error {
	product := productEntity.Product{}
	err := c.Bind(&product)
	if err != nil {
		return err
	}
	factory := productFactory.UpdateProductFactory()

	response, error := factory.Execute(product)
	if error != nil {
		return c.JSON(http.StatusBadRequest, error.Error())
	}

	return c.JSON(http.StatusOK, response)
}
