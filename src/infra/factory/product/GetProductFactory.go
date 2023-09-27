package productFactory

import (
	productUseCase "go-service/src/domain/product/use-cases"
	databaseProduct "go-service/src/providers/database/product"
)

func GetProductFactory() *productUseCase.GetProductUseCase {
	p := &databaseProduct.ProductRepository{}
	instance := productUseCase.GetProductUseCase{ProductRepository: p}
	return &instance
}
