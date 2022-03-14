package entities

import (
	"time"

	prDto "github.com/Foodut/backend/modules/product/rest-api/dto"
)

type Transaction struct {
	TransactionId   string                `form:"transactionId" json:"transactionId"`
	TransactionDate time.Time             `form:"transactionDate" json:"transactionDate"`
	Username        string                `form:"username" json:"username"`
	PaymentOption   string                `form:"paymentOption" json:"paymentOption"`
	SubTotal        float64               `form:"subTotal" json:"subTotal"`
	Cart            []prDto.DetailProduct `form:"cart" json:"cart"`
}

func (t Transaction) GenerateTotalPayment() float64 {
	// Calculate from Cart
	var result float64 = 0
	return result
}

// Getter Setter
func (t *Transaction) GetCart() []prDto.DetailProduct {
	return t.Cart
}

func (t *Transaction) SetCart(cart []prDto.DetailProduct) {
	t.Cart = cart
}

func (t *Transaction) GetPaymentOption() string {
	return t.PaymentOption
}

func (t *Transaction) SetPaymentOption(paymentOption string) {
	t.PaymentOption = paymentOption
}

func (t *Transaction) GetSubTotal() float64 {
	return t.SubTotal
}

// Udh ada di GenerateTotalPayment ?
func (t *Transaction) SetSubTotal(subTotal float64) {
	t.SubTotal = subTotal
}

func (t *Transaction) GetTransactionDate() time.Time {
	return t.TransactionDate
}

func (t *Transaction) SetTransactionDate(transactionDate time.Time) {
	t.TransactionDate = transactionDate
}

func (t *Transaction) GetTransactionId() string {
	return t.TransactionId
}

func (t *Transaction) SetTransactionId(transactionId string) {
	t.TransactionId = transactionId
}

func (t *Transaction) GetUsername() string {
	return t.Username
}

func (t *Transaction) SetUsername(username string) {
	t.Username = username
}
