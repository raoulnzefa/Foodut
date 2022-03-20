package model

type Cart struct {
	CustomerUserID int `form:"CustomerId" json:"CustomerId" gorm:"primaryKey;autoIncrement:false"`
	ProductID      int `form:"productId" json:"productId" gorm:"primaryKey;autoIncrement:false"`
	Quantity       int `form:"quantity" json:"quantity" gorm:"not null"`
}
