package productUseCase

import (
	iProductRepository "go-service/src/domain/product/data"
	productEntity "go-service/src/domain/product/entities"
)

type GetProductUseCase struct {
	ProductRepository iProductRepository.IProductRepository
}

func (e GetProductUseCase) Execute(slug string) (productEntity.Product, error) {
	product, error := e.ProductRepository.FindBySlug(slug)
	if error != nil {
		return product, error
	}
	return product, nil
}
