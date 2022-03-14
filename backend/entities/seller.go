package entities

import "fmt"

type Seller struct {
	city        string
	listProduct []Product
	storeName   string
	user        User
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
	return s.city
}

func (s *Seller) SetTransactionId(city string) {
	s.city = city
}

func (s *Seller) GetListProduct() []Product {
	return s.listProduct
}

func (s *Seller) SetListProduct(listProduct []Product) {
	s.listProduct = listProduct
}

func (s *Seller) GetStoreName() string {
	return s.storeName
}

func (s *Seller) SetStoreName(storeName string) {
	s.storeName = storeName
}

func (s *Seller) GetUser() User {
	return s.user
}

func (s *Seller) SetUser(user User) {
	s.user = user
}
