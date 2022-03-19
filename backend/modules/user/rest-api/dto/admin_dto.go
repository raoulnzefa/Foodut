package dto

import "fmt"

type Admin struct {
	User User `form:"user" json:"user"`
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
	return a.User
}

func (a *Admin) SetUser(user User) {
	a.User = user
}
