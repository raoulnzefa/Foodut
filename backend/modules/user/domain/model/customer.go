package model

import (
	"fmt"

	prModel "github.com/Foodut/backend/modules/product/domain/model"
	trModel "github.com/Foodut/backend/modules/transaction/domain/model"
)

type Customer struct {
	UserID      int  `gorm:"primaryKey"`
	User        User `gorm:"foreignKey:UserID;references:ID"`
	Address     string
	TempCart    []prModel.Product     `gorm:"many2many:carts"`
	ListHistory []trModel.Transaction `gorm:"foreignKey:CustomerID;references:UserID"`
}

func (c Customer) AddToCart(productId string) {
	fmt.Println("Add product to customer cart")
}

//
// Apa mending digabung aja sama BuyProduct?
// Jadi kek DoPayment, abis itu langsung -
// add history juga di dalem DoPayment??
func (c Customer) AddToHistory(productId string) {
	fmt.Println("Add transaction to customer history")
}

func (c Customer) BuyProduct(productId string) {
	fmt.Println("Add product to customer cart")
}
