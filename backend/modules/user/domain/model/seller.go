package model

import (
	"fmt"

	prModel "github.com/Foodut/backend/modules/product/domain/model"
)

type Seller struct {
	ID          int               `form:"id" json:"id" gorm:"primaryKey"`
	City        string            `form:"city" json:"city"`
	ListProduct []prModel.Product `gorm:"foreignKey:SellerID;references:ID"`
	StoreName   string            `form:"storeName" json:"storeName"`
}

func (s Seller) AddProduct(Product prModel.Product) {
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
