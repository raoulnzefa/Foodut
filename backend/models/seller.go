package models

import "fmt"

type Seller struct {
	User        User      `form:"user" json:"user" gorm:"primaryKey"`
	City        string    `form:"city" json:"city"`
	ListProduct []Product `form:"listProduct" json:"listProduct"`
	StoreName   string    `form:"storeName" json:"storeName"`
}

type SellerResponse struct {
	Status  int      `form:"status" json:"status"`
	Message string   `form:"message" json:"message"`
	Data    []Seller `form:"data" json:"data"`
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
