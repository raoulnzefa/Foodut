package entities

import "fmt"

type Admin struct {
	user User
}

func (a Admin) DeleteProduct(productId string) {
	fmt.Println("Hello Admin Delete Product")
	// If success or not success
}

func (a Admin) EditProduct(productId string) {
	fmt.Println("Hello Admin Edit Product")
	// If success or not success
}

// Getter Setter
func (a *Admin) GetUser() User {
	return a.user
}

func (a *Admin) SetUser(user User) {
	a.user = user
}
