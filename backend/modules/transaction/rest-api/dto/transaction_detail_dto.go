package dto

import "time"

type OrderDetail struct {
	TransactionID   int       `form:"transactionId" json:"transacionId"`
	ProductID       int       `form:"productId" json:"productId"`
	Quantity        int       `form:"quantity" json:"quantity"`
	Price           float64   `form:"price" json:"price"`
	SellerID        int       `form:"sellerId" json:"sellerId"`
	Status          string    `form:"status" json:"status"`
	ProductName     string    `form:"productName" json:"productName"`
	CustomerID      int       `form:"customerId" json:"customerId"`
	PaymentOption   string    `form:"paymentOption" json:"paymentOption"`
	TransactionDate time.Time `form:"transactionDate" json:"transactionDate"`
	Name            string    `form:"name" json:"name"`
	Address         string    `form:"address" json:"address"`
}

type TransactionDecentralize struct {
	PicturePath   string  `form:"picturePath" json:"picturePath"`
	ProductName   string  `form:"productName" json:"productName"`
	Price         float64 `form:"price" json:"price"`
	Quantity      int     `form:"quantity" json:"quantity"`
	Address       string  `form:"address" json:"address"`
	PaymentOption string  `form:"paymentOption" json:"paymentOption"`
	SubTotal      float64 `form:"subTotal" json:"subTotal"`
}
