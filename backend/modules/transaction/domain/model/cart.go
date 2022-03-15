package model

type Cart struct {
	CustomerID int `form:"id" json:"id" gorm:"primaryKey;autoIncrement:false"`
	ProductID  int `form:"productId" json:"productId" gorm:"primaryKey;autoIncrement:false"`
	Quantity   int `form:"quantity" json:"quantity"`
}
