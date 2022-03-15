package model

type TransactionDetail struct {
	TransactionID int `form:"transactionId" json:"transacionId" gorm:"primaryKey;autoIncrement:false"`
	ProductID     int `form:"productId" json:"productId" gorm:"primaryKey;autoIncrement:false"`
	Quantity      int `form:"quantity" json:"quantity"`
	Status        int `form:"status" json:"status"`
	// { 0: wishlist, 1: payed }
}
