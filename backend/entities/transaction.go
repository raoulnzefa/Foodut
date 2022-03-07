package entities

import "time"

type Transaction struct {
	cart            []DetailProduct
	paymentOption   string
	subTotal        float64
	transactionDate time.Time
	transactionId   string
	username        string
}

// Ini agak bingung, apa mau di Setter subTotal aja ?
func (t Transaction) GenerateTotalPayment() float64 {
	// Calculate from Cart
	var result float64 = 0
	return result
}

// Getter Setter
func (t *Transaction) GetCart() []DetailProduct {
	return t.cart
}

func (t *Transaction) SetCart(cart []DetailProduct) {
	t.cart = cart
}

func (t *Transaction) GetPaymentOption() string {
	return t.paymentOption
}

func (t *Transaction) SetPaymentOption(paymentOption string) {
	t.paymentOption = paymentOption
}

func (t *Transaction) GetSubTotal() float64 {
	return t.subTotal
}

// Udh ada di GenerateTotalPayment ?
func (t *Transaction) SetSubTotal(subTotal float64) {
	t.subTotal = subTotal
}

func (t *Transaction) GetTransactionDate() time.Time {
	return t.transactionDate
}

func (t *Transaction) SetTransactionDate(transactionDate time.Time) {
	t.transactionDate = transactionDate
}

func (t *Transaction) GetTransactionId() string {
	return t.transactionId
}

func (t *Transaction) SetTransactionId(transactionId string) {
	t.transactionId = transactionId
}

func (t *Transaction) GetUsername() string {
	return t.username
}

func (t *Transaction) SetUsername(username string) {
	t.username = username
}
