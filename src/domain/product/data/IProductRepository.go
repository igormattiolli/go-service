package productRepository

import productEntity "go-service/src/domain/product/entities"

type IProductRepository interface {
	FindBySlug(string) (productEntity.Product, error)
	Create(productEntity.Product) (productEntity.Product, error)
	Update(productEntity.Product) (productEntity.Product, error)
	Delete(id string) (bool, error)
}
