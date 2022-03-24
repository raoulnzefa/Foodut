package dto

import (
	prDto "github.com/Foodut/backend/modules/product/rest-api/dto"
	trDto "github.com/Foodut/backend/modules/transaction/rest-api/dto"
)

type Customer struct {
	User        User                  `form:"user" json:"user"`
	Address     string                `form:"address" json:"address"`
	Cart        []prDto.ProductDetail `form:"cart" json:"cart"`
	ListHistory []trDto.Transaction   `form:"transaction" json:"transaction"`
}

type PostCustomer struct {
	Name     string `form:"name" json:"name"`
	Username string `form:"username" json:"username"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	Address  string `form:"address" json:"address"`
}

type GetCustomer struct {
	ID           int                      `form:"id" json:"id" gorm:"primaryKey"`
	Username     string                   `form:"username" json:"username"`
	Email        string                   `form:"email" json:"email" gorm:"uniqueIndex:idx_email"`
	Name         string                   `form:"name" json:"name"`
	Password     string                   `form:"password" json:"password"`
	ProfilePhoto string                   `form:"profilePhoto" json:"profilePhoto"`
	Level        int                      `form:"level" json:"level"`
	Address      string                   `form:"address" json:"address"`
	Cart         []prDto.GetProductDetail `form:"cart" json:"cart"`
	ListHistory  []trDto.Transaction      `form:"transaction" json:"transaction"`
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
