package models

import "time"

type Transaction struct {
	TransactionID   int `form:"id" json:"id" gorm:"primaryKey"`
	CustomerID      int
	PaymentOption   string    `form:"paymentOption" json:"paymentOption"`
	SubTotal        float64   `form:"subTotal" json:"subTotal"`
	TransactionDate time.Time `form:"transactionDate" json:"transactionDate"`
}

type TransactionResponse struct {
	Status  int           `form:"status" json:"status"`
	Message string        `form:"message" json:"message"`
	Data    []Transaction `form:"data" json:"data"`
}

func (t Transaction) GenerateTotalPayment() float64 {
	return 0
	// Calculate from Cart
}
