package entities

import "fmt"

type Customer struct {
	address     string
	cart        []DetailProduct
	listHistory []Transaction
	user        User
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

// Getter Setter
func (c *Customer) GetAddress() string {
	return c.address
}

func (c *Customer) SetAddress(address string) {
	c.address = address
}

func (c *Customer) GetCart() []DetailProduct {
	return c.cart
}

func (c *Customer) SetCart(cart []DetailProduct) {
	c.cart = cart
}

func (c *Customer) GetListHistory() []Transaction {
	return c.listHistory
}

func (c *Customer) SetListHistory(listHistory []Transaction) {
	c.listHistory = listHistory
}

func (c *Customer) GetUser() User {
	return c.user
}

func (c *Customer) SetUser(user User) {
	c.user = user
}
