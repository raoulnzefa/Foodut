package dto

import (
	"time"

	prDto "github.com/Foodut/backend/modules/product/rest-api/dto"
)

type Transaction struct {
	ID              int                      `form:"id" json:"id" gorm:"primaryKey"`
	CustomerId      int                      `form:"customerId" json:"customerId"`
	Address         string                   `form:"address" json:"address"`
	PaymentOption   string                   `form:"paymentOption" json:"paymentOption"`
	SubTotal        float64                  `form:"subTotal" json:"subTotal"`
	TransactionDate time.Time                `form:"transactionDate" json:"transactionDate"`
	ProductDetail   []prDto.GetProductDetail `form:"productDetail" json:"productDetail"`
}

type PostTransaction struct {
	CustomerId    int    `form:"customerId" json:"customerId"`
	ExtraAddress  string `form:"extraAddress" json:"extraAddress"`
	PaymentOption string `form:"paymentOption" json:"paymentOption"`
}

// Getter Setter
func (t *PostTransaction) GetPaymentOption() string {
	return t.PaymentOption
}

func (t *PostTransaction) SetPaymentOption(paymentOption string) {
	t.PaymentOption = paymentOption
}

// Getter Setter Transaction
func (t *Transaction) GetProductDetail() []prDto.GetProductDetail {
	return t.ProductDetail
}

func (t *Transaction) SetProductDetail(pd []prDto.GetProductDetail) {
	t.ProductDetail = pd
}
