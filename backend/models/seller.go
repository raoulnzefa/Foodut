package models

import "fmt"

type Seller struct {
	City        string
	ListProduct []Product
	StoreName   string
	User        User
}

func (s Seller) AddProduct(Product Product) {
	fmt.Println("Seller Add Product")
	// Add Product to database
	// Assign to the seller
}

func (s Seller) DeleterProduct(productId string) {
	fmt.Println("Seller Delete Product")
	// Delete Product from database
}

func (s Seller) EditProduct(productId string) {
	fmt.Println("Seller Edit Product")
	// Update Product in database
}
