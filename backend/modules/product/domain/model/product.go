package model

type Product struct {
	ID           int     `form:"id" json:"id" gorm:"primaryKey"`
	ProductName  string  `form:"productName" json:"productName"`
	ProductPrice float64 `form:"productPrice" json:"productPrice"`
	ProductRate  float32 `form:"productRate" json:"productRate"`
	ProductStock int16   `form:"productStock" json:"productStock"`
	SellerID     int     `form:"seller" json:"seller"`
	CategoryID   int
	Picture      []Picture `gorm:"foreignKey:ProductID;references:ID"`
}
