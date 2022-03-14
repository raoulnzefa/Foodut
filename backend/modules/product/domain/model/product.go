package model

import (
	picModel "github.com/Foodut/backend/modules/picture/domain/model"
)

type Product struct {
	ID           int     `form:"id" json:"id" gorm:"primaryKey"`
	ProductName  string  `form:"productName" json:"productName"`
	ProductPrice float64 `form:"productPrice" json:"productPrice"`
	ProductRate  float32 `form:"productRate" json:"productRate"`
	ProductStock int16   `form:"productStock" json:"productStock"`
	SellerID     int     `form:"seller" json:"seller"`
	CategoryID   int
	Picture      []picModel.Picture `gorm:"foreignKey:ProductID;references:ID"`
}

type ProductResponse struct {
	Status  int       `form:"status" json:"status"`
	Message string    `form:"message" json:"message"`
	Data    []Product `form:"data" json:"data"`
}
