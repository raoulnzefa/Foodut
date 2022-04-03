package model

type Product struct {
	ID           int       `form:"id" json:"id" gorm:"primaryKey"`
	ProductName  string    `form:"productName" json:"productName"`
	ProductPrice float64   `form:"productPrice" json:"productPrice"`
	ProductRate  float32   `form:"productRate" json:"productRate"`
	ProductStock int       `form:"productStock" json:"productStock"`
	Description  string    `form:"description" json:"description"`
	SellerID     int       `form:"sellerId" json:"sellerId"`
	CategoryID   int       `form:"categoryId" json:"categoryId"`
	Picture      []Picture `form:"pictures" json:"pictures" gorm:"foreignKey:ProductID;references:ID"`
}
