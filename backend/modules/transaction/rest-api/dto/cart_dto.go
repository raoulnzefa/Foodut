package dto

import "github.com/Foodut/backend/modules/product/rest-api/dto"

type PostCart struct {
	UserId   int                     `form:"userId" json:"userId"`
	Products []dto.PostProductDetail `form:"products" json:"products"`
}

type DeleteSpecificCart struct {
	UserId    int `form:"userId" json:"userId"`
	ProductId int `form:"productId" json:"productId"`
}
