package productUseCase

import (
	iProductRepository "go-service/src/domain/product/data"
)

type DeleteProductUseCase struct {
	ProductRepository iProductRepository.IProductRepository
}

func (e DeleteProductUseCase) Execute(id string) (bool, error) {
	response, error := e.ProductRepository.Delete(id)
	if error != nil {
		return response, error
	}
	return response, nil
}
