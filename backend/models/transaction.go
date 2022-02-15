package models

import "time"

type Transaction struct {
	Cart            []Cart
	PaymentOption   string
	SubTotal        float64
	TransactionDate time.Time
	TransactionId   string
	Username        string
}

func (t Transaction) GenerateTotalPayment() float64 {
	return 0
	// Calculate from Cart
}
