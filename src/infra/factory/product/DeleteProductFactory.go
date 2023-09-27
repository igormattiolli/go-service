package productFactory

import (
	productUseCase "go-service/src/domain/product/use-cases"
	databaseProduct "go-service/src/providers/database/product"
)

func DeleteProductFactory() *productUseCase.DeleteProductUseCase {
	p := &databaseProduct.ProductRepository{}
	instance := productUseCase.DeleteProductUseCase{ProductRepository: p}
	return &instance
}
