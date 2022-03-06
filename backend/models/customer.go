package models

import "fmt"

type Customer struct {
	Address     string
	Cart        []Cart
	ListHistory []Transaction
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
