package productFactory

import (
	productUseCase "go-service/src/domain/product/use-cases"
	databaseProduct "go-service/src/providers/database/product"
)

func CreateProductFactory() *productUseCase.CreateProductUseCase {
	p := &databaseProduct.ProductRepository{}
	instance := productUseCase.CreateProductUseCase{ProductRepository: p}
	return &instance
}
