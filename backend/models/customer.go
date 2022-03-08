package models

import "fmt"

type Customer struct {
	ID            int             `form:"id" json:"id" gorm:"primaryKey"`
	Address       string          `form:"address" json:"address"`
	DetailProduct []DetailProduct `form:"detailProduct" json:"detailProduct" gorm:"foreignKey:ProductID;references:ID"`
	ListHistory   []Transaction   `form:"listHistory" json:"listHistory" gorm:"foreignKey:CustomerID;references:ID"`
}

type CustomerResponse struct {
	Status  int        `form:"status" json:"status"`
	Message string     `form:"message" json:"message"`
	Data    []Customer `form:"data" json:"data"`
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

//
//

// Apa skip getter setter aja?
func (c Customer) GetListHistory(productId string) {
	fmt.Println("Getter untuk History?")
}
