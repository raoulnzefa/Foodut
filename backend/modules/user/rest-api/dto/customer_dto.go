package dto

import (
	prDto "github.com/Foodut/backend/modules/product/rest-api/dto"
	trDto "github.com/Foodut/backend/modules/transaction/rest-api/dto"
)

type Customer struct {
	User        User                  `form:"user" json:"user"`
	Address     string                `form:"address" json:"address"`
	ListHistory []trDto.Transaction   `form:"transaction" json:"transaction"`
	Cart        []prDto.ProductDetail `form:"cart" json:"cart"`
}

// Getter Setter
func (c *Customer) GetAddress() string {
	return c.Address
}

func (c *Customer) SetAddress(address string) {
	c.Address = address
}

func (c *Customer) GetCart() []prDto.ProductDetail {
	return c.Cart
}

func (c *Customer) SetCart(cart []prDto.ProductDetail) {
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
