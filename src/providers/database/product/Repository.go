package databaseProduct

import (
	productEntity "go-service/src/domain/product/entities"
	productNotFoundError "go-service/src/domain/product/errors"
	"go-service/src/providers/database"

	_ "github.com/mattn/go-sqlite3"
)

type ProductRepository struct {
}

func (e ProductRepository) FindBySlug(slug string) (productEntity.Product, error) {
	db := database.InitDatabaseServer()
	stmt, err := db.Prepare("SELECT * FROM products WHERE slug = ?")
	if err != nil {
		panic(err)
	}

	var product productEntity.Product
	err = stmt.QueryRow(slug).Scan(&product.Id, &product.Name, &product.Price, &product.Slug)

	if err != nil {
		return product, productNotFoundError.ProductNotFoundError{}
	}

	database.CloseDatabaseServer(db)

	return product, nil
}

func (e ProductRepository) Create(product productEntity.Product) (productEntity.Product, error) {
	db := database.InitDatabaseServer()
	stmt, err := db.Prepare("INSERT INTO products(id, name, price, slug) values($1, $2, $3, $4)")
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(product.Id, product.Name, product.Price, product.Slug)
	if err != nil {
		panic(err)
	}
	database.CloseDatabaseServer(db)

	return product, nil
}

func (e ProductRepository) Update(product productEntity.Product) (productEntity.Product, error) {
	db := database.InitDatabaseServer()

	stmt, err := db.Prepare("UPDATE products SET name = $1, price = $2, slug = $3 WHERE id = $4")
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(product.Name, product.Price, product.Slug, product.Id)
	if err != nil {
		panic(err)
	}

	database.CloseDatabaseServer(db)

	return product, nil
}

func (e ProductRepository) Delete(id string) (bool, error) {
	db := database.InitDatabaseServer()
	stmt, err := db.Prepare("DELETE FROM products WHERE id = $1")
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(id)

	if err != nil {
		return false, productNotFoundError.ProductNotFoundError{}
	}

	database.CloseDatabaseServer(db)

	return true, nil
}
