package model

type TransactionDetail struct {
	TransactionID int     `form:"transactionId" json:"transacionId" gorm:"primaryKey;autoIncrement:false"`
	ProductID     int     `form:"productId" json:"productId" gorm:"primaryKey;autoIncrement:false"`
	SellerID      int     `form:"sellerId" json:"sellerId"`
	Quantity      int     `form:"quantity" json:"quantity"`
	Price         float64 `form:"price" json:"price"`
	Status        string  `form:"status" json:"status"`
}
