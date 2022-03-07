package models

import "time"

type Transaction struct {
	ID              string `form:"id" json:"id" gorm:"primaryKey"`
	CustomerId      string
	DetailProduct   []DetailProduct `form:"detailProduct" json:"detailProduct"`
	PaymentOption   string          `form:"paymentOption" json:"paymentOption"`
	SubTotal        float64         `form:"subTotal" json:"subTotal"`
	TransactionDate time.Time       `form:"transactionDate" json:"transactionDate"`
	Username        string          `form:"username" json:"username"`
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
