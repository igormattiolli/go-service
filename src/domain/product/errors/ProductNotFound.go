package productNotFoundError

type ProductNotFoundError struct {
}

func (e ProductNotFoundError) Error() string {
	return "Product not found"
}
