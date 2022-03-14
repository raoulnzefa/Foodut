package entities

import "fmt"

type Seller struct {
	User        User      `form:"user" json:"user"`
	StoreName   string    `form:"storeName" json:"storeName"`
	City        string    `form:"city" json:"city"`
	ListProduct []Product `form:"listProduct" json:"listProduct"`
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

// Getter Setter
func (s *Seller) GetCity() string {
	return s.City
}

func (s *Seller) SetTransactionId(city string) {
	s.City = city
}

func (s *Seller) GetListProduct() []Product {
	return s.ListProduct
}

func (s *Seller) SetListProduct(listProduct []Product) {
	s.ListProduct = listProduct
}

func (s *Seller) GetStoreName() string {
	return s.StoreName
}

func (s *Seller) SetStoreName(storeName string) {
	s.StoreName = storeName
}

func (s *Seller) GetUser() User {
	return s.User
}

func (s *Seller) SetUser(user User) {
	s.User = user
}
