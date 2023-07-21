package models

type Product struct {
	Code     int32   `json:"code"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Brand    string  `brand:"brand"`
	Supplier string  `suplier:"suplier"`
}

var products []*Product

func (product *Product) CreateProduct() error {
	products = append(products, product)
	return nil
}

func ReadAllProducts() []*Product {
	return products
}
