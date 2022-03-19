package dto

import (
	"time"

	prDto "github.com/Foodut/backend/modules/product/rest-api/dto"
)

type Transaction struct {
	CustomerId      int                   `form:"customerId" json:"customerId"`
	PaymentOption   string                `form:"paymentOption" json:"paymentOption"`
	TransactionDate time.Time             `form:"transactionDate" json:"transactionDate"`
	ProductDetail   []prDto.ProductDetail `form:"productDetail" json:"productDetail"`
}

func (t Transaction) GenerateTotalPayment() float64 {
	// Calculate from Cart
	var result float64 = 0
	return result
}

// Getter Setter
func (t *Transaction) GetCart() []prDto.ProductDetail {
	return t.ProductDetail
}

func (t *Transaction) SetCart(cart []prDto.ProductDetail) {
	t.ProductDetail = cart
}

func (t *Transaction) GetPaymentOption() string {
	return t.PaymentOption
}

func (t *Transaction) SetPaymentOption(paymentOption string) {
	t.PaymentOption = paymentOption
}

func (t *Transaction) GetTransactionDate() time.Time {
	return t.TransactionDate
}

func (t *Transaction) SetTransactionDate(transactionDate time.Time) {
	t.TransactionDate = transactionDate
}
