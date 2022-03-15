package model

import (
	"fmt"

	prModel "github.com/Foodut/backend/modules/product/domain/model"
)

type Seller struct {
	UserID      int               `gorm:"primaryKey"`
	User        User              `gorm:"foreignKey:UserID;references:ID"`
	StoreName   string            `form:"storeName" json:"storeName" gorm:"uniqueIndex:idx_store_name"`
	City        string            `form:"city" json:"city"`
	ListProduct []prModel.Product `gorm:"foreignKey:SellerID;references:UserID"`
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
