package entities

import (
	"fmt"

	prDto "github.com/Foodut/backend/modules/product/rest-api/dto"
	trDto "github.com/Foodut/backend/modules/transaction/rest-api/dto"
)

type Customer struct {
	User        User                  `form:"user" json:"user"`
	Address     string                `form:"address" json:"address"`
	ListHistory []trDto.Transaction   `form:"transaction" json:"transaction"`
	Cart        []prDto.DetailProduct `form:"cart" json:"cart"`
}

func (c Customer) AddToCart(productId string) {
	fmt.Println("Add product to customer cart")
}

//
// Apa mending digabung aja sama BuyProduct?
// Jadi kek DoPayment, abis itu langsung -
// add history juga di dalem DoPayment??
// Jangan lupa flush cart kalau sudah bayar
func (c Customer) AddToHistory(productId string) {
	fmt.Println("Add transaction to customer history")
}

func (c Customer) BuyProduct(productId string) {
	fmt.Println("Add product to customer cart")
}

// Getter Setter
func (c *Customer) GetAddress() string {
	return c.Address
}

func (c *Customer) SetAddress(address string) {
	c.Address = address
}

func (c *Customer) GetCart() []prDto.DetailProduct {
	return c.Cart
}

func (c *Customer) SetCart(cart []prDto.DetailProduct) {
	c.Cart = cart
}

func (c *Customer) GetListHistory() []trDto.Transaction {
	return c.ListHistory
}

func (c *Customer) SetListHistory(listHistory []trDto.Transaction) {
	c.ListHistory = listHistory
}

func (c *Customer) GetUser() User {
	return c.User
}

func (c *Customer) SetUser(user User) {
	c.User = user
}
