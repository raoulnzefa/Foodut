package dto

import "time"

type OrderDetail struct {
	SellerID        int       `form:"sellerId" json:"sellerId"`
	TransactionID   int       `form:"transactionId" json:"transacionId" gorm:"primaryKey;autoIncrement:false"`
	ProductID       int       `form:"productId" json:"productId" gorm:"primaryKey;autoIncrement:false"`
	ProductName     string    `form:"productName" json:"productName"`
	Price           float64   `form:"price" json:"price"`
	Quantity        int       `form:"quantity" json:"quantity"`
	PaymentOption   string    `form:"paymentOption" json:"paymentOption"`
	Status          string    `form:"status" json:"status"`
	CustomerID      int       `form:"customerId" json:"customerId"`
	Name            string    `form:"name" json:"name"`
	Address         string    `form:"address" json:"address"`
	TransactionDate time.Time `form:"transactionDate" json:"transactionDate"`
}
