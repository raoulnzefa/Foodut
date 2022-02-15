package models

import "fmt"

type Admin struct {
	User User
}

func (a Admin) DeleteProduct(productId string) {
	fmt.Println("Hello Admin Delete Product")
	// If success or not success
}

func (a Admin) EditProduct(productId string) {
	fmt.Println("Hello Admin Edit Product")
	// If success or not success
}
