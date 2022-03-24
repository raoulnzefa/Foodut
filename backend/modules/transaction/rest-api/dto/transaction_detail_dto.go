package dto

import "time"

type OrderDetail struct {
	TransactionID   int       `form:"transactionId" json:"transacionId" gorm:"primaryKey;autoIncrement:false"`
	ProductID       int       `form:"productId" json:"productId" gorm:"primaryKey;autoIncrement:false"`
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
