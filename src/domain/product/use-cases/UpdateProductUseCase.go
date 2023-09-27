package productUseCase

import (
	iProductRepository "go-service/src/domain/product/data"
	productEntity "go-service/src/domain/product/entities"
)

type UpdateProductUseCase struct {
	ProductRepository iProductRepository.IProductRepository
}

func (e UpdateProductUseCase) Execute(product productEntity.Product) (productEntity.Product, error) {
	product, error := e.ProductRepository.Update(product)
	if error != nil {
		return product, error
	}
	return product, nil
}
