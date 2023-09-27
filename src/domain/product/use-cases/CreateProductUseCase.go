package productUseCase

import (
	iProductRepository "go-service/src/domain/product/data"
	productEntity "go-service/src/domain/product/entities"

	"github.com/google/uuid"
)

type CreateProductUseCase struct {
	ProductRepository iProductRepository.IProductRepository
}

func (e CreateProductUseCase) Execute(product productEntity.Product) (productEntity.Product, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return product, err
	}
	product.Id = u.String()
	product, error := e.ProductRepository.Create(product)
	if error != nil {
		return product, error
	}
	return product, nil
}
