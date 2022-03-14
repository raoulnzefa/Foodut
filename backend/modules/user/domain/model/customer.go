package model

import (
	"fmt"

	prModel "github.com/Foodut/backend/modules/product/domain/model"
	trModel "github.com/Foodut/backend/modules/transaction/domain/model"
)

type Customer struct {
	ID            int                     `form:"id" json:"id" gorm:"primaryKey"`
	Address       string                  `form:"address" json:"address"`
	DetailProduct []prModel.DetailProduct `form:"detailProduct" json:"detailProduct" gorm:"foreignKey:ProductID;references:ID"`
	ListHistory   []trModel.Transaction   `form:"listHistory" json:"listHistory" gorm:"foreignKey:CustomerID;references:ID"`
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