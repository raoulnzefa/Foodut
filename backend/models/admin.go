package models

import "fmt"

type Admin struct {
	ID int `form:"id" json:"id" gorm:"primaryKey"`
	// Username string `form:"username" json:"username"`
	UserId int
	// User     User
}

type AdminResponse struct {
	Status  int     `form:"status" json:"status"`
	Message string  `form:"message" json:"message"`
	Data    []Admin `form:"data" json:"data"`
}

func (a Admin) DeleteProduct(productId string) {
	fmt.Println("Hello Admin Delete Product")
	// If success or not success
}

func (a Admin) EditProduct(productId string) {
	fmt.Println("Hello Admin Edit Product")
	// If success or not success
}
