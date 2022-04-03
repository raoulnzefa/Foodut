package dto

import (
	prDto "github.com/Foodut/backend/modules/product/rest-api/dto"
)

type Seller struct {
	User        User            `form:"user" json:"user"`
	StoreName   string          `form:"storeName" json:"storeName"`
	City        string          `form:"city" json:"city"`
	ListProduct []prDto.Product `form:"listProduct" json:"listProduct"`
}

type GetSeller struct {
	UserID    int    `form:"id" json:"id"`
	StoreName string `form:"storeName" json:"storeName"`
	Username  string `form:"username" json:"username"`
	Email     string `form:"email" json:"email"`
	Name      string `form:"name" json:"name"`
	City      string `form:"city" json:"city"`
}

type PostSeller struct {
	Name      string `form:"name" json:"name"`
	Username  string `form:"username" json:"username"`
	Email     string `form:"email" json:"email"`
	Password  string `form:"password" json:"password"`
	StoreName string `form:"storeName" json:"storeName"`
	City      string `form:"city" json:"city"`
}
