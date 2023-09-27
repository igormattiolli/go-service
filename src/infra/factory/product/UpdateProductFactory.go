package productFactory

import (
	productUseCase "go-service/src/domain/product/use-cases"
	databaseProduct "go-service/src/providers/database/product"
)

func UpdateProductFactory() *productUseCase.UpdateProductUseCase {
	p := &databaseProduct.ProductRepository{}
	instance := productUseCase.UpdateProductUseCase{ProductRepository: p}
	return &instance
}
